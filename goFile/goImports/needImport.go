package goImports

// NeedImport will return true if the given package path needs
// to be imported, when called inside the import list host package
func (i *GoImports) NeedImport(otherPackagePath string) bool {
	return i.packagePath != otherPackagePath
}
