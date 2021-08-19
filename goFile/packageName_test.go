package goFile

import "testing"

func TestPackageName(t *testing.T) {
	t.Run("Should return the file package name, acting as a getter", func(t *testing.T) {
		pkgName := "pkgName"
		f := &GoFile{packageName: pkgName}
		if f.PackageName() != pkgName {
			t.Fatalf("Wrong package name returned")
		}
	})
}
