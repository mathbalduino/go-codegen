package parser

import (
	"go/token"
	"testing"
)

func TestNewGoParser(t *testing.T) {
	t.Run("Should return errors of packages.Load", func(t *testing.T) {
		p, e := NewGoParser("wrongpattern=", Config{})
		if e == nil {
			t.Fatalf("The packages.Load error was expected to be returned")
		}
		if p != nil {
			t.Fatalf("When there's errors, the GoParser should be nil")
		}
	})
	t.Run("Should return a valid GoParser and nil error", func(t *testing.T) {
		p, e := NewGoParser("--inexistentPackage--", Config{})
		if e != nil {
			t.Fatalf("The error was expected to be nil")
		}
		if p == nil {
			t.Fatalf("GoParser expected to be not nil")
		}
	})
	t.Run("The returned GoParser pkgs, focus, log and fileSet should be filled", func(t *testing.T) {
		config := Config{Focus: &Focus{}, Fset: token.NewFileSet()}
		p, _ := NewGoParser("--inexistentPackage--", config)
		if p == nil {
			t.Fatalf("GoParser expected to be not nil")
		}
		if p.pkgs == nil {
			t.Fatalf("GoParser.pkgs expected to be not nil")
		}
		if p.focus != config.Focus {
			t.Fatalf("GoParser.focus expected to be equal to config.Focus")
		}
		if p.logger == nil {
			t.Fatalf("GoParser.log expected to not be nil")
		}
		if p.fileSet != config.Fset {
			t.Fatalf("GoParser.fileSet expected to be equal to config.Fset")
		}
	})
}
