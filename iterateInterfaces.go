package parser

import (
	"go/types"
)

type InterfacesIterator = func(type_ *types.TypeName, parentLog LoggerCLI) error

// IterateInterfaces will iterate only over the interfaces that are defined
// inside the parsed go code
func (p *GoParser) IterateInterfaces(callback InterfacesIterator) error {
	typeNamesIterator := func(type_ *types.TypeName, parentLog LoggerCLI) error {
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
	return p.iterateTypeNames(typeNamesIterator)
}
