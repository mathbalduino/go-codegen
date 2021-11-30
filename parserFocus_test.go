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
	t.Run("If the focus is not on PackagePath, then it should always return true", func(t *testing.T) {
		str := "focusStr"
		p := &Focus{filePath: &str}
		if !p.is(focusPackagePath, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("If the packagePath is set, should check the regexp", func(t *testing.T) {
		str := "focusStr"
		strRegexp := "^" + str + "$"
		p := &Focus{packagePath: &strRegexp}
		if p.is(focusPackagePath, "anythingElse") {
			t.Fatalf("Expected to be false")
		}
		if p.is(focusPackagePath, str+"someSuffix") {
			t.Fatalf("Expected to be false")
		}
		if p.is(focusPackagePath, "somePrefix"+str) {
			t.Fatalf("Expected to be false")
		}
		if !p.is(focusPackagePath, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("If the focus is not on FilePath, then it should always return true", func(t *testing.T) {
		str := "focusStr"
		p := &Focus{packagePath: &str}
		if !p.is(focusFilePath, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("If the filePath is set, should check the regexp", func(t *testing.T) {
		str := "focusStr"
		strRegexp := "^" + str + "$"
		p := &Focus{filePath: &strRegexp}
		if p.is(focusFilePath, "anythingElse") {
			t.Fatalf("Expected to be false")
		}
		if p.is(focusFilePath, str+"someSuffix") {
			t.Fatalf("Expected to be false")
		}
		if p.is(focusFilePath, "somePrefix"+str) {
			t.Fatalf("Expected to be false")
		}
		if !p.is(focusFilePath, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("If the focus is not on TypeName, then it should always return true", func(t *testing.T) {
		str := "focusStr"
		p := &Focus{packagePath: &str}
		if !p.is(focusTypeName, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("If the typeName is set, should check the regexp", func(t *testing.T) {
		str := "focusStr"
		strRegexp := "^" + str + "$"
		p := &Focus{typeName: &strRegexp}
		if p.is(focusTypeName, "anythingElse") {
			t.Fatalf("Expected to be false")
		}
		if p.is(focusTypeName, str+"someSuffix") {
			t.Fatalf("Expected to be false")
		}
		if p.is(focusTypeName, "somePrefix"+str) {
			t.Fatalf("Expected to be false")
		}
		if !p.is(focusTypeName, str) {
			t.Fatalf("Expected to be true")
		}
	})
	t.Run("Expects panic when the given focus is not recognizable", func(t *testing.T) {
		p := &Focus{}
		c := make(chan bool)
		go func() {
			defer func() {
				if e := recover(); e == nil {
					c <- false
				} else {
					c <- true
				}
			}()
			p.is("unrecognizable focus", "any value")
		}()
		if !<-c {
			t.Logf("It was expected to panic")
			t.Fail()
		}
	})
	t.Run("MergeFocus should copy all the values from the second argument, overriding the first", func(t *testing.T) {
		str1, str2 := "str_one", "str_two"
		f1 := &Focus{
			packagePath: &str1,
			filePath:    &str1,
			typeName:    &str1,
		}
		t.Run("Override packagePath only", func(t *testing.T) {
			f2 := &Focus{packagePath: &str2}
			f3 := MergeFocus(f1, f2)
			if f3.packagePath != &str2 {
				t.Fatalf("packagePath was expected to be overriden")
			}
			if f3.filePath != &str1 {
				t.Fatalf("filePath was not expected to change")
			}
			if f3.typeName != &str1 {
				t.Fatalf("typeName was not expected to change")
			}
		})
		t.Run("Override filePath only", func(t *testing.T) {
			f2 := &Focus{filePath: &str2}
			f3 := MergeFocus(f1, f2)
			if f3.packagePath != &str1 {
				t.Fatalf("packagePath was not expected to change")
			}
			if f3.filePath != &str2 {
				t.Fatalf("filePath was expected to be overriden")
			}
			if f3.typeName != &str1 {
				t.Fatalf("typeName was not expected to change")
			}
		})
		t.Run("Override typeName only", func(t *testing.T) {
			f2 := &Focus{typeName: &str2}
			f3 := MergeFocus(f1, f2)
			if f3.packagePath != &str1 {
				t.Fatalf("packagePath was not expected to change")
			}
			if f3.filePath != &str1 {
				t.Fatalf("filePath was not expected to change")
			}
			if f3.typeName != &str2 {
				t.Fatalf("typeName was expected to be overriden")
			}
		})
		t.Run("Override all fields", func(t *testing.T) {
			f2 := &Focus{
				packagePath: &str2,
				filePath:    &str2,
				typeName:    &str2,
			}
			f3 := MergeFocus(f1, f2)
			if f3.packagePath != &str2 {
				t.Fatalf("packagePath was expected to be overriden")
			}
			if f3.filePath != &str2 {
				t.Fatalf("filePath was expected to be overriden")
			}
			if f3.typeName != &str2 {
				t.Fatalf("typeName was expected to be overriden")
			}
		})
	})
}
