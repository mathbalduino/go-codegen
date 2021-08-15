package helpers

func emptyMockGoImports() *mockGoImports {
	return &mockGoImports{
		func(path string) bool { return false },
		func(path string) string { return "" },
	}
}

type mockGoImports struct {
	mockNeedImport    func(path string) bool
	mockAliasFromPath func(path string) string
}

func (m *mockGoImports) NeedImport(path string) bool {
	if m.mockNeedImport != nil {
		return m.mockNeedImport(path)
	}
	return false
}

func (m *mockGoImports) AliasFromPath(path string) string {
	if m.mockAliasFromPath != nil {
		return m.mockAliasFromPath(path)
	}
	return ""
}
