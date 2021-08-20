package tsImports

// tsImport represents a single Typescript - TS
// import line to a single path
type tsImport struct {
	// path is the TS import path of the file being
	// imported
	path string

	// defaultImport, if present, represents a default
	// import, that must be unique in a single import line
	defaultImport string

	// namedImports is a collection of named TS imports
	namedImports []string
}

// TsImports is a collection of single
// Typescript - TS line imports
type TsImports []*tsImport

// New will just return a pointer to
// an empty TsImports
func New() *TsImports {
	t := make(TsImports, 0, 5)
	return &t
}
