package parser

import (
	"fmt"
	"github.com/mathbalduino/go-log/loggerCLI"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestIteratePackages(t *testing.T) {
	t.Run("Should use the optionalLogger, if provided, giving it to the callback", func(t *testing.T) {
		mock := &mockLoggerCLI{}
		mock.mockTrace = func(string, ...interface{}) LoggerCLI { return mock }
		p := &GoParser{pkgs: []*packages.Package{{}}, logger: loggerCLI.New(false, 0)}
		e := p.iteratePackages(func(_ *packages.Package, logger LoggerCLI) error {
			m, isMock := logger.(*mockLoggerCLI)
			if !isMock || m != mock {
				t.Fatalf("The LoggerCLI given to the callback is not the expected one")
			}
			return nil
		}, mock)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should return nil when there are no packages", func(t *testing.T) {
		p := &GoParser{pkgs: nil, logger: loggerCLI.New(false, 0)}
		e := p.iteratePackages(nil)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should skip packages with error", func(t *testing.T) {
		m := loggerCLI.New(false, 0)
		okPkg := "c"
		p := &GoParser{
			pkgs: []*packages.Package{
				{PkgPath: "a", Errors: []packages.Error{{}, {}}},
				{PkgPath: "b", Errors: []packages.Error{{}}},
				{PkgPath: okPkg, Errors: nil},
				{PkgPath: "d", Errors: []packages.Error{{}}},
			},
			logger: m,
		}
		calls := 0
		e := p.iteratePackages(func(pkg *packages.Package, parentLog LoggerCLI) error {
			calls += 1
			if pkg.PkgPath != okPkg {
				t.Fatalf("Packages with error should be skipped")
			}
			return nil
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if calls != 1 {
			t.Fatalf("Expected to iterate over just one pkg")
		}
	})
	t.Run("Should stop iteration when the callback returns any error, forwarding the returned error", func(t *testing.T) {
		m := loggerCLI.New(false, 0)
		p := &GoParser{
			pkgs: []*packages.Package{
				{Errors: nil},
				{Errors: nil},
			},
			logger: m,
		}
		calls := 0
		e := p.iteratePackages(func(pkg *packages.Package, parentLog LoggerCLI) error {
			calls += 1
			return fmt.Errorf("any")
		})
		if e == nil {
			t.Fatalf("Expected to be not nil")
		}
		if calls != 1 {
			t.Fatalf("Expected to stop at the first callback returned error")
		}
	})
	t.Run("Should skip any package that is not the focus", func(t *testing.T) {
		m := loggerCLI.New(false, 0)
		focusPkg := "focusedPkg"
		p := &GoParser{
			pkgs: []*packages.Package{
				{PkgPath: "a", Errors: nil},
				{PkgPath: focusPkg, Errors: nil},
				{PkgPath: "b", Errors: nil},
			},
			logger: m,
			focus:  FocusPackagePath(focusPkg),
		}
		calls := 0
		e := p.iteratePackages(func(pkg *packages.Package, parentLog LoggerCLI) error {
			calls += 1
			if pkg.PkgPath != focusPkg {
				t.Fatalf("Expected to iterate only over the focused package")
			}
			return nil
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if calls != 1 {
			t.Fatalf("Expected to iterate only over the focused package")
		}
	})
	t.Run("Should call the callback for every package that needs to be iterated and return nil error", func(t *testing.T) {
		m := loggerCLI.New(false, 0)
		p := &GoParser{
			pkgs: []*packages.Package{
				{PkgPath: "a", Errors: nil},
				{PkgPath: "b", Errors: nil},
				{PkgPath: "c", Errors: nil},
			},
			logger: m,
		}
		aCalls, bCalls, cCalls := 0, 0, 0
		e := p.iteratePackages(func(pkg *packages.Package, parentLog LoggerCLI) error {
			switch pkg.PkgPath {
			case "a":
				aCalls += 1
				return nil
			case "b":
				bCalls += 1
				return nil
			case "c":
				cCalls += 1
				return nil
			default:
				t.Fatalf("Unexpected package iteration")
				return nil
			}
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if aCalls != 1 || bCalls != 1 || cCalls != 1 {
			t.Fatalf("Each package must be iterated one time")
		}
	})
}
