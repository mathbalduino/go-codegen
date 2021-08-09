package goParser

import (
	"fmt"
)

// printFinalConfig will print, as a Debug log, the final configuration of the parser
func printFinalConfig(pattern string, config Config, log LogCLI) {
	focus := "Focus not defined (will not skip anything)"
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

	fset := "Using the FileSet of the library"
	if config.Fset != nil {
		fset = "Using the FileSet provided by the client"
	}

	dir := "./"
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
