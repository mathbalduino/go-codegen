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
		expectedA := "import (\n" +
			"aliasA \"pkg/path/a\"\n" +
			"aliasB \"pkg/path/b\"\n" +
			")\n"
		expectedB := "import (\n" +
			"aliasB \"pkg/path/b\"\n" +
			"aliasA \"pkg/path/a\"\n" +
			")\n"

		// Maps don't preserve the order of the elements,
		// that's why we test two possibilities
		received := i.SourceCode()
		if received != expectedA && received != expectedB {
			t.Fatalf("Not the expected source code")
		}
	})
}
