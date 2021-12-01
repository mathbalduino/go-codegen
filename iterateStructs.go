package parser

import (
	"go/types"
)

type StructsIterator = func(type_ *types.TypeName, logger LoggerCLI) error

// IterateStructs will iterate only over the structs that are defined inside
// the parsed go code
func (p *GoParser) IterateStructs(callback StructsIterator, optionalLogger ...LoggerCLI) error {
	typeNamesIterator := func(type_ *types.TypeName, logger LoggerCLI) error {
		logger = logger.Trace("Analysing *types.TypeName '%s'...", type_.Name())
		_, isStruct := type_.Type().Underlying().(*types.Struct)
		if !isStruct {
			logger.Trace("Skipped (not a struct)...")
			return nil
		}

		e := callback(type_, logger)
		if e != nil {
			return e
		}

		return nil
	}
	return p.iterateTypeNames(typeNamesIterator, optionalLogger...)
}
