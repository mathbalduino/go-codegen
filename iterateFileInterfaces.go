package parser

import (
	"go/types"
)

type FileInterfacesIterator = func(type_ *types.TypeName, parentLog LogCLI) error

// IterateFileInterfaces will iterate only over the interfaces that are defined
// inside the parsed files
func (p *GoParser) IterateFileInterfaces(callback FileInterfacesIterator) error {
	fileTypeNamesIterator := func(type_ *types.TypeName, parentLog LogCLI) error {
		log := parentLog.Debug("Analysing *types.TypeName '%s'...", type_.Name())
		_, isInterface := type_.Type().Underlying().(*types.Interface)
		if !isInterface {
			log.Debug("Skipped (not a interface)...")
			return nil
		}

		e := callback(type_, log)
		if e != nil {
			return e
		}

		return nil
	}
	return p.iterateFileTypeNames(fileTypeNamesIterator)
}
