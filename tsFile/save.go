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
//
// The headerTitle param is used to fill the signature header, at the beginning of
// the file
func (f *TsFile) Save(headerTitle, folder string) error {
	sourceCode := f.SourceCode(headerTitle)
	filename := fmt.Sprintf("%s/%s", folder, f.name)
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
