package parser

import (
	"golang.org/x/tools/go/packages"
)

type packagesIterator = func(pkg *packages.Package, logger LoggerCLI) error

// iteratePackages will iterate over the parsed packages.
//
// Note that if the focus is set to packagePath, it will iterate only over
// the specified package
func (p *GoParser) iteratePackages(callback packagesIterator, optionalLogger ...LoggerCLI) error {
	logger := p.logger
	if len(optionalLogger) != 0 {
		logger = optionalLogger[0]
	}

	if len(p.pkgs) == 0 {
		logger.Trace("There are no packages to iterate...")
		return nil
	}

	logger = logger.Trace("Iterating over packages...")
	for _, currPkg := range p.pkgs {
		currLogger := logger.Trace("Analysing *packages.Package '%s %s'...", currPkg.Name, currPkg.PkgPath)

		if len(currPkg.Errors) != 0 {
			errorsLog := currLogger.Error("Package '%s %s' contain errors. Skipping it...", currPkg.Name, currPkg.PkgPath)
			for _, currError := range currPkg.Errors {
				errorsLog.Error(currError.Error())
			}
			continue
		}
		if !p.focus.is(focusPackagePath, currPkg.PkgPath) {
			currLogger.Trace("Skipped (not the focus)...")
			continue
		}

		e := callback(currPkg, currLogger)
		if e != nil {
			return e
		}
	}

	return nil
}
