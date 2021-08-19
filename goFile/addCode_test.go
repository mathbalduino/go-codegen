package goFile

import "testing"

func TestAddCode(t *testing.T) {
	t.Run("Should add a break line before appending the given source code", func(t *testing.T) {
		f := &GoFile{}
		code := "My fictional source code"
		f.AddCode(code)
		if f.sourceCode != "\n" + code {
			t.Fatalf("Wrong file source code")
		}
		f.AddCode(code)
		if f.sourceCode != "\n" + code + "\n" + code {
			t.Fatalf("Wrong file source code")
		}
	})
}
