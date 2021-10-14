package helpers

import (
	goParser "github.com/mathbalduino/go-codegen"
	"go/types"
)

// ObjectIsAccessible will recursively check an object to see if it is accessible from the given
// package path (is it public? Private? Builtin?...)
func ObjectIsAccessible(obj types.Object, fromPackagePath string, parentLog goParser.LoggerCLI) bool {
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
