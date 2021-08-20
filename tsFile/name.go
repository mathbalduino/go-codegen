package tsFile

// Name will return the name of the file,
// without the folderpath
func (f *TsFile) Name() string {
	return f.name
}
