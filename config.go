package parser

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

// Config holds some information about the
// parser behaviour.
//
// Some fields of the packages.Config struct
// may appear inside this Config struct, exposing
// them to client customization
type Config struct {
	Tests      bool
	Dir        string
	Env        []string
	BuildFlags []string
	Focus      *ParserFocus
	Fset       *token.FileSet
}

// packagesLoadConfig is the configuration of the packages.Load function
func packagesLoadConfig(config Config, log LogCLI) *packages.Config {
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
		Logf: func(format string, args ...interface{}) { log.Debug(format, args...) },

		// Not used
		Context:   nil,
		ParseFile: nil,
		Overlay:   nil,
	}
}

const packagesConfigMode =
	packages.NeedImports |
	packages.NeedSyntax |
	packages.NeedName |
	packages.NeedTypes |
	packages.NeedTypesInfo
