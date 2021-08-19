package goFile

import (
	"fmt"
	"gitlab.com/matheuss-leonel/go-codegen/goFile/goImports"
	"io/ioutil"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	t.Run("Should abort and return any errors when building the source code", func(t *testing.T) {
		// Note that the source code is syntactically wrong (missing '}')
		f := &GoFile{packageName: "somePkg", privateImports: goImports.New(""), sourceCode: "func abc() {"}
		e := f.Save("", "/")
		if e == nil {
			t.Fatalf("Error expected")
		}
	})
	t.Run("Should abort and return any errors when creating the file", func(t *testing.T) {
		folder := "/folderName"
		filename := "filename"
		f := &GoFile{name: filename, packageName: "somePkg", privateImports: goImports.New("")}
		calls := 0
		osCreate = func(name string) (*os.File, error) {
			calls += 1
			if name != folder+"/"+filename {
				t.Fatalf("Wrong filename given to osCreate")
			}
			return nil, fmt.Errorf("some error")
		}

		e := f.Save("", folder)
		if e == nil {
			t.Fatalf("Error expected")
		}
		if calls != 1 {
			t.Fatalf("Expected one call to osCreate")
		}
	})
	t.Run("Should return any errors when trying to write to the fil", func(t *testing.T) {
		f := &GoFile{name: "filename", packageName: "somePkg", privateImports: goImports.New("")}
		tmpFile, e := ioutil.TempFile(os.TempDir(), "tmpfile-")
		defer os.Remove(tmpFile.Name())
		tmpFile.Close() // Just to get an error when trying to write
		osCreate = func(name string) (*os.File, error) { return tmpFile, e }
		e = f.Save("", "/folder")
		if e == nil {
			t.Fatalf("Error expected")
		}
	})
	t.Run("Should correctly write the source code to the file", func(t *testing.T) {
		f := &GoFile{name: "filename", packageName: "somePkg", privateImports: goImports.New("")}
		tmpFile, e := ioutil.TempFile(os.TempDir(), "tmpfile-")
		defer os.Remove(tmpFile.Name())
		osCreate = func(name string) (*os.File, error) { return tmpFile, e }
		e = f.Save("", "/folder")
		if e != nil {
			t.Fatalf("Error not expected")
		}

		expectedSourceCode, e := f.SourceCode("", "/folder")
		if e != nil {
			t.Fatalf("Error not expected")
		}
		fileData := make([]byte, len(expectedSourceCode))
		tmpFile, e = os.Open(tmpFile.Name())
		if e != nil {
			t.Fatalf("Error not expected")
		}
		_, e = tmpFile.Read(fileData)
		if e != nil {
			t.Fatalf("Error not expected")
		}
		if string(fileData) != string(expectedSourceCode) {
			t.Fatalf("File with wrong content")
		}
		tmpFile.Close()
	})
}
