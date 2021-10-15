package parser

import (
	"github.com/mathbalduino/go-log/loggerCLI"
	"golang.org/x/tools/go/packages"
)

type LoggerCLI = loggerCLI.LoggerCLI

type packagesIterator = func(pkg *packages.Package, parentLog LoggerCLI) error

// iteratePackages will iterate over the parsed packages.
//
// Note that if the focus is set to packagePath, it will iterate only over
// the specified package
func (p *GoParser) iteratePackages(callback packagesIterator) error {
	if len(p.pkgs) == 0 {
		p.logger.Debug("There are no packages to iterate...")
		return nil
	}

	for _, currPkg := range p.pkgs {
		log := p.logger.Debug("Analysing *packages.Package '%s %s'...", currPkg.Name, currPkg.PkgPath)

		if len(currPkg.Errors) != 0 {
			errorsLog := log.Error("Package '%s %s' contain errors. Skipping it...", currPkg.Name, currPkg.PkgPath)
			for _, currError := range currPkg.Errors {
				errorsLog.Error(currError.Error())
			}
			continue
		}
		if !p.focus.is(focusPackagePath, currPkg.PkgPath) {
			log.Debug("Skipped (not the focus)...")
			continue
		}

		e := callback(currPkg, log)
		if e != nil {
			return e
		}
	}

	return nil
}
