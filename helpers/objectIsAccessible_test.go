package helpers

import (
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/types"
	"testing"
)

func TestObjectIsAccessible(t *testing.T) {
	t.Run("Should return true for builtin types", func(t *testing.T) {
		type_ := types.NewVar(0, nil, "any", &types.Basic{})
		b := ObjectIsAccessible(type_, "pkgName", loggerCLI.New(false, 0))
		if !b {
			t.Fatalf("Expected to return true")
		}
	})
	t.Run("Should return true for types inside the given package", func(t *testing.T) {
		p := types.NewPackage("packagePath", "packageName")
		type_ := types.NewVar(0, p, "any", &types.Basic{})
		b := ObjectIsAccessible(type_, p.Path(), loggerCLI.New(false, 0))
		if !b {
			t.Fatalf("Expected to return true")
		}
	})
	t.Run("Should return false for not exported types of a different package", func(t *testing.T) {
		p := types.NewPackage("packagePath", "packageName")
		type_ := types.NewVar(0, p, "any", &types.Basic{})
		b := ObjectIsAccessible(type_, "otherPkg", loggerCLI.New(false, 0))
		if b {
			t.Fatalf("Expected to return false")
		}
	})
	t.Run("Should return true for exported types of a different package", func(t *testing.T) {
		p := types.NewPackage("packagePath", "packageName")
		type_ := types.NewVar(0, p, "Any", &types.Basic{})
		b := ObjectIsAccessible(type_, "otherPkg", loggerCLI.New(false, 0))
		if !b {
			t.Fatalf("Expected to return true")
		}
	})
}
