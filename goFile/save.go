package goFile

import (
	"fmt"
	"os"
)

// Save will save the file to disk, ready to be compiled
//
// The headerTitle param is used to fill the signature header,
// at the beginning of the file (usually, the library name)
func (f *GoFile) Save(headerTitle, folder string) error {
	filename := fmt.Sprintf("%s/%s", folder, f.name)
	sourceCode, e := f.SourceCode(headerTitle, filename)
	if e != nil {
		return e
	}

	createdFile, e := osCreate(filename)
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

// just to ease tests
var osCreate = os.Create
