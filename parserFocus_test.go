package parser

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
		if f.filePath != nil {
			t.Fatalf("Focus.filePath expected to be nil")
		}
	})
	t.Run("Focus.is should return true if the receiver is nil", func(t *testing.T) {
		p := (*Focus)(nil)
		if !p.is("", "") {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("When PackagePaths are being focused, the other focuses should be ignored", func(t *testing.T) {
		str := "focusStr"
		p := &Focus{
			packagePath: &str,
		}
		if p.is(focusPackagePath, "---") {
			t.Fatalf("Expected to be false")
		}
		if !p.is(focusPackagePath, str) {
			t.Fatalf("Expected to be true")
		}

		// This focus is in packagePaths, so the other comparisons should always return true
		if !p.is(focusFilePath, str) || !p.is(focusTypeName, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("When FilePaths are being focused, the other focuses should be ignored", func(t *testing.T) {
		str := "focusStr"
		p := &Focus{
			filePath: &str,
		}
		if p.is(focusFilePath, "---") {
			t.Fatalf("Expected to be false")
		}
		if !p.is(focusFilePath, str) {
			t.Fatalf("Expected to be true")
		}

		// This focus is in packagePaths, so the other comparisons should always return true
		if !p.is(focusPackagePath, str) || !p.is(focusTypeName, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("When TypeNames are being focused, the other focuses should be ignored", func(t *testing.T) {
		str := "focusStr"
		p := &Focus{
			typeName: &str,
		}
		if p.is(focusTypeName, "---") {
			t.Fatalf("Expected to be false")
		}
		if !p.is(focusTypeName, str) {
			t.Fatalf("Expected to be true")
		}

		// This focus is in packagePaths, so the other comparisons should always return true
		if !p.is(focusPackagePath, str) || !p.is(focusFilePath, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Expects panic when the given focus is not recognizable", func(t *testing.T) {
		p := &Focus{}
		c := make(chan bool)
		go func() {
			defer func() {
				if e := recover(); e == nil {
					t.Logf("It was expected to panic")
					t.Fail()
				}
				c <- true
			}()

			p.is("unrecognizable focus", "any value")
		}()
		<-c
	})
}
