// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-30 11:16 AM

package parser

import (
	logger "github.com/mathbalduino/go-log"
	"github.com/mathbalduino/go-log/loggerCLI"
)

// Exports some LoggerCLI things, to avoid forcing the user
// of go-codegen to directly depend on go-log

type LoggerCLI = loggerCLI.LoggerCLI

var NewLoggerCLI = loggerCLI.New
var ParseLogLevel = loggerCLI.ParseLogLevel
var ValidateLogLevels = loggerCLI.ValidateLogLevels

const (
	LogLevelsValues = loggerCLI.LogLevelsValues

	// LogTrace is a flag that if used will enable
	// Trace logs
	LogTrace = logger.LvlTrace

	// LogDebug is a flag that if used will enable
	// Debug logs
	LogDebug = logger.LvlDebug

	// LogInfo is a flag that if used will enable
	// Info logs
	LogInfo = logger.LvlInfo

	// LogWarn is a flag that if used will enable
	// Warn logs
	LogWarn = logger.LvlWarn

	// LogError is a flag that if used will enable
	// Error logs
	LogError = logger.LvlError

	// LogFatal is a flag that if used will enable
	// Fatal logs
	LogFatal = logger.LvlFatal

	// LogJSON set the logs to be parsed
	// to JSON before printing it to the
	// stdout (one per line)
	LogJSON = uint64(1 << 6)
)
