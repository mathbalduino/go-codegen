package helpers

import (
	"fmt"
	goParser "github.com/mathbalduino/go-codegen"
	"go/types"
	"strings"
)

// ResolveTypeIdentifier will take some type and return a string that represents it
//
// Example: an integer pointer type will return "*int"
func ResolveTypeIdentifier(t types.Type, pkgImports GoImports, log goParser.LogCLI) string {
	switch type_ := t.(type) {

	case *types.Basic:
		return type_.Name()

	case *types.Pointer:
		typeIdentifier := ResolveTypeIdentifier(type_.Elem(), pkgImports, log)
		return fmt.Sprintf("*%s", typeIdentifier)

	case *types.Array:
		typeIdentifier := ResolveTypeIdentifier(type_.Elem(), pkgImports, log)
		return fmt.Sprintf("[%d]%s", type_.Len(), typeIdentifier)

	case *types.Slice:
		typeIdentifier := ResolveTypeIdentifier(type_.Elem(), pkgImports, log)
		return fmt.Sprintf("[]%s", typeIdentifier)

	case *types.Map:
		keyTypeIdentifier := ResolveTypeIdentifier(type_.Key(), pkgImports, log)
		elemTypeIdentifier := ResolveTypeIdentifier(type_.Elem(), pkgImports, log)
		return fmt.Sprintf("map[%s]%s", keyTypeIdentifier, elemTypeIdentifier)

	case *types.Chan:
		typeIdentifier := ResolveTypeIdentifier(type_.Elem(), pkgImports, log)
		if type_.Dir() == types.SendOnly {
			return fmt.Sprintf("chan<- %s", typeIdentifier)
		}
		if type_.Dir() == types.RecvOnly {
			return fmt.Sprintf("<-chan %s", typeIdentifier)
		}
		return fmt.Sprintf("chan %s", typeIdentifier)

	case *types.Struct:
		str := "struct{"
		for i := 0; i < type_.NumFields(); i++ {
			field := type_.Field(i)
			fieldTypeIdentifier := ResolveTypeIdentifier(field.Type(), pkgImports, log)
			str += fmt.Sprintf("%s %s; ", field.Name(), fieldTypeIdentifier)
		}
		str = strings.TrimSuffix(str, "; ")
		return str + "}"

	case *types.Tuple:
		str := ""
		for i := 0; i < type_.Len(); i++ {
			typeIdentifier := ResolveTypeIdentifier(type_.At(i).Type(), pkgImports, log)

			// Note that the name is ignored. If this Tuple doesn't belongs to a Signature
			// (it is a multiple assignment), problems can arise (always expected to be part of a signature)
			str += fmt.Sprintf("%s, ", typeIdentifier)
		}
		return strings.TrimSuffix(str, ", ")

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
				withoutLastParam := ResolveTypeIdentifier(types.NewTuple(vars...), pkgImports, log)

				// Resolve the last variadic param identifier alone
				lastParam := ResolveTypeIdentifier(
					types.NewTuple(type_.Params().At(type_.Params().Len()-1)), pkgImports, log)

				// Check if there's just the variadic param
				params = withoutLastParam
				if params != "" {
					params += ", "
				}

				// Replace the two chars that are representing the variadic type as an slice
				params += "..." + lastParam[2:]
			} else {
				params = ResolveTypeIdentifier(type_.Params(), pkgImports, log)
			}
		}
		if !hasResults {
			return fmt.Sprintf("func(%s)", params)
		}

		results := ResolveTypeIdentifier(type_.Results(), pkgImports, log)
		if type_.Results().Len() > 1 {
			results = "(" + results + ")"
		}
		return fmt.Sprintf("func(%s) %s", params, results)

	case *types.Named:
		if type_.Obj().Pkg() == nil {
			// Types with (package == nil) are basic types, like the "error" interface
			return type_.Obj().Name()
		}

		if pkgImports.NeedImport(type_.Obj().Pkg().Path()) {
			return fmt.Sprintf("%s.%s",
				pkgImports.AliasFromPath(type_.Obj().Pkg().Path()),
				type_.Obj().Name())
		}
		return type_.Obj().Name()

	case *types.Interface:
		str := "interface{"
		for i := 0; i < type_.NumMethods(); i++ {
			currMethod := type_.Method(i)
			typeIdentifier := ResolveTypeIdentifier(currMethod.Type(), pkgImports, log)
			str += fmt.Sprintf("%s; ", strings.ReplaceAll(typeIdentifier, "func", currMethod.Name()))
		}
		str = strings.TrimSuffix(str, "; ")
		return str + "}"

	default:
		log.Fatal(unexpectedTypeMsg, type_.String())
		return ""
	}
}
