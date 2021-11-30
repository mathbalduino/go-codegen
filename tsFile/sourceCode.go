package tsFile

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/internal/signature"
)

// SourceCode will build the file source code, including the file imports
// source code
//
// Note that you need to pass the name of the library that
// generated it, in order to provide good debug information
// inside the generated file
//
// The headerTitle param is used to fill the signature header, at the beginning of
// the file
func (f *TsFile) SourceCode(headerTitle string) []byte {
	sourceCode := fmt.Sprintf(sourceCodeTmpl,
		signature.NewSignature(headerTitle),
		f.privateImports.SourceCode(),
		f.sourceCode)

	return []byte(sourceCode)
}

const sourceCodeTmpl = `/*%s*/
%s
%s
`
