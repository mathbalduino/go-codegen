package goImports

import "fmt"

// AddImport will take some alias and package path and try to add it to the
// list of imports.
//
// The given alias is just a suggestion, it can be changed if there's a clash
// with another import
//
// If the package path is already present in the import list, the already
// existing alias will be returned
func (i *GoImports) AddImport(suggestedAlias, packagePath string) string {
	if !i.NeedImport(packagePath) {
		panic(fmt.Errorf(addUnnecessaryImportError, packagePath))
	}

	alreadyExistentAlias := i.AliasFromPath(packagePath)
	if alreadyExistentAlias != "" {
		return alreadyExistentAlias
	}

	n := 2
	possibleAlias := suggestedAlias
	for {
		_, alreadyExists := i.imports[possibleAlias]
		if !alreadyExists {
			break
		}

		possibleAlias = fmt.Sprintf("%s_%d", suggestedAlias, n)
		n += 1
	}

	i.imports[possibleAlias] = packagePath
	return possibleAlias
}

const addUnnecessaryImportError = "trying to add unnecessary import: %s"
