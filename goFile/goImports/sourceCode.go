package goImports

import (
	"fmt"
)

// SourceCode will transform the list of imports in
// GO source code, ready to be compiled
func (i *GoImports) SourceCode() string {
	if len(i.imports) == 0 {
		return ""
	}

	sourceCode := "import (\n"
	for alias, packagePath := range i.imports {
		sourceCode += fmt.Sprintf(`%s "%s"`, alias, packagePath) + "\n"
	}
	sourceCode += ")\n"

	return sourceCode
}
