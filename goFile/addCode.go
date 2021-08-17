package goFile

// AddCode will append the given param to the file
// source code body (below imports)
func (f *GoFile) AddCode(newSourceCode string) {
	f.sourceCode += "\n" + newSourceCode
}
