package goImports

// AliasFromPath will take some package path as argument
// and return the alias associated with it
//
// Note that if, and only if, there's no package with the
// given path in the import list, the returned alias will
// be empty
func (i *GoImports) AliasFromPath(packagePath string) string {
	for alias, existentPath := range i.imports {
		if existentPath == packagePath {
			return alias
		}
	}

	// There's no import with the given packagePath
	return ""
}
