package parser

import "fmt"

// FocusPackagePath will tell the parser to look for a
// specific package
//
// Note that the packagePath argument refers to the import
// path to the target package, not the package name
func FocusPackagePath(packagePath string) *Focus {
	return &Focus{
		&packagePath,
		nil,
		nil,
	}
}

// FocusFilePath will tell the parser to look for a
// specific file, based on it's ABSOLUTE path
func FocusFilePath(filePath string) *Focus {
	return &Focus{
		nil,
		&filePath,
		nil,
	}
}

// FocusTypeName will tell the parser to look for a
// specific GO typename
func FocusTypeName(typeName string) *Focus {
	return &Focus{
		nil,
		nil,
		&typeName,
	}
}

// -----

// Focus tells to the parser
// what it needs to focus on
type Focus struct {
	// packagePath is the import
	// path to focused package.
	packagePath *string

	// filePath is the file system absolute
	// path to the focused file
	filePath *string

	// typeName is the name of the
	// focused GO typename
	typeName *string
}

// is is used to check if the parser focus is equal to the given one
func (f *Focus) is(lvl focusLevel, value string) bool {
	if f == nil {
		// If it's nil, there's no focus
		return true
	}

	// If the focus lvl equivalent is nil, then return true
	// because the focus is in something else

	switch lvl {
	case focusPackagePath:
		return f.packagePath == nil || *f.packagePath == value
	case focusFilePath:
		return f.filePath == nil || *f.filePath == value
	case focusTypeName:
		return f.typeName == nil || *f.typeName == value
	default:
		panic(fmt.Errorf("unrecognizable focus: %s", lvl))
	}
}

type focusLevel string

// focusLevel options
const (
	focusPackagePath  focusLevel = "packagePath"
	focusFilePath     focusLevel = "filePath"
	focusTypeName     focusLevel = "typeName"
)
