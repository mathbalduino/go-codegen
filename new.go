package parser

import (
	"github.com/mathbalduino/go-log/loggerCLI"
	"go/token"
	"golang.org/x/tools/go/packages"
)

// GoParser is a type that can parse
// GO source code and iterate over the
// parsed code
type GoParser struct {
	// pkgs is a slice that contains all the
	// parsed packages
	pkgs []*packages.Package

	// focus tell the parser what kind of thing
	// it must focus on. It can be a type, a file,
	// an entire package, etc
	focus   *ParserFocus

	// logger is used to print information to stdout
	// about each step
	logger *loggerCLI.LoggerCLI

	// fileSet will store information about the
	// files paths. This information can be used
	// by some methods
	fileSet *token.FileSet
}

// NewGoParser creates a new parser for GO source files
func NewGoParser(pattern string, config Config) (*GoParser, error) {
	logger := loggerCLI.New(
		config.LogFlags & LogJSON != 0,
			config.LogFlags & LogDebug != 0,
			config.LogFlags & LogTrace != 0,
	)
	printFinalConfig(pattern, config, logger)
	packagesLoadConfig := packagesLoadConfig(config, logger)
	pkgs, e := packages.Load(packagesLoadConfig, pattern)
	if e != nil {
		return nil, e
	}

	return &GoParser{
		pkgs,
		config.Focus,
		logger,
		packagesLoadConfig.Fset,
	}, nil
}
