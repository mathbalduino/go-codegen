package tsFile

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/tsFile/tsImports"
	"io/ioutil"
	"os"
	"testing"
)

func TestSave(t *testing.T) {
	t.Run("Should abort and return any errors when creating the file", func(t *testing.T) {
		folder := "/folderName"
		filename := "filename"
		f := &TsFile{name: filename, privateImports: tsImports.New()}
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
	t.Run("Should return any errors when trying to write to the file", func(t *testing.T) {
		f := &TsFile{name: "filename", privateImports: tsImports.New()}
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
		f := &TsFile{name: "filename", privateImports: tsImports.New()}
		tmpFile, e := ioutil.TempFile(os.TempDir(), "tmpfile-")
		defer os.Remove(tmpFile.Name())
		osCreate = func(name string) (*os.File, error) { return tmpFile, e }
		e = f.Save("", "/folder")
		if e != nil {
			t.Fatalf("Error not expected")
		}

		expectedSourceCode := f.SourceCode("")
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
