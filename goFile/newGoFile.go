package goFile

import (
	"fmt"
	"gitlab.com/matheuss-leonel/go-codegen/goFile/goImports"
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

// NewGoFile will create a new GO file
//
// Note that the 'filename' cannot contain the folderpath, and the
// 'packageName'/'packagePath' refers to the package that the file
// will belong to
func NewGoFile(filename, packageName, packagePath string) *GoFile {
	return &GoFile{
		name:           fmt.Sprintf("%s%s.go", filename, loxe.GeneratedFileSuffix),
		packageName:    packageName,
		sourceCode:     "",
		privateImports: goImports.New(packagePath),
	}
}

// NewGoTestFile will create a new GO test file
//
// Note that the 'filename' cannot contain the folderpath, and the
// 'packageName'/'packagePath' refers to the package that the file
// will belong to
func NewGoTestFile(filename, packageName, packagePath string) *GoFile {
	return &GoFile{
		name:           fmt.Sprintf("%s%s_test.go", filename, loxe.GeneratedFileSuffix),
		packageName:    packageName,
		sourceCode:     "",
		privateImports: goImports.New(packagePath),
	}
}
