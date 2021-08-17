package helpers

import (
	goParser "gitlab.com/matheuss-leonel/go-codegen"
	"go/types"
)

// TypeIsAccessible will check to see if some type is fully accessible from the given
// package (is it public? Private? Builtin? All struct fields exported?...)
func TypeIsAccessible(t types.Type, fromPackagePath string, log goParser.LogCLI) bool {
	switch type_ := t.(type) {
	case *types.Basic:
		log.Debug("Accessible: Is a basic type (*types.Basic)")
		return true

	case *types.Pointer:
		ptrLog := log.Debug("Pointer type (*types.Pointer). Checking it's element type...")
		return TypeIsAccessible(type_.Elem(), fromPackagePath, ptrLog)

	case *types.Array:
		arrLog := log.Debug("Array type (*types.Array). Checking it's element type...")
		return TypeIsAccessible(type_.Elem(), fromPackagePath, arrLog)

	case *types.Slice:
		sliceLog := log.Debug("Slice type (*types.Slice). Checking it's element type...")
		return TypeIsAccessible(type_.Elem(), fromPackagePath, sliceLog)

	case *types.Map:
		mapLog := log.Debug("Map type (*types.Map). Checking it's element and key type...")
		return TypeIsAccessible(type_.Elem(), fromPackagePath, mapLog) &&
			TypeIsAccessible(type_.Key(), fromPackagePath, mapLog)

	case *types.Chan:
		chanLog := log.Debug("Channel type (*types.Chan). Checking it's element type...")
		return TypeIsAccessible(type_.Elem(), fromPackagePath, chanLog)

	case *types.Struct:
		structLog := log.Debug("Anonymous struct type (*types.Struct). Checking it's fields...")
		if type_.NumFields() == 0 {
			structLog.Debug("Accessible: doesn't have any fields")
			return true
		}

		for i := 0; i < type_.NumFields(); i++ {
			field := type_.Field(i)
			fieldLog := structLog.Debug("Field '%s'...", field.Name())

			// struct fields will never have nil Pkg()
			if field.Pkg().Path() != fromPackagePath {
				if !field.Exported() {
					fieldLog.Debug("Not accessible: different package (%s %s) and not exported",
						field.Pkg().Name(), field.Pkg().Path())
					return false
				}

				fieldLog.Debug("Accessible: different package (%s %s) but exported",
					field.Pkg().Name(), field.Pkg().Path())
			} else {
				fieldLog.Debug("Accessible: same package")
			}

			fieldType := field.Type()
			fieldTypeLog := fieldLog.Debug("Checking field type '%s'...", fieldType.String())
			if !TypeIsAccessible(fieldType, fromPackagePath, fieldTypeLog) {
				return false
			}
		}

		structLog.Debug("Accessible: all fields (and it's types) are exported or are in the same package")
		return true

	case *types.Tuple:
		tupleLog := log.Debug("Tuple type (*types.Tuple, function parameters or multiple assignments). Checking it's members...")
		if type_.Len() == 0 {
			tupleLog.Debug("Accessible: doesn't have any members")
			return true
		}

		for i := 0; i < type_.Len(); i++ {
			member := type_.At(i)
			memberType := member.Type()
			memberLog := tupleLog.Debug("Member '%s', type '%s'...", member.Name(), memberType.String())

			if !TypeIsAccessible(memberType, fromPackagePath, memberLog) {
				return false
			}
		}

		tupleLog.Debug("Accessible: all members types are exported or are in the same package")
		return true

	case *types.Signature:
		signLog := log.Debug("Signature/Function type (*types.Signature). Checking it's parameters and return values...")
		return TypeIsAccessible(type_.Params(), fromPackagePath, signLog) && TypeIsAccessible(type_.Results(), fromPackagePath, signLog)

	case *types.Named:
		namedLog := log.Debug("Named type (*types.Named)")
		if type_.Obj().Pkg() == nil {
			namedLog.Debug("Accessible: builtin type")
			return true
		}
		if type_.Obj().Pkg().Path() != fromPackagePath {
			if !type_.Obj().Exported() {
				namedLog.Debug("Not accessible: different package (%s %s) and not exported",
					type_.Obj().Pkg().Name(), type_.Obj().Pkg().Path())
				return false
			}

			namedLog.Debug("Accessible: different package (%s %s) but exported",
				type_.Obj().Pkg().Name(), type_.Obj().Pkg().Path())
		} else {
			namedLog.Debug("Accessible: same package")
		}
		return true

	case *types.Interface:
		interfaceLog := log.Debug("Interface type (*types.Interface). Checking it's methods...")
		if type_.NumMethods() == 0 {
			interfaceLog.Debug("Accessible: doesn't have any methods")
			return true
		}

		for i := 0; i < type_.NumMethods(); i++ {
			method := type_.Method(i)
			methodLog := interfaceLog.Debug("Method '%s'...", method.Name())

			// methods will never have nil Pkg()
			if method.Pkg().Path() != fromPackagePath {
				if !method.Exported() {
					methodLog.Debug("Not accessible: different package (%s %s) and not exported",
						method.Pkg().Name(), method.Pkg().Path())
					return false
				}
				methodLog.Debug("Accessible: different package (%s %s) but exported",
					method.Pkg().Name(), method.Pkg().Path())
			} else {
				methodLog.Debug("Accessible: same package")
			}

			methodType := method.Type()
			methodTypeLog := methodLog.Debug("Checking method type '%s'...", methodType.String())
			if !TypeIsAccessible(methodType, fromPackagePath, methodTypeLog) {
				return false
			}
		}

		interfaceLog.Debug("Accessible: all methods (and it's types) are exported or are in the same package")
		return true

	default:
		log.Fatal(unexpectedTypeMsg, type_.String())
		return false
	}
}
