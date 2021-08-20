package tsFile

import (
	"fmt"
	"os"
)

// Save will take all the file information and save it to a real file
// in disk
//
// Note that you need to pass the name of the library that
// generated it, in order to provide good debug information
// inside the generated file
func (f *TsFile) Save(libraryThatGeneratedIt, folder string) error {
	sourceCode, e := f.SourceCode(libraryThatGeneratedIt)
	if e != nil {
		return e
	}

	filename := fmt.Sprintf("%s/%s", folder, f.name)
	createdFile, e := os.Create(filename)
	if e != nil {
		return e
	}
	_, e = createdFile.Write(sourceCode)
	if e != nil {
		createdFile.Close()
		return e
	}
	return createdFile.Close()
}
