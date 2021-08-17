package goImports

// GoImports holds information about
// the host package and the imports
// that exist inside it
type GoImports struct {
	// packagePath is the path to the
	// host package
	packagePath string

	// imports is the list of imports
	// that are registered.
	// It is a map of alias -> path
	imports map[string]string
}

// New will create a new import list to a GO file.
//
// Note that the packagePath param is the path to the host
// package, and can be used to calculate the other imports
func New(packagePath string) *GoImports {
	return &GoImports{
		packagePath: packagePath,
		imports:     map[string]string{},
	}
}
