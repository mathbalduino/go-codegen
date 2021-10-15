package goFile

// PackageName will return the name of
// the package that the file belongs to
func (f *GoFile) PackageName() string {
	return f.packageName
}
