package tsFile

import (
	"fmt"
	"gitlab.com/matheuss-leonel/go-codegen/internal/signature"
	"gitlab.com/matheuss-leonel/go-codegen/tsFile/tsImports"
	"strings"
	"testing"
)

func TestSourceCode(t *testing.T) {
	t.Run("Should add the correct header, between long commentaries (/* */)", func(t *testing.T) {
		f := &TsFile{privateImports: tsImports.New()}
		headerTitle := "Header Title - LibName v1.0.4"
		code := f.SourceCode(headerTitle)
		if !strings.HasPrefix(string(code), fmt.Sprintf("/*%s*/", signature.NewSignature(headerTitle))) {
			t.Fatalf("Header missing/wrong")
		}
	})
	t.Run("Should add imports", func(t *testing.T) {
		f := &TsFile{privateImports: tsImports.New()}
		f.privateImports.AddNamedImport("NamedA", "some/path")
		headerTitle := "Header Title - LibName v1.0.4"
		code := f.SourceCode(headerTitle)
		if !strings.Contains(string(code), "import { NamedA } from 'some/path'") {
			t.Fatalf("Missing imports")
		}
	})
}
