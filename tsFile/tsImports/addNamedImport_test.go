package tsImports

import "testing"

func TestAddNamedImport(t *testing.T) {
	t.Run("Should return an error if there's another defaultImport with the given namedImport", func(t *testing.T) {
		path := "path"
		namedImport := "namedImport"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{path, namedImport, nil})
		e := (&i).AddNamedImport(namedImport, path)
		if e == nil {
			t.Fatalf("Expected to be not nil")
		}
	})
	t.Run("Return nil error if there's an import with the given path that already contains the given namedImport", func(t *testing.T) {
		path := "path"
		namedImport := "namedImport"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{path, "", []string{namedImport}})
		e := (&i).AddNamedImport(namedImport, path)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if len(i) != 1 {
			t.Fatalf("Not expected to change the import list")
		}
		if len(i[0].namedImports) != 1 {
			t.Fatalf("Not expected to change the tsImport namedImports")
		}
	})
	t.Run("Should return an error if there's another namedImport with the given namedImport, for another path", func(t *testing.T) {
		path := "path"
		namedImport := "namedImport"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{path, "", []string{namedImport}})
		e := (&i).AddNamedImport(namedImport, "anotherPath")
		if e == nil {
			t.Fatalf("Expected to be not nil")
		}
	})
	t.Run("Should return nil error and use any already existing import with the same path", func(t *testing.T) {
		path := "path"
		namedImport := "namedImport"
		i := make(TsImports, 0, 1)
		tsI := &tsImport{path, "", []string{"anotherNamedImport"}}
		i = append(i, tsI)
		e := (&i).AddNamedImport(namedImport, path)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if len(i) != 1 {
			t.Fatalf("Not expected to change the import list size")
		}
		if len(tsI.namedImports) != 2 {
			t.Fatalf("Expected to add the new named import to the tsImport")
		}
		if tsI.namedImports[1] != namedImport {
			t.Fatalf("Wrong named import added")
		}
	})
	t.Run("Should add a new import to the list, if the given path isn't already present", func(t *testing.T) {
		path := "path"
		namedImport := "namedImport"
		i := make(TsImports, 0, 1)
		i = append(i, &tsImport{"anotherPath", "", []string{"anotherNamedImport"}})
		e := (&i).AddNamedImport(namedImport, path)
		if e != nil {
			t.Fatalf("Expected to be nil")
		}
		if len(i) != 2 {
			t.Fatalf("Expected to append the new named import to the import list")
		}
		if len(i[1].namedImports) != 1 {
			t.Fatalf("Newl created named imports must have just one named import")
		}
		if i[1].namedImports[0] != namedImport {
			t.Fatalf("Wrong named import appended")
		}
		if i[1].defaultImport != "" {
			t.Fatalf("Default import expected to be nil")
		}
		if i[1].path != path {
			t.Fatalf("Wrong import path")
		}
	})
}
