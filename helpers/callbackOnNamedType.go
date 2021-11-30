package helpers

import (
	"fmt"
	"go/types"
)

// CallbackOnNamedType will iterate over a types.Type, trying to find some NamedType to pass as argument
// to the callback function
//
// If the Type is an anonymous struct, for example, it will recursively iterate over the fields until it
// finds some Named type, or a basic type (in which case, the callback is not called). Note that the callback
// can be called more than just once (maybe not even called)
func CallbackOnNamedType(fieldType types.Type, callback func(obj *types.Named)) error {
	switch type_ := fieldType.(type) {
	case *types.Basic:
		return nil
	case *types.Pointer:
		return CallbackOnNamedType(type_.Elem(), callback)
	case *types.Array:
		return CallbackOnNamedType(type_.Elem(), callback)
	case *types.Slice:
		return CallbackOnNamedType(type_.Elem(), callback)
	case *types.Map:
		e := CallbackOnNamedType(type_.Elem(), callback)
		if e != nil {
			return e
		}
		return CallbackOnNamedType(type_.Key(), callback)
	case *types.Chan:
		return CallbackOnNamedType(type_.Elem(), callback)
	case *types.Struct:
		for i := 0; i < type_.NumFields(); i++ {
			e := CallbackOnNamedType(type_.Field(i).Type(), callback)
			if e != nil {
				return e
			}
		}
		return nil
	case *types.Tuple:
		for i := 0; i < type_.Len(); i++ {
			e := CallbackOnNamedType(type_.At(i).Type(), callback)
			if e != nil {
				return e
			}
		}
		return nil
	case *types.Signature:
		// Receiver is ignored
		e := CallbackOnNamedType(type_.Params(), callback)
		if e != nil {
			return e
		}
		return CallbackOnNamedType(type_.Results(), callback)
	case *types.Named:
		callback(type_)
		return nil
	case *types.Interface:
		for i := 0; i < type_.NumMethods(); i++ {
			e := CallbackOnNamedType(type_.Method(i).Type(), callback)
			if e != nil {
				return e
			}
		}
		return nil
	default:
		return UnexpectedTypeError
	}
}

var UnexpectedTypeError = fmt.Errorf("the given types.Type interface underlying type is not recognizable")
