package helpers

import (
	goParser "gitlab.com/matheuss-leonel/go-codegen"
	"go/types"
)

func ObjectIsAccessible(obj types.Object, fromPackagePath string, parentLog goParser.LogCLI) bool {
	log := parentLog.Debug("Checking to see if '%s' is accessible from '%s'...", obj.Name(), fromPackagePath)
	if obj.Pkg() == nil {
		log.Debug("Accessible: builtin type")
		return true
	}
	if obj.Pkg().Path() == fromPackagePath {
		log.Debug("Accessible: same package")
		return true
	}
	if !obj.Exported() {
		log.Debug("Not accessible: different package (%s %s) and not exported", obj.Pkg().Name(), obj.Pkg().Path())
		return false
	}

	log.Debug("Accessible: different package (%s %s) but exported", obj.Pkg().Name(), obj.Pkg().Path())
	return true
}
