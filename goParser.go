package goParser

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

// LogCLI is used to log the parser actions, using only
// two levels: Debug and Error
type LogCLI interface {
	Debug(msgFormat string, args ...interface{}) LogCLI
	Error(msgFormat string, args ...interface{}) LogCLI
}

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

	log     LogCLI
	fileSet *token.FileSet
}

// NewGoParser creates a new parser for GO source files
func NewGoParser(pattern string, config Config, log LogCLI) (*GoParser, error) {
	printFinalConfig(pattern, config, log)

	packagesLoadConfig := packagesLoadConfig(config, log)
	pkgs, e := packages.Load(packagesLoadConfig, pattern)
	if e != nil {
		return nil, e
	}

	return &GoParser{
		pkgs,
		config.Focus,
		log,
		packagesLoadConfig.Fset,
	}, nil
}
