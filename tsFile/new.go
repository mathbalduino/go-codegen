package tsFile

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/internal/signature"
	"github.com/mathbalduino/go-codegen/tsFile/tsImports"
)

// privateImports is just an alias that
// provide a way to embed an Exported type
// as a private field
type privateImports = *tsImports.TsImports

// TsFile holds information about
// a Typescript - TS file, with name
// and source code
type TsFile struct {
	// name will be the name of the
	// generated file (without folderpath)
	name string

	// sourceCode is the file source code
	// without the imports
	sourceCode string

	// privateImports holds information about
	// the imports that the file contains
	privateImports
}

// New will create a new pointer to a TsFile
// with the given filename
//
// Note that the given filename cannot contain information
// about the folderpath
func New(filename string) *TsFile {
	return &TsFile{
		fmt.Sprintf("%s%s.ts", filename, signature.FileSuffix),
		"",
		tsImports.New(),
	}
}
