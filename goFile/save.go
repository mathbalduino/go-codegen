package goFile

import (
	"fmt"
	"os"
)

// Save will save the file to disk, ready to be compiled
//
// Note that you need to pass the name of the library that
// generated it, in order to provide good debug information
// inside the generated file
func (f *GoFile) Save(libraryThatGeneratedIt, folder string) error {
	filename := fmt.Sprintf("%s/%s", folder, f.name)
	sourceCode, e := f.SourceCode(libraryThatGeneratedIt, filename)
	if e != nil {
		return e
	}

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
