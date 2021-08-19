package goFile

import (
	"gitlab.com/matheuss-leonel/go-codegen/internal/signature"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Should never return a nil pointer", func(t *testing.T) {
		if New("", "", "") == nil {
			t.Fatalf("Not expected to be nil")
		}
	})
	t.Run("Should set the filename correctly", func(t *testing.T) {
		file := New("meuFile", "", "")
		if file.name != "meuFile" + signature.FileSuffix + ".go" {
			t.Fatalf("Filename is wrong")
		}
	})
	t.Run("Should set the packageName correctly", func(t *testing.T) {
		file := New("", "pkgName", "")
		if file.packageName != "pkgName" {
			t.Fatalf("PackageName is wrong")
		}
	})
	t.Run("The source code must be an empty string", func(t *testing.T) {
		file := New("", "", "")
		if file.sourceCode != "" {
			t.Fatalf("Source code must be empty")
		}
	})
	t.Run("Should create a GoImports object point to the given packagePath", func(t *testing.T) {
		file := New("", "", "pkgPath")
		if file.privateImports.PackagePath() != "pkgPath" {
			t.Fatalf("Wrong GoImports packagePath")
		}
	})
}

func TestNewTestFile(t *testing.T) {
	t.Run("Should never return a nil pointer", func(t *testing.T) {
		if NewTestFile("", "", "") == nil {
			t.Fatalf("Not expected to be nil")
		}
	})
	t.Run("Should set the filename correctly", func(t *testing.T) {
		file := NewTestFile("meuFile", "", "")
		if file.name != "meuFile" + signature.FileSuffix + "._test.go" {
			t.Fatalf("Filename is wrong")
		}
	})
	t.Run("Should set the packageName correctly", func(t *testing.T) {
		file := NewTestFile("", "pkgName", "")
		if file.packageName != "pkgName" {
			t.Fatalf("PackageName is wrong")
		}
	})
	t.Run("The source code must be an empty string", func(t *testing.T) {
		file := NewTestFile("", "", "")
		if file.sourceCode != "" {
			t.Fatalf("Source code must be empty")
		}
	})
	t.Run("Should create a GoImports object point to the given packagePath", func(t *testing.T) {
		file := NewTestFile("", "", "pkgPath")
		if file.privateImports.PackagePath() != "pkgPath" {
			t.Fatalf("Wrong GoImports packagePath")
		}
	})
}
