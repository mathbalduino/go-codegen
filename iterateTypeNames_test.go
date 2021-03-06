package parser

import (
	"fmt"
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestIterateTypeNames(t *testing.T) {
	fakeScopeObjects := func(typeNameFocus *string) (map[string]*ast.Object, *types.Info) {
		focusedOne := "TypeName_1"
		if typeNameFocus != nil {
			focusedOne = *typeNameFocus
		}
		typeNameIdents := []*ast.Ident{{Name: "TypeName_0"}, {Name: focusedOne}, {Name: "TypeName_2"}, {}}
		typesInfo := &types.Info{
			Defs: map[*ast.Ident]types.Object{
				typeNameIdents[0]: types.NewTypeName(0, nil, typeNameIdents[0].Name, nil),
				typeNameIdents[1]: types.NewTypeName(0, nil, typeNameIdents[1].Name, nil),
				typeNameIdents[2]: types.NewTypeName(0, nil, typeNameIdents[2].Name, nil),
				typeNameIdents[3]: types.NewVar(0, nil, "anyOtherThing", nil),
			},
		}
		return map[string]*ast.Object{
			"0":   {Decl: &ast.TypeSpec{Name: typeNameIdents[0]}},
			"-":   {Decl: "anyOtherThing"},
			"1":   {Decl: &ast.TypeSpec{Name: typeNameIdents[1]}},
			"--":  {Decl: &ast.TypeSpec{Name: &ast.Ident{ /* not in typesInfo */ }}},
			"---": {Decl: &ast.TypeSpec{Name: typeNameIdents[3]}},
			"2":   {Decl: &ast.TypeSpec{Name: typeNameIdents[2]}},
		}, typesInfo
	}
	fakeTypeNames := func(typeNameFocus *string) *GoParser {
		fileSet := token.NewFileSet()
		fileSet.AddFile("a", 1, 5)

		objects, typesInfo := fakeScopeObjects(typeNameFocus)
		focus := (*Focus)(nil)
		if typeNameFocus != nil {
			focus = FocusTypeName(*typeNameFocus)
		}
		return &GoParser{
			pkgs: []*packages.Package{{
				Syntax: []*ast.File{{
					Package: 2,
					Scope:   &ast.Scope{Objects: objects},
				}},
				TypesInfo: typesInfo,
			}},
			logger:  loggerCLI.New(false, 0),
			fileSet: fileSet,
			focus:   focus,
		}
	}

	t.Run("Should forward the optionalLogger, if provided, to the iteratePackageFiles method (this test depends on the iteratePackageFiles forwarding its own optionalLogger to its callback)", func(t *testing.T) {
		p := fakeTypeNames(nil)
		mock := &mockLoggerCLI{}
		mock.mockTrace = func(string, ...interface{}) LoggerCLI { return mock }
		e := p.iterateTypeNames(func(type_ *types.TypeName, logger LoggerCLI) error {
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
	t.Run("Should return nil errors when there are no Scope.Objects to iterate", func(t *testing.T) {
		p := fakeTypeNames(nil)
		p.pkgs[0].Syntax[0].Scope.Objects = map[string]*ast.Object{}
		e := p.iterateTypeNames(func(type_ *types.TypeName, parentLog LoggerCLI) error { return nil })
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should return errors returned by the callback", func(t *testing.T) {
		p := fakeTypeNames(nil)
		e := p.iterateTypeNames(func(type_ *types.TypeName, parentLog LoggerCLI) error {
			return fmt.Errorf("error")
		})
		if e == nil {
			t.Fatalf("Expected to be not nil")
		}
	})
	t.Run("Skip everything that is not a TypeName", func(t *testing.T) {
		p := fakeTypeNames(nil)
		calls := 0
		e := p.iterateTypeNames(func(type_ *types.TypeName, parentLog LoggerCLI) error {
			calls += 1
			return nil
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if calls != 3 {
			t.Fatalf("Callback was expected to be called 3 times")
		}
	})
	t.Run("Iterate only over the focused TypeName", func(t *testing.T) {
		focus := "focus"
		p := fakeTypeNames(&focus)
		calls := 0
		e := p.iterateTypeNames(func(type_ *types.TypeName, parentLog LoggerCLI) error {
			if type_.Name() != focus {
				t.Fatalf("Callback was not expected to be called with non focused TypeNames")
				return nil
			}
			calls += 1
			return nil
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if calls != 1 {
			t.Fatalf("Callback was expected to be called 3 times")
		}
	})
	t.Run("Should call the callback for every TypeName that needs to be iterated", func(t *testing.T) {
		p := fakeTypeNames(nil)
		callsA, callsB, callsC := 0, 0, 0
		e := p.iterateTypeNames(func(type_ *types.TypeName, parentLog LoggerCLI) error {
			switch type_.Name() {
			// Strings from "fakeScopeObjects"
			case "TypeName_0":
				callsA += 1
				return nil
			case "TypeName_1":
				callsB += 1
				return nil
			case "TypeName_2":
				callsC += 1
				return nil
			default:
				t.Fatalf("Unexpected TypeName callback call")
				return nil
			}
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if callsA != 1 || callsB != 1 || callsC != 1 {
			t.Fatalf("Callback was expected to be called one time for every TypeName")
		}
	})
}
