package goImports

import "testing"

func TestSourceCode(t *testing.T) {
	t.Run("If the import list is empty, return an empty string", func(t *testing.T) {
		i := &GoImports{"", map[string]string{}}
		if i.SourceCode() != "" {
			t.Fatalf("Expected an empty string")
		}
	})
	t.Run("Write valid GO source code from the import list", func(t *testing.T) {
		i := &GoImports{"", map[string]string{
			"aliasA": "pkg/path/a",
			"aliasB": "pkg/path/b",
		}}
		expected := "import (\n" +
			"aliasA \"pkg/path/a\"\n" +
			"aliasB \"pkg/path/b\"\n" +
			")\n"
		if i.SourceCode() != expected {
			t.Fatalf("Not the expected source code")
		}
	})
}
