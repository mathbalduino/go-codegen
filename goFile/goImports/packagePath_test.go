package goImports

import "testing"

func TestPackagePath(t *testing.T) {
	t.Run("Should act just as a getter", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		i := &GoImports{pkgPath, nil}
		if i.PackagePath() != pkgPath {
			t.Fatalf("Expected to act as a getter")
		}
	})
}
