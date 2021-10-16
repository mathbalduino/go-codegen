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
	focus *ParserFocus

	// logger is used to print information to stdout
	// about each step
	logger LoggerCLI

	// fileSet is the one forwarded to the underlying
	// packages.Load function
	fileSet *token.FileSet
}

// NewGoParser creates a new parser for GO source files
//
// The pattern argument will be forwarded directly to the
// packages.Load function. Take a look at the docs:
// 		Load passes most patterns directly to the underlying build tool,
//		but all patterns with the prefix "query=", where query is a non-empty
//		string of letters from [a-z], are reserved and may be interpreted
//		as query operators
//
// 		Two query operators are currently supported: "file" and "pattern"
//
//		The query "file=path/to/file.go" matches the package or packages
//		enclosing the Go source file path/to/file.go. For example
//		"file=~/go/src/fmt/print.go" might return the packages "fmt" and
//		"fmt [fmt.test]"
//
//		The query "pattern=string" causes "string" to be passed directly to
//		the underlying build tool. In most cases this is unnecessary, but an
//		application can use Load("pattern=" + x) as an escaping mechanism to
//		ensure that x is not interpreted as a query operator if it contains
//		'='
//
// 		All other query operators are reserved for future use and currently
//		cause Load to report an error
//
// 		 Note that one pattern can match multiple packages and that a package
//		 might be matched by multiple patterns: in general it is not possible
//		 to determine which packages correspond to which patterns.
func NewGoParser(pattern string, config Config) (*GoParser, error) {
	logger := loggerCLI.New(
		config.LogFlags&LogJSON != 0,
		config.LogFlags&LogDebug != 0,
		config.LogFlags&LogTrace != 0,
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
