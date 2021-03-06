package parser

import (
	"fmt"
	"strings"
)

// printFinalConfig will print, as a Debug log, the final configuration of the parser
func printFinalConfig(pattern string, config Config, logger LoggerCLI) {
	focus := nilFocusStr
	if config.Focus != nil {
		packagePath, filePath, typeName := "nil", "nil", "nil"
		if config.Focus.packagePath != nil {
			packagePath = *config.Focus.packagePath
		}
		if config.Focus.filePath != nil {
			filePath = *config.Focus.filePath
		}
		if config.Focus.typeName != nil {
			typeName = *config.Focus.typeName
		}

		focus = fmt.Sprintf(focusTemplate, packagePath, filePath, typeName)
	}

	fset := nilFsetStr
	if config.Fset != nil {
		fset = notNilFsetStr
	}

	dir := emptyDirStr
	if config.Dir != "" {
		dir = config.Dir
	}

	loggerMsg := nilLoggerStr
	if config.Logger != nil {
		loggerMsg = notNilLoggerStr
	}

	logFlags := emptyLogFlagsStr
	if config.Logger != nil {
		logFlags = ignoredLogFlagsStr
	} else if config.LogFlags != 0 {
		if config.LogFlags&LogJSON != 0 {
			logFlags = "LogJSON | "
		}
		if config.LogFlags&LogTrace != 0 {
			logFlags = "LogTrace | "
		}
		if config.LogFlags&LogDebug != 0 {
			logFlags = "LogDebug | "
		}
		if config.LogFlags&LogInfo != 0 {
			logFlags = "LogInfo | "
		}
		if config.LogFlags&LogWarn != 0 {
			logFlags = "LogWarn | "
		}
		if config.LogFlags&LogError != 0 {
			logFlags = "LogError | "
		}
		if config.LogFlags&LogFatal != 0 {
			logFlags = "LogFatal"
		}
		logFlags = strings.TrimSuffix(logFlags, " | ")
	}

	logger.Debug(finalConfigTemplate,
		pattern,
		config.Tests,
		dir,
		config.Env,
		config.BuildFlags,
		focus,
		fset,
		loggerMsg,
		logFlags,
	)
}

const emptyDirStr = "./"
const nilFsetStr = "Using the FileSet of the library"
const notNilFsetStr = "Using the FileSet provided by the client"
const nilFocusStr = "Focus not defined (will not skip anything)"
const nilLoggerStr = "Logger not defined (will be created using LogFlags)"
const notNilLoggerStr = "Using the Logger instance provided by the client"
const emptyLogFlagsStr = "unset"
const ignoredLogFlagsStr = "Ignored (Logger field is set)"

const finalConfigTemplate = `New GoParser created. Final configuration:
Pattern: %s
Config: {
	Tests: %t
	Dir: %s
	Env: %v
	BuildFlags: %v
	Focus: %s
	Fset: %s
	Logger: %s
	LogFlags: %s
}`

const focusTemplate = `{
		packagePath: %s
		filePath: %s
		typeName: %s
	}`
