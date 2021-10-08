package goFile

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/goFile/goImports"
	"github.com/mathbalduino/go-codegen/internal/signature"
)

// privateImports is just an alias that will
// provide a way to embed the type as a
// private field
type privateImports = *goImports.GoImports

// GoFile holds information about
// a generate GO file
type GoFile struct {
	// name will be the name of the
	// generated file (without folderpath)
	name string

	// packageName is the name of the
	// package that the file will belong
	packageName string

	// sourceCode is the file source code
	// that isn't the 'package' and 'import'
	// keywords
	sourceCode string

	// privateImports holds information about
	// the imports that the file contains
	privateImports
}

// New will create a new GO file
//
// Note that the 'filename' cannot contain the folderpath, and the
// 'packageName'/'packagePath' refers to the package that the file
// will belong to
func New(filename, packageName, packagePath string) *GoFile {
	return &GoFile{
		fmt.Sprintf("%s%s.go", filename, signature.FileSuffix),
		packageName,
		"",
		goImports.New(packagePath),
	}
}

// NewTestFile will create a new GO test file
//
// Note that the 'filename' cannot contain the folderpath, and the
// 'packageName'/'packagePath' refers to the package that the file
// will belong to
func NewTestFile(filename, packageName, packagePath string) *GoFile {
	return &GoFile{
		fmt.Sprintf("%s%s._test.go", filename, signature.FileSuffix),
		packageName,
		"",
		goImports.New(packagePath),
	}
}
