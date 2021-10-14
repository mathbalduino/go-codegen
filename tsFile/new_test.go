package tsFile

import (
	"github.com/mathbalduino/go-codegen/internal/signature"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Should never return nil pointer", func(t *testing.T) {
		if New("") == nil {
			t.Fatalf("Not expected to be nil")
		}
	})
	t.Run("Should set filename to the given one", func(t *testing.T) {
		filename := "name"
		file := New(filename)
		if file.name != filename+signature.FileSuffix+".ts" {
			t.Fatalf("Incorrect filename")
		}
	})
	t.Run("Should set source code to an empty", func(t *testing.T) {
		file := New("")
		if file.sourceCode != "" {
			t.Fatalf("Source code must e empty")
		}
	})
	t.Run("Should set a non nil imports", func(t *testing.T) {
		file := New("")
		if file.privateImports == nil {
			t.Fatalf("TsImports must not be nil")
		}
	})
}
