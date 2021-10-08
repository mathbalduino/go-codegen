package parser

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestIterateFileInterfaces(t *testing.T) {
	fakeScopeObjects := func() (map[string]*ast.Object, *types.Info) {
		typeNameIdents := []*ast.Ident{{Name: "TypeName_0"}, {}, {Name: "TypeName_1"}, {}, {Name: "TypeName_2"}}
		typesInfo := &types.Info{
			Defs: map[*ast.Ident]types.Object{
				typeNameIdents[0]: types.NewTypeName(0, nil, typeNameIdents[0].Name, &types.Interface{}),
				typeNameIdents[1]: types.NewTypeName(0, nil, typeNameIdents[1].Name, &types.Struct{}),
				typeNameIdents[2]: types.NewTypeName(0, nil, typeNameIdents[2].Name, &types.Interface{}),
				typeNameIdents[3]: types.NewTypeName(0, nil, typeNameIdents[3].Name, &types.Struct{}),
				typeNameIdents[4]: types.NewTypeName(0, nil, typeNameIdents[4].Name, &types.Interface{}),
			},
		}
		return map[string]*ast.Object{
			"0": {Decl: &ast.TypeSpec{Name: typeNameIdents[0]}},
			"1": {Decl: &ast.TypeSpec{Name: typeNameIdents[1]}},
			"2": {Decl: &ast.TypeSpec{Name: typeNameIdents[2]}},
			"3": {Decl: &ast.TypeSpec{Name: typeNameIdents[3]}},
			"4": {Decl: &ast.TypeSpec{Name: typeNameIdents[4]}},
		}, typesInfo
	}
	fakeTypeNames := func() *GoParser {
		fileSet := token.NewFileSet()
		fileSet.AddFile("a", 1, 5)

		objects, typesInfo := fakeScopeObjects()
		return &GoParser{
			pkgs: []*packages.Package{{
				Syntax: []*ast.File{{
					Package: 2,
					Scope:   &ast.Scope{Objects: objects},
				}},
				TypesInfo: typesInfo,
			}},
			log:     emptyMockLogCLI(),
			fileSet: fileSet,
		}
	}
	t.Run("Should return nil errors when there are no Interfaces to iterate", func(t *testing.T) {
		p := fakeTypeNames()
		// Remove the interfaces from the Objects map (see fakeScopeObjects above)
		delete(p.pkgs[0].Syntax[0].Scope.Objects, "0")
		delete(p.pkgs[0].Syntax[0].Scope.Objects, "2")
		delete(p.pkgs[0].Syntax[0].Scope.Objects, "4")

		e := p.IterateFileInterfaces(func(type_ *types.TypeName, parentLog LogCLI) error { return nil })
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
	})
	t.Run("Should return errors returned by the callback", func(t *testing.T) {
		p := fakeTypeNames()
		e := p.IterateFileInterfaces(func(type_ *types.TypeName, parentLog LogCLI) error {
			return fmt.Errorf("error")
		})
		if e == nil {
			t.Fatalf("Expected to be not nil")
		}
	})
	t.Run("Skip everything that is not a Interface", func(t *testing.T) {
		p := fakeTypeNames()
		calls := 0
		e := p.IterateFileInterfaces(func(type_ *types.TypeName, parentLog LogCLI) error {
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
	t.Run("Should call the callback for every Interface that needs to be iterated", func(t *testing.T) {
		p := fakeTypeNames()
		callsA, callsB, callsC := 0, 0, 0
		e := p.IterateFileInterfaces(func(type_ *types.TypeName, parentLog LogCLI) error {
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
				t.Fatalf("Unexpected Interface callback call")
				return nil
			}
		})
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if callsA != 1 || callsB != 1 || callsC != 1 {
			t.Fatalf("Callback was expected to be called one time for every Interface")
		}
	})
}
