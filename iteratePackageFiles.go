package parser

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type packageFilesIterator = func(currFile *ast.File, filePkg *packages.Package, parentLog LoggerCLI) error

// iteratePackageFiles will iterate over the files inside the parsed packages.
//
// Note that if the focus is set to filepath, it will iterate only over the specified
// file
func (p *GoParser) iteratePackageFiles(callback packageFilesIterator) error {
	packagesIterator := func(pkg *packages.Package, parentLog LoggerCLI) error {
		if len(pkg.Syntax) == 0 {
			parentLog.Debug("Skipped (zero Syntax objects)...")
			return nil
		}

		for _, currFile := range pkg.Syntax {
			currFilePath := p.fileSet.File(currFile.Pos()).Name()
			log := parentLog.Debug("Analysing *ast.File '%s'...", currFilePath)

			if !p.focus.is(focusFilePath, currFilePath) {
				log.Debug("Skipped (not the focus)...")
				continue
			}

			e := callback(currFile, pkg, log)
			if e != nil {
				return e
			}
		}

		return nil
	}
	return p.iteratePackages(packagesIterator)
}
