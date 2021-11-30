package helpers

import (
	"fmt"
	"go/types"
	"strings"
)

// ResolveTypeIdentifier will take some type and return a string that represents it
//
// Example: the type "pointer to integer" will return the string "*int"
func ResolveTypeIdentifier(t types.Type, pkgImports GoImports) (string, error) {
	switch type_ := t.(type) {

	case *types.Basic:
		return type_.Name(), nil

	case *types.Pointer:
		typeIdentifier, e := ResolveTypeIdentifier(type_.Elem(), pkgImports)
		if e != nil {
			return "", e
		}
		return fmt.Sprintf("*%s", typeIdentifier), nil

	case *types.Array:
		typeIdentifier, e := ResolveTypeIdentifier(type_.Elem(), pkgImports)
		if e != nil {
			return "", e
		}
		return fmt.Sprintf("[%d]%s", type_.Len(), typeIdentifier), nil

	case *types.Slice:
		typeIdentifier, e := ResolveTypeIdentifier(type_.Elem(), pkgImports)
		if e != nil {
			return "", e
		}
		return fmt.Sprintf("[]%s", typeIdentifier), nil

	case *types.Map:
		keyTypeIdentifier, e := ResolveTypeIdentifier(type_.Key(), pkgImports)
		if e != nil {
			return "", e
		}
		elemTypeIdentifier, e := ResolveTypeIdentifier(type_.Elem(), pkgImports)
		if e != nil {
			return "", e
		}
		return fmt.Sprintf("map[%s]%s", keyTypeIdentifier, elemTypeIdentifier), nil

	case *types.Chan:
		typeIdentifier, e := ResolveTypeIdentifier(type_.Elem(), pkgImports)
		if e != nil {
			return "", e
		}
		if type_.Dir() == types.SendOnly {
			return fmt.Sprintf("chan<- %s", typeIdentifier), nil
		}
		if type_.Dir() == types.RecvOnly {
			return fmt.Sprintf("<-chan %s", typeIdentifier), nil
		}
		return fmt.Sprintf("chan %s", typeIdentifier), nil

	case *types.Struct:
		str := "struct{"
		for i := 0; i < type_.NumFields(); i++ {
			field := type_.Field(i)
			fieldTypeIdentifier, e := ResolveTypeIdentifier(field.Type(), pkgImports)
			if e != nil {
				return "", e
			}
			str += fmt.Sprintf("%s %s; ", field.Name(), fieldTypeIdentifier)
		}
		str = strings.TrimSuffix(str, "; ")
		return str + "}", nil

	case *types.Tuple:
		str := ""
		for i := 0; i < type_.Len(); i++ {
			typeIdentifier, e := ResolveTypeIdentifier(type_.At(i).Type(), pkgImports)
			if e != nil {
				return "", e
			}

			// Note that the name is ignored. If this Tuple doesn't belongs to a Signature
			// (it is a multiple assignment), problems can arise (always expected to be part of a signature)
			str += fmt.Sprintf("%s, ", typeIdentifier)
		}
		return strings.TrimSuffix(str, ", "), nil

	case *types.Signature:
		hasParams := type_.Params().Len() > 0
		hasResults := type_.Results().Len() > 0

		params := ""
		if hasParams {
			if type_.Variadic() {
				// Make a new Tuple without the last variadic param (possibly empty)
				var vars []*types.Var
				for i := 0; i < type_.Params().Len()-1; i++ {
					vars = append(vars, type_.Params().At(i))
				}
				withoutLastParam, e := ResolveTypeIdentifier(types.NewTuple(vars...), pkgImports)
				if e != nil {
					return "", e
				}

				// Resolve the last variadic param identifier alone
				lastParam, e := ResolveTypeIdentifier(types.NewTuple(type_.Params().At(type_.Params().Len()-1)), pkgImports)
				if e != nil {
					return "", e
				}

				// Check if there's just the variadic param
				params = withoutLastParam
				if params != "" {
					params += ", "
				}

				// Replace the two chars that are representing the variadic type as an slice
				params += "..." + lastParam[2:]
			} else {
				var e error
				params, e = ResolveTypeIdentifier(type_.Params(), pkgImports)
				if e != nil {
					return "", e
				}
			}
		}
		if !hasResults {
			return fmt.Sprintf("func(%s)", params), nil
		}

		results, e := ResolveTypeIdentifier(type_.Results(), pkgImports)
		if e != nil {
			return "", e
		}
		if type_.Results().Len() > 1 {
			results = "(" + results + ")"
		}
		return fmt.Sprintf("func(%s) %s", params, results), nil

	case *types.Named:
		if type_.Obj().Pkg() == nil {
			// Types with (package == nil) are basic types, like the "error" interface
			return type_.Obj().Name(), nil
		}

		if pkgImports.NeedImport(type_.Obj().Pkg().Path()) {
			return fmt.Sprintf("%s.%s",
				pkgImports.AliasFromPath(type_.Obj().Pkg().Path()),
				type_.Obj().Name()), nil
		}
		return type_.Obj().Name(), nil

	case *types.Interface:
		str := "interface{"
		for i := 0; i < type_.NumMethods(); i++ {
			currMethod := type_.Method(i)
			typeIdentifier, e := ResolveTypeIdentifier(currMethod.Type(), pkgImports)
			if e != nil {
				return "", e
			}
			str += fmt.Sprintf("%s; ", strings.ReplaceAll(typeIdentifier, "func", currMethod.Name()))
		}
		str = strings.TrimSuffix(str, "; ")
		return str + "}", nil

	default:
		return "", UnexpectedTypeError
	}
}
