package parser

import (
	logger "github.com/mathbalduino/go-log"
	"go/token"
	"golang.org/x/tools/go/packages"
)

// Config holds some information about the
// parser behaviour.
//
// Some fields of the packages.Config struct may
// appear inside this Config struct, exposing
// them to client customization (others hidden)
type Config struct {
	// Tests (from packages.Config) if set, the
	// loader includes not just the packages matching
	// a particular pattern but also any related test
	// packages, including test-only variants of the
	// package and the test executable.
	//
	// For example, when using the go command, loading
	// "fmt" with Tests=true returns four packages, with
	// IDs "fmt" (the standard package), "fmt [fmt.test]"
	// (the package as compiled for the test), "fmt_test"
	// (the test functions from source files in package fmt_test),
	// and "fmt.test" (the test binary).
	//
	// In build systems with explicit names for tests,
	// setting Tests may have no effect.
	Tests bool

	// Dir (from packages.Config) is the directory
	// in which to run the build system's query tool
	// that provides information about the packages.
	//
	// If Dir is empty, the tool is run in the current
	// directory
	Dir string

	// Env (from packages.Config) is the environment
	// to use when invoking the build system's query
	// tool.
	//
	// If Env is nil, the current environment is used.
	// As in os/exec's Cmd, only the last value in the
	// slice for each environment key is used. To specify
	// the setting of only a few variables, append to the
	// current environment, as in:
	// 		opt.Env = append(os.Environ(), "GOOS=plan9", "GOARCH=386")
	Env []string

	// Fset (from packages.Config) provides source
	// position information for syntax trees and
	// types.
	//
	// If Fset is nil, go-codegen will create a new
	// fileset.
	Fset *token.FileSet

	// BuildFlags (from packages.Config) is a list of
	// command-line flags to be passed through to the
	// build system's query tool.
	BuildFlags []string

	// Focus sets the typename/filename/pkgname/etc that
	// the parser should focus.
	//
	// Note that if you set the focus to be some pkg, the
	// filenames, typenames, etc, will continue to be used.
	// Only other pkgs will be ignored (the same for filenames,
	// typenames, etc)
	Focus *Focus

	// LogFlags controls the logging configuration of the parser.
	// It can be set using the constants LogTrace, LogDebug and
	// LogJSON.
	//
	// Note that you can use a bitwise-AND operator to combine
	// multiple flags
	LogFlags uint64
}

// packagesLoadConfig returns the configuration struct that is used to
// call the packages.Load function
func packagesLoadConfig(config Config, log LoggerCLI) *packages.Config {
	if config.Fset == nil {
		config.Fset = token.NewFileSet()
	}

	return &packages.Config{
		// Customizable configurations
		Env:        config.Env,
		BuildFlags: config.BuildFlags,
		Tests:      config.Tests,
		Dir:        config.Dir,
		Fset:       config.Fset,

		// Constant values, don't exposed
		Mode: packagesConfigMode,
		Logf: func(format string, args ...interface{}) {
			log.Debug(format, args...)
		},

		// Not used
		Context:   nil,
		ParseFile: nil,
		Overlay:   nil,
	}
}

const packagesConfigMode = packages.NeedImports |
	packages.NeedSyntax |
	packages.NeedName |
	packages.NeedTypes |
	packages.NeedTypesInfo

const (
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
