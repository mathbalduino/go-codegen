package goImports

import (
	"fmt"
	"testing"
)

func TestAddImport(t *testing.T) {
	t.Run("Should panic if the given alias is empty", func(t *testing.T) {
		i := &GoImports{"some/pkg/path", map[string]string{}}
		panicked := false
		panicMsg := ""
		c := make(chan bool)
		go func() {
			defer func() {
				e := recover()
				if e != nil {
					panicked = true
					panicMsg = e.(error).Error()
				}
				c <- true
			}()
			i.AddImport("   \n ", "some/another/path")
			c <- true
		}()

		<-c
		if !panicked {
			t.Fatalf("Expected to panic")
		}
		if panicMsg != emptyAliasError {
			t.Fatalf("Wrong panic msg")
		}
	})
	t.Run("Should panic if the given new import doesn't need to be imported", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		i := &GoImports{pkgPath, map[string]string{}}
		panicked := false
		c := make(chan bool)
		go func() {
			defer func() {
				e := recover()
				if e != nil {
					panicked = true
					if e.(error).Error() != fmt.Sprintf(addUnnecessaryImportError, pkgPath) {
						t.Fatalf("Wrong panic error")
					}
				}
				c <- true
			}()
			i.AddImport("someAlias", pkgPath)
			c <- true
		}()
		<-c
		if !panicked {
			t.Fatalf("Expected to panic")
		}
	})
	t.Run("If the import is duplicated, just return the actual alias", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		anotherPkgPath := "another/pkg"
		alias1 := "alias1"
		i := &GoImports{pkgPath, map[string]string{alias1: anotherPkgPath}}
		alias2 := i.AddImport("pkg", anotherPkgPath)
		if alias1 != alias2 {
			t.Fatalf("Returned alias expected to be equal")
		}
	})
	t.Run("Should use the suggested alias if there's no duplicate", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		i := &GoImports{pkgPath, map[string]string{"anotherAlias": "another/pkg"}}
		alias := "alias"
		if alias != i.AddImport(alias, "another/one") {
			t.Fatalf("Expected to return the suggested alias")
		}
	})
	t.Run("Should suffix the alias with a counter when it's duplicated", func(t *testing.T) {
		pkgPath := "some/pkg/path"
		alias := "alias"
		i := &GoImports{pkgPath, map[string]string{alias: "another/pkg"}}
		if i.AddImport(alias, "another/one/2") != alias+"_2" {
			t.Fatalf("Returned alias expected to have a suffix")
		}
		if i.AddImport(alias, "another/one/3") != alias+"_3" {
			t.Fatalf("Returned alias expected to have a suffix")
		}
		if i.AddImport(alias, "another/one/4") != alias+"_4" {
			t.Fatalf("Returned alias expected to have a suffix")
		}
	})
}
