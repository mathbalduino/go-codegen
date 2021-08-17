package goImports

import "testing"

func TestAliasFromPath(t *testing.T) {
	t.Run("If the given package path is not in the import list, return an empty string", func(t *testing.T) {
		i := &GoImports{"", map[string]string{}}
		if i.AliasFromPath("some/pkg/path") != "" {
			t.Fatalf("Expected to be empty string")
		}
	})
	t.Run("Should return the alias if the given package path is in the import list", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		alias := "alias"
		i := &GoImports{"", map[string]string{alias: pkgPath}}
		if i.AliasFromPath(pkgPath) != alias {
			t.Fatalf("Expected to return the alias")
		}
	})
}
