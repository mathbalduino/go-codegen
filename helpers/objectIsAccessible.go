package helpers

import (
	goParser "github.com/mathbalduino/go-codegen"
	"go/types"
)

// ObjectIsAccessible will recursively check an object to see if it is accessible from the given
// package path (is it public? Private? Builtin?...)
func ObjectIsAccessible(obj types.Object, fromPackagePath string, logger goParser.LoggerCLI) bool {
	logger = logger.Trace("Checking to see if '%s' is accessible from '%s'...", obj.Name(), fromPackagePath)
	if obj.Pkg() == nil {
		logger.Trace("Accessible: builtin type")
		return true
	}
	if obj.Pkg().Path() == fromPackagePath {
		logger.Trace("Accessible: same package")
		return true
	}
	if !obj.Exported() {
		logger.Trace("Not accessible: different package and not exported (%s %s)", obj.Pkg().Name(), obj.Pkg().Path())
		return false
	}

	logger.Trace("Accessible: different package but exported (%s %s)", obj.Pkg().Name(), obj.Pkg().Path())
	return true
}
