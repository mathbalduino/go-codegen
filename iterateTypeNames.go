package parser

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type typeNamesIterator = func(type_ *types.TypeName, logger LoggerCLI) error

// iterateTypeNames will iterate over all the typeNames inside the parsed files.
//
// Note that if the focus is set to typeName, it will iterate only over the specified
// typeName
func (p *GoParser) iterateTypeNames(callback typeNamesIterator, optionalLogger ...LoggerCLI) error {
	packageFilesIterator := func(file *ast.File, typePkg *packages.Package, logger LoggerCLI) error {
		logger = logger.Trace("Iterating over file TypeNames...")
		if len(file.Scope.Objects) == 0 {
			logger.Trace("Skipped (zero objects)...")
			return nil
		}

		for _, currObj := range file.Scope.Objects {
			currLogger := logger.Trace("Analysing *ast.Object '%s'...", currObj.Name)

			typeSpec, isTypeSpec := currObj.Decl.(*ast.TypeSpec)
			if !isTypeSpec {
				currLogger.Trace("Skipped (not a TypeSpec)...")
				continue
			}
			typeObj, exists := typePkg.TypesInfo.Defs[typeSpec.Name]
			if !exists {
				currLogger.Trace("Skipped (missing TypesInfo.Defs information)...")
				continue
			}
			typeName, isTypeName := typeObj.(*types.TypeName)
			if !isTypeName {
				currLogger.Trace("Skipped (not a TypeName)...")
				continue
			}
			if !p.focus.is(focusTypeName, typeName.Name()) {
				currLogger.Trace("Skipped (not the focus)...")
				continue
			}

			e := callback(typeName, currLogger)
			if e != nil {
				return e
			}
		}

		return nil
	}
	return p.iteratePackageFiles(packageFilesIterator, optionalLogger...)
}
