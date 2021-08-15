package helpers

type GoImports interface {
	NeedImport(string) bool
	AliasFromPath(string) string
}
