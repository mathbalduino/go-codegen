package goFile

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/goFile/goImports"
	"github.com/mathbalduino/go-codegen/internal/signature"
	"strings"
	"testing"
)

func TestSourceCode(t *testing.T) {
	t.Run("Should add the correct header, between long commentaries (/* */)", func(t *testing.T) {
		f := &GoFile{packageName: "somePkg", privateImports: goImports.New("")}
		headerTitle := "Header Title - LibName v1.0.4"
		filepath := "/"
		code, e := f.SourceCode(headerTitle, filepath)
		if e != nil {
			t.Fatalf("Error not expected")
		}
		if !strings.HasPrefix(string(code), fmt.Sprintf("/*%s*/", signature.NewSignature(headerTitle))) {
			t.Fatalf("Header missing/wrong")
		}
	})
	t.Run("Should add imports", func(t *testing.T) {
		imports := goImports.New("")
		imports.AddImport("fmt", "fmt")
		f := &GoFile{packageName: "somePkg", privateImports: imports}
		f.AddCode("var abc = fmt.Sprintf(\"abc%s\", \"d\")")
		filepath := "/"
		code, e := f.SourceCode("", filepath)
		if e != nil {
			t.Fatalf("Error not expected")
		}
		if !strings.Contains(string(code), "import (\n\tfmt \"fmt\"\n)") {
			t.Fatalf("Expected to add the imports")
		}
	})
	t.Run("Should format the source code", func(t *testing.T) {
		sourceCode := "  const   abc    =  `some string`"
		f := &GoFile{packageName: "somePkg", privateImports: goImports.New(""), sourceCode: sourceCode}
		filepath := "/"
		code, e := f.SourceCode("", filepath)
		if e != nil {
			t.Fatalf("Error not expected")
		}
		if !strings.HasSuffix(string(code), "const abc = `some string`\n") {
			t.Fatalf("Code not formatted")
		}
	})
	t.Run("Should optimize imports", func(t *testing.T) {
		imports := goImports.New("")
		imports.AddImport("fmt", "fmt") // should be removed, since its not used
		f := &GoFile{packageName: "somePkg", privateImports: imports}
		filepath := "/"
		code, e := f.SourceCode("", filepath)
		if e != nil {
			t.Fatalf("Error not expected")
		}
		if strings.Contains(string(code), "import") {
			t.Fatalf("Expected to optimize imports")
		}
	})
	t.Run("Should abort and return nil data + errors with the formatting operation, if any", func(t *testing.T) {
		// Note that the source code is syntactically wrong (missing '}')
		f := &GoFile{packageName: "somePkg", privateImports: goImports.New(""), sourceCode: "func abc() {"}
		filepath := "/"
		code, e := f.SourceCode("", filepath)
		if e == nil {
			t.Fatalf("Error expected")
		}
		if code != nil {
			t.Fatalf("Expected to be nil when there are errors")
		}
	})
}
