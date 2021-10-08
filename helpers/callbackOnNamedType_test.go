package helpers

import (
	goParser "github.com/mathbalduino/go-codegen"
	"go/types"
	"testing"
)

func TestCallbackOnNamedType(t *testing.T) {
	t.Run("Basic types should not do any recursive calls, not call the callback, just return", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(&types.Basic{}, callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Pointer type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(types.NewPointer(&types.Basic{}), callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Array type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(types.NewArray(&types.Basic{}, 1), callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Slice type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(types.NewSlice(&types.Basic{}), callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Map type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(types.NewMap(&types.Basic{}, &types.Basic{}), callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Chan type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(types.NewChan(types.SendOnly, &types.Basic{}), callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Struct type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		struct_ := types.NewStruct([]*types.Var{types.NewVar(0, nil, "", &types.Basic{})}, nil)
		CallbackOnNamedType(struct_, callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should recursively iterate over the Tuple type, until Basic/Named type", func(t *testing.T) {
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(types.NewTuple(types.NewVar(0, nil, "", &types.Basic{})), callback, nil)
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
		CallbackOnNamedType(signature, callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should immediately call the callback", func(t *testing.T) {
		calls := 0
		typeName := types.NewNamed(&types.TypeName{}, nil, nil)
		callback := func(obj *types.Named) {
			calls += 1
			if typeName != obj {
				t.Fatalf("Wrong named type passed as argument to the callback")
			}
		}
		CallbackOnNamedType(typeName, callback, nil)
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
		CallbackOnNamedType(interface_, callback, nil)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
	})
	t.Run("Should call Log.Fatal when facing unrecognizable types", func(t *testing.T) {
		fatalCalls := 0
		mock := &mockLogCLI{mockFatal: func(msgFormat string, args ...interface{}) {
			fatalCalls += 1
			if msgFormat != unexpectedTypeMsg {
				t.Fatalf("Wrong msg for unexpected types error")
			}
			if len(args) != 1 || args[0].(string) != "fake" {
				t.Fatalf("Wrong variadic argument for unexpected types error")
			}
		}}
		calls := 0
		callback := func(obj *types.Named) { calls += 1 }
		CallbackOnNamedType(&fakeType{}, callback, mock)
		if calls != 0 {
			t.Fatalf("Callback was not expected to be called")
		}
		if fatalCalls != 1 {
			t.Fatalf("Log.Fatal was expected to be called")
		}
	})
}

type fakeType struct{}

func (f *fakeType) Underlying() types.Type { return nil }
func (f *fakeType) String() string         { return "fake" }

// -----

func emptyMockLogCLI() goParser.LogCLI {
	m := &mockLogCLI{}
	m.mockDebug = func(msgFormat string, args ...interface{}) goParser.LogCLI { return m }
	m.mockError = func(msgFormat string, args ...interface{}) goParser.LogCLI { return m }
	m.mockFatal = func(msgFormat string, args ...interface{}) {}
	return m
}

type mockLogCLI struct {
	mockDebug func(msgFormat string, args ...interface{}) goParser.LogCLI
	mockError func(msgFormat string, args ...interface{}) goParser.LogCLI
	mockFatal func(msgFormat string, args ...interface{})
}

func (m *mockLogCLI) Debug(msgFormat string, args ...interface{}) goParser.LogCLI {
	if m.mockDebug == nil {
		return nil
	}
	return m.mockDebug(msgFormat, args...)
}

func (m *mockLogCLI) Error(msgFormat string, args ...interface{}) goParser.LogCLI {
	if m.mockError == nil {
		return nil
	}
	return m.mockError(msgFormat, args...)
}

func (m *mockLogCLI) Fatal(msgFormat string, args ...interface{}) {
	if m.mockFatal != nil {
		m.mockFatal(msgFormat, args...)
	}
}
