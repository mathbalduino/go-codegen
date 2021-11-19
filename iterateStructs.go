package parser

import (
	"go/types"
)

type StructsIterator = func(type_ *types.TypeName, parentLog LoggerCLI) error

// IterateStructs will iterate only over the structs that are defined inside
// the parsed go code
func (p *GoParser) IterateStructs(callback StructsIterator) error {
	typeNamesIterator := func(type_ *types.TypeName, parentLog LoggerCLI) error {
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
	return p.iterateTypeNames(typeNamesIterator)
}
