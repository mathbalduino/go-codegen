package goParser

import "testing"

func TestParserFocus(t *testing.T) {
	t.Run("FocusPackagePath should fill only the packagePath field", func(t *testing.T) {
		pkgPath := "packagePath"
		f := FocusPackagePath(pkgPath)
		if f == nil {
			t.Fatalf("Focus not expected to be nil")
		}
		if f.packagePath == nil {
			t.Fatalf("Focus.packagePath not expected to be nil")
		}
		if *f.packagePath != pkgPath {
			t.Fatalf("Focus.packagePath expected to be equal to the given argument")
		}
		if f.filePath != nil {
			t.Fatalf("Focus.filePath expected to be nil")
		}
		if f.typeName != nil {
			t.Fatalf("Focus.typeName expected to be nil")
		}
		if f.functionName != nil {
			t.Fatalf("Focus.functionName expected to be nil")
		}
		if f.varName != nil {
			t.Fatalf("Focus.varName expected to be nil")
		}
	})
	t.Run("FocusFilePath should fill only the filePath field", func(t *testing.T) {
		filePath := "filePath"
		f := FocusFilePath(filePath)
		if f == nil {
			t.Fatalf("Focus not expected to be nil")
		}
		if f.filePath == nil {
			t.Fatalf("Focus.filePath not expected to be nil")
		}
		if *f.filePath != filePath {
			t.Fatalf("Focus.filePath expected to be equal to the given argument")
		}
		if f.packagePath != nil {
			t.Fatalf("Focus.packagePath expected to be nil")
		}
		if f.typeName != nil {
			t.Fatalf("Focus.typeName expected to be nil")
		}
		if f.functionName != nil {
			t.Fatalf("Focus.functionName expected to be nil")
		}
		if f.varName != nil {
			t.Fatalf("Focus.varName expected to be nil")
		}
	})
	t.Run("FocusFunctionName should fill only the functionName field", func(t *testing.T) {
		functionName := "functionName"
		f := FocusFunctionName(functionName)
		if f == nil {
			t.Fatalf("Focus not expected to be nil")
		}
		if f.functionName == nil {
			t.Fatalf("Focus.functionName not expected to be nil")
		}
		if *f.functionName != functionName {
			t.Fatalf("Focus.functionName expected to be equal to the given argument")
		}
		if f.packagePath != nil {
			t.Fatalf("Focus.packagePath expected to be nil")
		}
		if f.typeName != nil {
			t.Fatalf("Focus.typeName expected to be nil")
		}
		if f.filePath != nil {
			t.Fatalf("Focus.filePath expected to be nil")
		}
		if f.varName != nil {
			t.Fatalf("Focus.varName expected to be nil")
		}
	})
	t.Run("FocusTypeName should fill only the typeName field", func(t *testing.T) {
		typeName := "typeName"
		f := FocusTypeName(typeName)
		if f == nil {
			t.Fatalf("Focus not expected to be nil")
		}
		if f.typeName == nil {
			t.Fatalf("Focus.typeName not expected to be nil")
		}
		if *f.typeName != typeName {
			t.Fatalf("Focus.typeName expected to be equal to the given argument")
		}
		if f.packagePath != nil {
			t.Fatalf("Focus.packagePath expected to be nil")
		}
		if f.functionName != nil {
			t.Fatalf("Focus.functionName expected to be nil")
		}
		if f.filePath != nil {
			t.Fatalf("Focus.filePath expected to be nil")
		}
		if f.varName != nil {
			t.Fatalf("Focus.varName expected to be nil")
		}
	})
	t.Run("FocusVarName should fill only the varName field", func(t *testing.T) {
		varName := "varName"
		f := FocusVarName(varName)
		if f == nil {
			t.Fatalf("Focus not expected to be nil")
		}
		if f.varName == nil {
			t.Fatalf("Focus.varName not expected to be nil")
		}
		if *f.varName != varName {
			t.Fatalf("Focus.varName expected to be equal to the given argument")
		}
		if f.packagePath != nil {
			t.Fatalf("Focus.packagePath expected to be nil")
		}
		if f.functionName != nil {
			t.Fatalf("Focus.functionName expected to be nil")
		}
		if f.filePath != nil {
			t.Fatalf("Focus.filePath expected to be nil")
		}
		if f.typeName != nil {
			t.Fatalf("Focus.typeName expected to be nil")
		}
	})
	t.Run("ParserFocus.is should return true if the receiver is nil", func(t *testing.T) {
		p := (*ParserFocus)(nil)
		if !p.is("", "") {
			t.Fatalf("Expected to be true")
		}
	})
}
