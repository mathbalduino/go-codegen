package parser

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type packageFilesIterator = func(currFile *ast.File, filePkg *packages.Package, logger LoggerCLI) error

// iteratePackageFiles will iterate over the files inside the parsed packages.
//
// Note that if the focus is set to filepath, it will iterate only over the specified
// file
func (p *GoParser) iteratePackageFiles(callback packageFilesIterator, optionalLogger ...LoggerCLI) error {
	packagesIterator := func(pkg *packages.Package, logger LoggerCLI) error {
		logger = logger.Trace("Iterating over package files...")
		if len(pkg.Syntax) == 0 {
			logger.Trace("Skipped (zero Syntax objects)...")
			return nil
		}

		for _, currFile := range pkg.Syntax {
			currFilePath := p.fileSet.File(currFile.Pos()).Name()
			logger = logger.Trace("Analysing *ast.File '%s'...", currFilePath)

			if !p.focus.is(focusFilePath, currFilePath) {
				logger.Trace("Skipped (not the focus)...")
				continue
			}

			e := callback(currFile, pkg, logger)
			if e != nil {
				return e
			}
		}

		return nil
	}
	return p.iteratePackages(packagesIterator, optionalLogger...)
}
