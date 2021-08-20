package tsImports

import "fmt"

// AddNamedImport will add a new named import to the given path
//
// If there's another default/named import with the name equal to the
// given namedImport, in the entire list of imports, an error will be
// returned, since an import name must be unique
//
// If there's no error (return value == nil), the given namedImport will
// never be changed inside the imports
func (i *TsImports) AddNamedImport(namedImport string, path string) error {
	var targetPath *tsImport
	for _, currImport := range *i {
		if currImport.path == path {
			targetPath = currImport
		}

		if currImport.defaultImport == namedImport {
			return fmt.Errorf("the default import of the path '%s' is already using the '%s' name",
				currImport.path, namedImport)
		}

		for _, currNamedImport := range currImport.namedImports {
			if currNamedImport == namedImport {
				if currImport.path == path {
					return nil
				}

				return fmt.Errorf("there's already a named import from the '%s' path using the '%s' name",
					currImport.path, namedImport)
			}
		}
	}

	if targetPath != nil {
		targetPath.namedImports = append(targetPath.namedImports, namedImport)
		return nil
	}

	*i = append(*i, &tsImport{
		path:          path,
		defaultImport: "",
		namedImports:  []string{namedImport},
	})
	return nil
}
