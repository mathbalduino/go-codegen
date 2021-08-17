package goImports

import "testing"

func TestNeedImport(t *testing.T) {
	t.Run("If the provided string is equal to the packagePath, return false", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		i := &GoImports{pkgPath, nil}
		if i.NeedImport(pkgPath) {
			t.Fatalf("Expected to be false")
		}
	})
	t.Run("If the provided string is different from the packagePath, return true", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		i := &GoImports{pkgPath, nil}
		if !i.NeedImport("another/pkg/path") {
			t.Fatalf("Expected to be true")
		}
	})
}
