package goFile

import "testing"

func TestName(t *testing.T) {
	t.Run("Should return the name of the file, acting as a getter", func(t *testing.T) {
		fileName := "fileName"
		f := &GoFile{name: fileName}
		if f.Name() != fileName {
			t.Fatalf("Wrong fileName returned")
		}
	})
}
