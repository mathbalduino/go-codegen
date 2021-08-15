package helpers

import (
	goParser "gitlab.com/matheuss-leonel/go-codegen"
	"go/types"
)

// CallbackOnNamedType will iterate over a types.Type, trying to find some NamedType to pass as argument
// to the callback function. If the Type is an anonymous struct, for example, it will recursively iterate
// over the fields until it finds some Named type, or a basic type (in which case, the callback is not called)
func CallbackOnNamedType(fieldType types.Type, callback func(obj *types.Named), log goParser.LogCLI) {
	switch type_ := fieldType.(type) {
	case *types.Basic:
		return
	case *types.Pointer:
		CallbackOnNamedType(type_.Elem(), callback, log)
	case *types.Array:
		CallbackOnNamedType(type_.Elem(), callback, log)
	case *types.Slice:
		CallbackOnNamedType(type_.Elem(), callback, log)
	case *types.Map:
		CallbackOnNamedType(type_.Elem(), callback, log)
		CallbackOnNamedType(type_.Key(), callback, log)
	case *types.Chan:
		CallbackOnNamedType(type_.Elem(), callback, log)
	case *types.Struct:
		for i := 0; i < type_.NumFields(); i++ {
			CallbackOnNamedType(type_.Field(i).Type(), callback, log)
		}
	case *types.Tuple:
		for i := 0; i < type_.Len(); i++ {
			CallbackOnNamedType(type_.At(i).Type(), callback, log)
		}
	case *types.Signature:
		// Receiver is ignored
		CallbackOnNamedType(type_.Params(), callback, log)
		CallbackOnNamedType(type_.Results(), callback, log)
	case *types.Named:
		callback(type_)
	case *types.Interface:
		for i := 0; i < type_.NumMethods(); i++ {
			CallbackOnNamedType(type_.Method(i).Type(), callback, log)
		}
	default:
		log.Fatal(unexpectedTypeMsg, type_.String())
		return
	}
}

const unexpectedTypeMsg = "unexpected type: %s"
