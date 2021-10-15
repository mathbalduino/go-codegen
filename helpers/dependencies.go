package helpers

// GoImports just to ease tests
type GoImports interface {
	NeedImport(string) bool
	AliasFromPath(string) string
}
