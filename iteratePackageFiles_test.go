package parser

import (
	"fmt"
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestIteratePackageFiles(t *testing.T) {
	t.Run("Should return nil errors when there are no Package.Syntax objects", func(t *testing.T) {
		p := &GoParser{
			pkgs: []*packages.Package{
				{Syntax: []*ast.File{}},
			},
			logger: loggerCLI.New(false, 0),
		}
		e := p.iteratePackageFiles(func(currFile *ast.File, filePkg *packages.Package, parentLog LoggerCLI) error {
			return nil
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should skip files that are not the focus", func(t *testing.T) {
		// Mocks the focused file in order to recognize it correctly
		focusedFilePos, focusedFileSize := token.Pos(100), 200
		focusedFile := "focusedFile"
		fileSet := token.NewFileSet()
		fileSet.AddFile(focusedFile, int(focusedFilePos), focusedFileSize)
		fileSet.AddFile("otherFiles", int(focusedFilePos)+focusedFileSize+1, 10)

		p := &GoParser{
			pkgs: []*packages.Package{
				{Syntax: []*ast.File{
					{Package: focusedFilePos + token.Pos(focusedFileSize+1)},
					{Package: focusedFilePos},
					{Package: focusedFilePos + token.Pos(focusedFileSize+1)},
				}},
			},
			logger:  loggerCLI.New(false, 0),
			fileSet: fileSet,
			focus:   FocusFilePath(focusedFile),
		}
		calls := 0
		e := p.iteratePackageFiles(func(currFile *ast.File, filePkg *packages.Package, parentLog LoggerCLI) error {
			calls += 1
			if fileSet.File(currFile.Pos()).Name() != focusedFile {
				t.Fatalf("Callback was not expected to be called with non focused files")
			}
			return nil
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if calls != 1 {
			t.Fatalf("The callback was expected to be called one time")
		}
	})
	t.Run("Should return any errors returned by the callback", func(t *testing.T) {
		fileSet := token.NewFileSet()
		fileSet.AddFile("file", 1, 5)

		p := &GoParser{
			pkgs: []*packages.Package{
				{Syntax: []*ast.File{
					{Package: 2},
					{Package: 2},
					{Package: 2},
				}},
			},
			logger:  loggerCLI.New(false, 0),
			fileSet: fileSet,
		}
		calls := 0
		e := p.iteratePackageFiles(func(currFile *ast.File, filePkg *packages.Package, parentLog LoggerCLI) error {
			calls += 1
			return fmt.Errorf("any error")
		})
		if e == nil {
			t.Fatalf("Expected to return the callback error")
		}
		if calls != 1 {
			t.Fatalf("The callback was expected to be called one time")
		}
	})
	t.Run("Should call the callback for every file that needs to be iterated and return nil error", func(t *testing.T) {
		fileSet := token.NewFileSet()
		fileSet.AddFile("a", 1, 5)
		fileSet.AddFile("b", 10, 5)
		fileSet.AddFile("c", 20, 5)

		p := &GoParser{
			pkgs: []*packages.Package{
				{Syntax: []*ast.File{
					{Package: 2},
					{Package: 11},
					{Package: 21},
				}},
			},
			logger:  loggerCLI.New(false, 0),
			fileSet: fileSet,
		}
		callsA, callsB, callsC := 0, 0, 0
		e := p.iteratePackageFiles(func(currFile *ast.File, filePkg *packages.Package, parentLog LoggerCLI) error {
			switch fileSet.File(currFile.Pos()).Name() {
			case "a":
				callsA += 1
				return nil
			case "b":
				callsB += 1
				return nil
			case "c":
				callsC += 1
				return nil
			default:
				t.Fatalf("Unexpected file iteration")
				return nil
			}
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if callsA != 1 || callsB != 1 || callsC != 1 {
			t.Fatalf("Each file must be iterated one time")
		}
	})
}
