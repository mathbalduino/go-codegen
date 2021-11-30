package tsImports

import (
	"fmt"
	"strings"
)

// SourceCode will return valid Typescript - TS code,
// for the stored imports, ready to be compiled
func (i TsImports) SourceCode() string {
	if len(i) == 0 {
		return ""
	}

	sourceCode := "\n"
	for _, currImport := range i {
		importCode := "import "
		if currImport.defaultImport != "" {
			importCode += currImport.defaultImport

			if len(currImport.namedImports) != 0 {
				importCode += ", "
			}
		}

		if len(currImport.namedImports) != 0 {
			importCode += "{ "
			for _, currNamedImport := range currImport.namedImports {
				importCode += currNamedImport + ", "
			}
			importCode = strings.TrimSuffix(importCode, ", ") + " }"
		}

		importCode += fmt.Sprintf(" from '%s'", strings.TrimSuffix(currImport.path, ".ts"))
		sourceCode += importCode + "\n"
	}

	return sourceCode
}
