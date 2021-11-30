package helpers

import (
	"go/types"
)

// TypeIsAccessible will check to see if some type is fully accessible from the given
// package (is it public? Private? Builtin? All struct fields exported?...)
func TypeIsAccessible(t types.Type, fromPackagePath string) (bool, error) {
	switch type_ := t.(type) {
	case *types.Basic:
		return true, nil

	case *types.Pointer:
		return TypeIsAccessible(type_.Elem(), fromPackagePath)

	case *types.Array:
		return TypeIsAccessible(type_.Elem(), fromPackagePath)

	case *types.Slice:
		return TypeIsAccessible(type_.Elem(), fromPackagePath)

	case *types.Map:
		elem, e := TypeIsAccessible(type_.Elem(), fromPackagePath)
		if e != nil {
			return false, e
		}
		key, e := TypeIsAccessible(type_.Key(), fromPackagePath)
		return elem && key, e

	case *types.Chan:
		return TypeIsAccessible(type_.Elem(), fromPackagePath)

	case *types.Struct:
		if type_.NumFields() == 0 {
			return true, nil
		}
		for i := 0; i < type_.NumFields(); i++ {
			field := type_.Field(i)

			// struct fields will never have nil Pkg()
			if field.Pkg().Path() != fromPackagePath && !field.Exported() {
				return false, nil
			}

			fieldType := field.Type()
			fieldBool, e := TypeIsAccessible(fieldType, fromPackagePath)
			if e != nil {
				return false, e
			}
			if !fieldBool {
				return false, nil
			}
		}
		return true, nil

	case *types.Tuple:
		if type_.Len() == 0 {
			return true, nil
		}
		for i := 0; i < type_.Len(); i++ {
			member := type_.At(i)
			memberType := member.Type()
			memberBool, e := TypeIsAccessible(memberType, fromPackagePath)
			if e != nil {
				return false, e
			}
			if !memberBool {
				return false, nil
			}
		}
		return true, nil

	case *types.Signature:
		a, e := TypeIsAccessible(type_.Params(), fromPackagePath)
		if e != nil {
			return false, e
		}
		b, e := TypeIsAccessible(type_.Results(), fromPackagePath)
		if e != nil {
			return false, e
		}
		return a && b, nil

	case *types.Named:
		if type_.Obj().Pkg() == nil {
			return true, nil
		}
		if type_.Obj().Pkg().Path() != fromPackagePath && !type_.Obj().Exported() {
			return false, nil
		}
		return true, nil

	case *types.Interface:
		if type_.NumMethods() == 0 {
			return true, nil
		}
		for i := 0; i < type_.NumMethods(); i++ {
			method := type_.Method(i)

			// methods will never have nil Pkg()
			if method.Pkg().Path() != fromPackagePath && !method.Exported() {
				return false, nil
			}

			methodType := method.Type()
			methodBool, e := TypeIsAccessible(methodType, fromPackagePath)
			if e != nil {
				return false, e
			}
			if !methodBool {
				return false, nil
			}
		}
		return true, nil

	default:
		return false, UnexpectedTypeError
	}
}
