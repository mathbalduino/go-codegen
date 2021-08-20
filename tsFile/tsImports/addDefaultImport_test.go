package tsImports

import "testing"

func TestAddDefaultImport(t *testing.T) {
	t.Run("Should return nil if there's another import with the given defaultImport name and path", func(t *testing.T) {
		path := "path"
		defaultImport := "defaultImportName"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{path, defaultImport, nil})
		e := (&i).AddDefaultImport(defaultImport, path)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if len(i) != 1 {
			t.Fatalf("Not expected to change the import list")
		}
		if i[0].defaultImport != defaultImport {
			t.Fatalf("Wrong defaultImport")
		}
	})
	t.Run("Should return error if there's another import wih the given defaultImport name and different path", func(t *testing.T) {
		path := "path"
		defaultImport := "defaultImportName"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{path, defaultImport, nil})
		e := (&i).AddDefaultImport(defaultImport, "anotherPath")
		if e == nil {
			t.Fatalf("Expected an error")
		}
	})
	t.Run("Should return error if there's another named import with the given defaultImport", func(t *testing.T) {
		defaultImport := "defaultImportName"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{"path", "", []string{defaultImport}})
		e := (&i).AddDefaultImport(defaultImport, "anotherPath")
		if e == nil {
			t.Fatalf("Expected an error")
		}
	})
	t.Run("If there's another import with the same path, add the given defaultImport to it", func(t *testing.T) {
		path := "path"
		defaultImport := "defaultImportName"
		i := make(TsImports, 0, 1)
		tsI := &tsImport{path, "", nil}
		i = append(i, tsI)
		e := (&i).AddDefaultImport(defaultImport, path)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if len(i) != 1 {
			t.Fatalf("Not expected to change the import list size")
		}
		if tsI.defaultImport != defaultImport {
			t.Fatalf("Expected to set the existent tsImport.defaultImport")
		}
	})
	t.Run("If there isn't another import with the same path, append a new tsImport to the list", func(t *testing.T) {
		path := "newPath"
		defaultImport := "defaultImportName"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{"path", "", nil})
		e := (&i).AddDefaultImport(defaultImport, path)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if len(i) != 2 {
			t.Fatalf("Expected to append a new tsImport")
		}
		if i[1].defaultImport != defaultImport {
			t.Fatalf("New appended import defaultImport is wrong")
		}
		if i[1].path != path {
			t.Fatalf("New appended import path is wrong")
		}
	})
}
