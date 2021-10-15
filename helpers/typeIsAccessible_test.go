package helpers

import (
	"fmt"
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/types"
	"testing"
)

func TestTypeIsAccessible(t *testing.T) {
	t.Run("Basic types should always return true", func(t *testing.T) {
		ok := TypeIsAccessible(types.Typ[1], "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively iterate over pointer elements type", func(t *testing.T) {
		type_ := types.NewPointer(types.Typ[2])
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively iterate over array elements type", func(t *testing.T) {
		type_ := types.NewArray(types.Typ[2], 2)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively iterate over slice elements type", func(t *testing.T) {
		type_ := types.NewSlice(types.Typ[2])
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively iterate over map elements and keys type", func(t *testing.T) {
		type_ := types.NewMap(types.Typ[1], types.Typ[2])
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively iterate over chan elements type", func(t *testing.T) {
		type_ := types.NewChan(types.SendRecv, types.Typ[2])
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Empty structs should always return true", func(t *testing.T) {
		type_ := types.NewStruct([]*types.Var{}, nil)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Structs with at least one field from another package, not exported, return false", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath", "pkgName")
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, pkg, "FieldA", types.Typ[1]),
			types.NewVar(0, pkg, "fieldB", types.Typ[1]),
			types.NewVar(0, pkg, "FieldC", types.Typ[1]),
		}, nil)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Structs with all fields from another package, all exported, return true", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath", "pkgName")
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, pkg, "FieldA", types.Typ[1]),
			types.NewVar(0, pkg, "FieldB", types.Typ[1]),
			types.NewVar(0, pkg, "FieldC", types.Typ[1]),
		}, nil)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Structs with all fields from same package, exported or not, return true", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath", "pkgName")
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, pkg, "FieldA", types.Typ[1]),
			types.NewVar(0, pkg, "fieldB", types.Typ[1]),
			types.NewVar(0, pkg, "FieldC", types.Typ[1]),
			types.NewVar(0, pkg, "fieldD", types.Typ[1]),
		}, nil)
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively check every struct field type. If they're all accessible, return true", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "PublicTypeName", nil),
			types.Typ[1],
			nil)

		pkg := types.NewPackage("pkgPath", "pkgName")
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, pkg, "fieldA", types.Typ[1]),
			types.NewVar(0, pkg, "fieldB", namedType),
			types.NewVar(0, pkg, "fieldC", types.Typ[1]),
		}, nil)
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Recursively check every struct field type. If there a single one not accessible, return false", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "privateTypeName", nil),
			types.Typ[1],
			nil)

		pkg := types.NewPackage("pkgPath", "pkgName")
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, pkg, "fieldA", types.Typ[1]),
			types.NewVar(0, pkg, "fieldB", namedType),
			types.NewVar(0, pkg, "fieldC", types.Typ[1]),
		}, nil)
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Empty tuples should always return true", func(t *testing.T) {
		type_ := types.NewTuple()
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Should iterate over every tuple element type. If they're all accessible, return true", func(t *testing.T) {
		type_ := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[2]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Should iterate over every tuple element type. If there's a single one not accessible, return false", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "privateTypeName", nil),
			types.Typ[1],
			nil)

		type_ := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", namedType),
			types.NewVar(0, nil, "", types.Typ[3]),
		)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Signatures with at least one param type and one result type not accessible should return false", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "privateTypeName", nil),
			types.Typ[1],
			nil)

		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", namedType),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", namedType),
			types.NewVar(0, nil, "", types.Typ[3]),
		)
		type_ := types.NewSignature(nil, params, results, false)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Signatures with one param type not accessible should return false", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "privateTypeName", nil),
			types.Typ[1],
			nil)

		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", namedType),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[4]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)
		type_ := types.NewSignature(nil, params, results, false)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Signatures with one result type not accessible should return false", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "privateTypeName", nil),
			types.Typ[1],
			nil)

		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", types.Typ[4]),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", namedType),
			types.NewVar(0, nil, "", types.Typ[3]),
		)
		type_ := types.NewSignature(nil, params, results, false)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Signatures with all param and result types accessible should return true", func(t *testing.T) {
		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", types.Typ[4]),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)
		type_ := types.NewSignature(nil, params, results, false)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Named types with nil package should return true (builtin type, like 'error')", func(t *testing.T) {
		type_ := types.NewNamed(
			types.NewTypeName(0, nil, "someTypeName", nil),
			types.Typ[1],
			nil)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Private named types from another package should return false", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath2", "pkgName2")
		type_ := types.NewNamed(
			types.NewTypeName(0, pkg, "someTypeName", nil),
			types.Typ[1],
			nil)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Public named types from another package should return true", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath2", "pkgName2")
		type_ := types.NewNamed(
			types.NewTypeName(0, pkg, "SomeTypeName", nil),
			types.Typ[1],
			nil)
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Named types from same package should always return true", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath2", "pkgName2")
		type_ := types.NewNamed(
			types.NewTypeName(0, pkg, "someTypeName", nil),
			types.Typ[1],
			nil)
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}

		type_ = types.NewNamed(
			types.NewTypeName(0, pkg, "SomeTypeName", nil),
			types.Typ[1],
			nil)
		ok = TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Empty interfaces should always return true", func(t *testing.T) {
		type_ := types.NewInterfaceType(nil, nil)
		type_.Complete()
		ok := TypeIsAccessible(type_, "", loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Interfaces with all methods from same package, and accessible signatures, should always return true", func(t *testing.T) {
		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)

		pkg := types.NewPackage("pkgPath2", "pkgName2")
		type_ := types.NewInterfaceType([]*types.Func{
			types.NewFunc(0, pkg, "PublicMethod", types.NewSignature(nil, params, results, false)),
			types.NewFunc(0, pkg, "privateMethod", types.NewSignature(nil, results, params, false)),
		}, nil)
		type_.Complete()
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Interfaces with at least one not accessible signature, should return false", func(t *testing.T) {
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		namedType := types.NewNamed(
			types.NewTypeName(0, pkg2, "privateNamedType", nil),
			types.Typ[1],
			nil)
		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", namedType),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)

		pkg := types.NewPackage("pkgPath", "pkgName")
		type_ := types.NewInterfaceType([]*types.Func{
			types.NewFunc(0, pkg, "PublicMethod", types.NewSignature(nil, params, results, false)),
			types.NewFunc(0, pkg, "privateMethod", types.NewSignature(nil, results, params, false)),
		}, nil)
		type_.Complete()
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Interfaces with at least one private method from another package, should return false", func(t *testing.T) {
		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)

		pkg := types.NewPackage("pkgPath", "pkgName")
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		type_ := types.NewInterfaceType([]*types.Func{
			types.NewFunc(0, pkg, "PublicMethod", types.NewSignature(nil, params, results, false)),
			types.NewFunc(0, pkg2, "privateMethod", types.NewSignature(nil, results, params, false)),
		}, nil)
		type_.Complete()
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if ok {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("Interfaces with all methods from another package being public, and all methods with accessible signature, should return true", func(t *testing.T) {
		params := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[5]),
			types.NewVar(0, nil, "", types.Typ[2]),
		)
		results := types.NewTuple(
			types.NewVar(0, nil, "", types.Typ[1]),
			types.NewVar(0, nil, "", types.Typ[3]),
		)

		pkg := types.NewPackage("pkgPath", "pkgName")
		pkg2 := types.NewPackage("pkgPath2", "pkgName2")
		type_ := types.NewInterfaceType([]*types.Func{
			types.NewFunc(0, pkg2, "PublicMethod", types.NewSignature(nil, params, results, false)),
			types.NewFunc(0, pkg, "privateMethod", types.NewSignature(nil, results, params, false)),
		}, nil)
		type_.Complete()
		ok := TypeIsAccessible(type_, pkg.Path(), loggerCLI.New(false, false, false))
		if !ok {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Not recognizable types should call Log.Fatal, panicking", func(t *testing.T) {
		fatalCalls := 0
		ch := make(chan bool)
		go func() {
			defer func() {
				e := recover()
				if e != nil && e.(error).Error() == fmt.Sprintf(unexpectedTypeMsg, (&fakeType{}).String()) {
					fatalCalls += 1
				}
				ch <- true
			}()

			TypeIsAccessible(&fakeType{}, "", loggerCLI.New(false, false, false))
		}()

		<-ch
		if fatalCalls != 1 {
			t.Fatalf("Expected to call panic one time")
		}
	})
}
