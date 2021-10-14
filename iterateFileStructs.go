package parser

import (
	"go/types"
)

type FileStructsIterator = func(type_ *types.TypeName, parentLog LoggerCLI) error

// IterateFileStructs will iterate only over the structs that are defined inside
// the parsed files
func (p *GoParser) IterateFileStructs(callback FileStructsIterator) error {
	fileTypeNamesIterator := func(type_ *types.TypeName, parentLog LoggerCLI) error {
		log := parentLog.Debug("Analysing *types.TypeName '%s'...", type_.Name())
		_, isStruct := type_.Type().Underlying().(*types.Struct)
		if !isStruct {
			log.Debug("Skipped (not a struct)...")
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
