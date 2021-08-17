package goImports

import "testing"

func TestNew(t *testing.T) {
	t.Run("Should never return nil pointer", func(t *testing.T) {
		if New("") == nil {
			t.Fatalf("Nil pointer received from New function")
		}
	})
	t.Run("Should correctly set the packagePath field", func(t *testing.T) {
		pkgPath := "packagePathName"
		i := New(pkgPath)
		if i.packagePath != pkgPath {
			t.Fatalf("Wrong package path")
		}
	})
	t.Run("Should set non-nil, zero empty, imports map", func(t *testing.T) {
		pkgPath := "packagePathName"
		i := New(pkgPath)
		if i.imports == nil {
			t.Fatalf("imports field cannot be nil")
		}
		if len(i.imports) != 0 {
			t.Fatalf("imports field cannot be filled")
		}
	})
}
