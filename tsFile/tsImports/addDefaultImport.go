package tsImports

import "fmt"

// AddDefaultImport will add a new default import for the given path
//
// If the given path already has a default import, and it is not equal
// the given defaultImport param, an error will be returned, since a path
// can not contain more than one default import
//
// If there's another default/named import with the name equal to the
// given defaultImport param, in the entire list of imports, an error will be
// returned, since an import name must be unique
//
// If there's no error (return value == nil), the given defaultImport will
// never be changed inside the imports
func (i *TsImports) AddDefaultImport(defaultImport string, path string) error {
	var targetPath *tsImport
	for _, currImport := range *i {
		if currImport.path == path {
			targetPath = currImport
		}

		if currImport.defaultImport == defaultImport {
			if currImport.path == path {
				return nil
			}

			return fmt.Errorf("the default import of the path '%s' is already using the '%s' name",
				currImport.path, defaultImport)
		}

		for _, currNamedImport := range currImport.namedImports {
			if currNamedImport == defaultImport {
				return fmt.Errorf("there's already a named import from the '%s' path using the '%s' name",
					currImport.path, defaultImport)
			}
		}
	}

	if targetPath != nil {
		targetPath.defaultImport = defaultImport
		return nil
	}

	*i = append(*i, &tsImport{
		path:          path,
		defaultImport: defaultImport,
		namedImports:  nil,
	})
	return nil
}
