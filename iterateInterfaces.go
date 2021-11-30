package parser

import (
	"go/types"
)

type InterfacesIterator = func(type_ *types.TypeName, logger LoggerCLI) error

// IterateInterfaces will iterate only over the interfaces that are defined
// inside the parsed go code
func (p *GoParser) IterateInterfaces(callback InterfacesIterator, optionalLogger ...LoggerCLI) error {
	typeNamesIterator := func(type_ *types.TypeName, logger LoggerCLI) error {
		logger = logger.Trace("Analysing *types.TypeName '%s'...", type_.Name())
		_, isInterface := type_.Type().Underlying().(*types.Interface)
		if !isInterface {
			logger.Trace("Skipped (not a interface)...")
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
