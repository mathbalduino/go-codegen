package goParser

import (
	"fmt"
)

// printFinalConfig will print, as a Debug log, the final configuration of the parser
func printFinalConfig(pattern string, config Config, log LogCLI) {
	focus := nilFocusStr
	if config.Focus != nil {
		packagePath, filePath, typeName, varName, functionName := "nil", "nil", "nil", "nil", "nil"
		if config.Focus.packagePath != nil {
			packagePath = *config.Focus.packagePath
		}
		if config.Focus.filePath != nil {
			filePath = *config.Focus.filePath
		}
		if config.Focus.typeName != nil {
			typeName = *config.Focus.typeName
		}
		if config.Focus.varName != nil {
			varName = *config.Focus.varName
		}
		if config.Focus.functionName != nil {
			functionName = *config.Focus.functionName
		}

		focus = fmt.Sprintf(focusTemplate, packagePath, filePath, typeName, varName, functionName)
	}

	fset := nilFsetStr
	if config.Fset != nil {
		fset = notNilFsetStr
	}

	dir := emptyDirStr
	if config.Dir != "" {
		dir = config.Dir
	}

	log.Debug(finalConfigTemplate,
		pattern,
		config.Tests,
		dir,
		config.Env,
		config.BuildFlags,
		focus,
		fset,
	)
}

const emptyDirStr = "./"
const nilFsetStr = "Using the FileSet of the library"
const notNilFsetStr = "Using the FileSet provided by the client"
const nilFocusStr = "Focus not defined (will not skip anything)"

const finalConfigTemplate = `New GoParser created. Final configuration:
Pattern: %s
Config: {
	Tests: %t
	Dir: %s
	Env: %v
	BuildFlags: %v
	Focus: %s
	Fset: %s
}`

const focusTemplate = `{
		packagePath: %s
		filePath: %s
		typeName: %s
		varName: %s
		functionName: %s
	}`
