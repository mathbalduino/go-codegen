package goFile

// AddCode will append the given string to the file
// source code body (below imports)
//
// Note that the given string is expected to be valid
// go code. This method doesn't do any checks
func (f *GoFile) AddCode(newSourceCode string) {
	f.sourceCode += "\n" + newSourceCode
}
