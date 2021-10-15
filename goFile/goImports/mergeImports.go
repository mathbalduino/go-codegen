package goImports

// MergeImports will merge the given imports list with the method
// receiver *GoImports
//
// When merging, some aliases maybe change due to clashes. All renamed
// aliases will be returned in a map of oldAlias -> newAlias
func (i *GoImports) MergeImports(other *GoImports) map[string]string {
	changedAliases := map[string]string{}
	for suggestedAlias, packagePath := range other.imports {
		realAlias := i.AddImport(suggestedAlias, packagePath)
		if realAlias != suggestedAlias {
			changedAliases[suggestedAlias] = realAlias
		}
	}

	return changedAliases
}
