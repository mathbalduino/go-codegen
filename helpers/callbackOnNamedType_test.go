package helpers

import (
	"go/types"
	"testing"
)

func TestCallbackOnNamedType(t *testing.T) {
	t.Run("Basic types should not do any recursive calls, not call the callback, just return", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(&types.Basic{}, callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Pointer type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewPointer(&types.Basic{}), callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Array type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewArray(&types.Basic{}, 1), callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Slice type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewSlice(&types.Basic{}), callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Map type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewMap(&types.Basic{}, &types.Basic{}), callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Map type, returning any errors", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewMap(&types.Basic{}, &fakeType{}), callback)
		if e != UnexpectedTypeError {
			t.Fatalf("Returned error is not the expected one")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Chan type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewChan(types.SendOnly, &types.Basic{}), callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Struct type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		struct_ := types.NewStruct([]*types.Var{types.NewVar(0, nil, "", &types.Basic{})}, nil)
		e := CallbackOnNamedType(struct_, callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Struct type, returning any errors", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		struct_ := types.NewStruct([]*types.Var{types.NewVar(0, nil, "", &fakeType{})}, nil)
		e := CallbackOnNamedType(struct_, callback)
		if e != UnexpectedTypeError {
			t.Fatalf("Error is not the expected one")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Tuple type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})), callback)
		if e != nil {
			t.Fatalf("Error expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Tuple type, returning any errors", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(types.NewTuple(types.NewVar(0, nil, "", &fakeType{})), callback)
		if e != UnexpectedTypeError {
			t.Fatalf("Error is not the expected one")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Signature type, ignoring the receiver type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		signature := types.NewSignature(
			// Should ignore this receiver NamedType
			types.NewVar(0, nil, "", types.NewNamed(&types.TypeName{}, nil, nil)),

			types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})),
			types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})),
			false,
		)
		e := CallbackOnNamedType(signature, callback)
		if e != nil {
			t.Fatalf("Error is expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Signature type, ignoring the receiver type, returning any errors", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		signature := types.NewSignature(
			// Should ignore this receiver NamedType
			types.NewVar(0, nil, "", types.NewNamed(&types.TypeName{}, nil, nil)),

			types.NewTuple(types.NewVar(0, nil, "", &fakeType{})),
			types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})),
			false,
		)
		e := CallbackOnNamedType(signature, callback)
		if e != UnexpectedTypeError {
			t.Fatalf("Error is not the expected one")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should immediately call the callback if it's a Named type", func(t *testing.T) {
		calls := 0
		typeName := types.NewNamed(&types.TypeName{}, nil, nil)
		callback := func(obj *types.Named) {
			calls += 1
			if typeName != obj {
				t.Fatalf("Wrong named type passed as argument to the callback")
			}
		}
		e := CallbackOnNamedType(typeName, callback)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if calls != 1 {
			t.Fatalf("Callback was expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Interface type, ignoring any receiver type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		signature := types.NewSignature(
			// Should ignore this receiver NamedType
			types.NewVar(0, nil, "", types.NewNamed(&types.TypeName{}, nil, nil)),

			types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})),
			types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})),
			false,
		)
		interface_ := types.NewInterfaceType([]*types.Func{types.NewFunc(0, nil, "", signature)}, nil)
		interface_.Complete()
		e := CallbackOnNamedType(interface_, callback)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Interface type, ignoring any receiver type, returning errors", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		signature := types.NewSignature(
			// Should ignore this receiver NamedType
			types.NewVar(0, nil, "", types.NewNamed(&types.TypeName{}, nil, nil)),

			types.NewTuple(types.NewVar(0, nil, "", &fakeType{})),
			types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})),
			false,
		)
		interface_ := types.NewInterfaceType([]*types.Func{types.NewFunc(0, nil, "", signature)}, nil)
		interface_.Complete()
		e := CallbackOnNamedType(interface_, callback)
		if e != UnexpectedTypeError {
			t.Fatalf("Error is not the expected one")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should return an error when facing unrecognizable types", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		e := CallbackOnNamedType(&fakeType{}, callback)
		if e != UnexpectedTypeError {
			t.Fatalf("Not the expected error")
		}
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
}

type fakeType struct{}

func (f *fakeType) Underlying() types.Type { return nil }
func (f *fakeType) String() string         { return "fake" }
