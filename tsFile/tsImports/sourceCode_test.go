package tsImports

import "testing"

func TestSourceCode(t *testing.T) {
	t.Run("If the import list is empty, return and empty string", func(t *testing.T) {
		i := make(TsImports, 0, 1)
		if i.SourceCode() != "" {
			t.Fatalf("Expected an empty string")
		}
	})
	t.Run("Should build the expected source code", func(t *testing.T) {
		i := make(TsImports, 0, 1)
		i = append(i,
			&tsImport{"some/path", "DefaultImportName", []string{"NamedA", "NamedB", "NamedC"}},
			&tsImport{"some/path/2", "DefaultImportName_2", []string{"NamedA_2", "NamedB_2", "NamedC_2"}})
		code := i.SourceCode()
		expectedCode := "import DefaultImportName, { NamedA, NamedB, NamedC } from 'some/path'\n" +
			"import DefaultImportName_2, { NamedA_2, NamedB_2, NamedC_2 } from 'some/path/2'\n"
		if code != expectedCode {
			t.Fatalf("Expected an empty string")
		}
	})
}
