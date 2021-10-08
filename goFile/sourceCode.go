package goFile

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/internal/signature"
	"golang.org/x/tools/imports"
)

// SourceCode will return the content of the file, with the 'package' and 'import'
// keywords
//
// The filepath param is used to optimize the file imports. It must contain the
// folderpath information with the filename
//
// The headerTitle param is used to fill the signature header, at the beginning of
// the file
func (f *GoFile) SourceCode(headerTitle, filepath string) ([]byte, error) {
	sourceCode := fmt.Sprintf(sourceCodeTmpl,
		signature.NewSignature(headerTitle),
		f.packageName,
		f.privateImports.SourceCode(),
		f.sourceCode)

	importOptimizedCode, e := imports.Process(filepath, []byte(sourceCode), nil)
	if e != nil {
		return nil, e
	}

	return importOptimizedCode, nil
}

const sourceCodeTmpl = `/*%s*/
package %s

%s

%s
`
