package tsFile

// AddCode will just append the given sourceCode to
// the file.
//
// Note that a "\n" will be added as a prefix to the
// given code
func (f *TsFile) AddCode(newSourceCode string) {
	f.sourceCode += "\n" + newSourceCode
}
