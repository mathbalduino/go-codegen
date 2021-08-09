package goParser

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type fileTypeNamesIterator = func(type_ *types.TypeName, parentLog LogCLI) error

// iterateFileTypeNames will iterate over all the typeNames inside the parsed files.
//
// Note that if the focus is set to typeName, it will iterate only over the specified
// typeName
func (p *GoParser) iterateFileTypeNames(callback fileTypeNamesIterator) error {
	packageFilesIterator := func(file *ast.File, typePkg *packages.Package, parentLog LogCLI) error {
		if len(file.Scope.Objects) == 0 {
			parentLog.Debug("Skipped (zero objects)...")
			return nil
		}

		for _, currObj := range file.Scope.Objects {
			log := parentLog.Debug("Analysing *ast.Object '%s'...", currObj.Name)

			typeSpec, isTypeSpec := currObj.Decl.(*ast.TypeSpec)
			if !isTypeSpec {
				log.Debug("Skipped (not a TypeSpec)...")
				continue
			}
			typeObj, exists := typePkg.TypesInfo.Defs[typeSpec.Name]
			if !exists {
				log.Debug("Skipped (missing TypesInfo.Defs information)...")
				continue
			}
			typeName, isTypeName := typeObj.(*types.TypeName)
			if !isTypeName {
				log.Debug("Skipped (not a TypeName)...")
				continue
			}
			if !p.focus.is(focusTypeName, typeName.Name()) {
				log.Debug("Skipped (not the focus)...")
				continue
			}

			e := callback(typeName, log)
			if e != nil {
				return e
			}
		}

		return nil
	}
	return p.iteratePackageFiles(packageFilesIterator)
}
