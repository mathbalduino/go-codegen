package goParser

import "fmt"

// FocusPackagePath will tell the parser to look for a
// specific package
//
// Note that the packagePath argument refers to the import
// path to the target package, not the package name
func FocusPackagePath(packagePath string) *ParserFocus {
	return &ParserFocus{
		&packagePath,
		nil,
		nil,
		nil,
		nil,
	}
}

// FocusFilePath will tell the parser to look for a
// specific file, based on it's ABSOLUTE path
func FocusFilePath(filePath string) *ParserFocus {
	return &ParserFocus{
		nil,
		&filePath,
		nil,
		nil,
		nil,
	}
}

// FocusTypeName will tell the parser to look for a
// specific GO type name
func FocusTypeName(typeName string) *ParserFocus {
	return &ParserFocus{
		nil,
		nil,
		&typeName,
		nil,
		nil,
	}
}

// FocusVarName will tell the parser to look for a
// specific GO variable name
func FocusVarName(varName string) *ParserFocus {
	return &ParserFocus{
		packagePath:  nil,
		filePath:     nil,
		typeName:     nil,
		varName:      &varName,
		functionName: nil,
	}
}

// FocusFunctionName will tell the parser to look for a
// specific GO function name
func FocusFunctionName(functionName string) *ParserFocus {
	return &ParserFocus{
		packagePath:  nil,
		filePath:     nil,
		typeName:     nil,
		varName:      nil,
		functionName: &functionName,
	}
}

// -----

// ParserFocus tells to the parser
// what it needs to focus on
type ParserFocus struct {
	// packagePath is the import
	// path to focused package.
	packagePath *string

	// filePath is the file system path
	// to the focused file
	filePath *string

	// typeName is the name of a GO
	// type that is the focus
	typeName *string

	// varName is the name of a GO
	// variable that is the focus
	varName *string

	// functionName is the name of a GO
	// function that is the focus
	functionName *string
}

// is is used to check if the focus is equal to the given one
func (f *ParserFocus) is(lvl focusLevel, value string) bool {
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
	case focusVarName:
		return f.varName == nil || *f.varName == value
	case focusFunctionName:
		return f.functionName == nil || *f.functionName == value
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
	focusVarName      focusLevel = "varName"
	focusFunctionName focusLevel = "functionName"
)
