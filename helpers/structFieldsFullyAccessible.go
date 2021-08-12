package helpers

import (
	goParser "gitlab.com/matheuss-leonel/go-codegen"
	"go/types"
)

func StructFieldsFullyAccessible(struct_ *types.Struct, fromPackagePath string, parentLog goParser.LogCLI) bool {
	log := parentLog.Debug("Iterating struct fields to see if they're fully accessible from package '%s'...",
		fromPackagePath)
	if struct_.NumFields() == 0 {
		log.Debug("Accessible: doesn't contain any fields")
		return true
	}

	for i := 0; i < struct_.NumFields(); i++ {
		field := struct_.Field(i)
		fieldName := field.Name()
		fieldLog := log.Debug("Field '%s'...", fieldName)

		// struct fields will never have nil Pkg()
		if field.Pkg().Path() != fromPackagePath {
			if !field.Exported() {
				fieldLog.Debug("Not accessible: different package (%s %s) and not exported",
					field.Pkg().Name(), field.Pkg().Path())
				return false
			}

			fieldLog.Debug("Accessible: different package (%s %s) but exported",
				field.Pkg().Name(), field.Pkg().Path())
		} else {
			fieldLog.Debug("Accessible: same package")
		}

		fieldType := field.Type()
		fieldUnderlyingType, isStruct := fieldType.Underlying().(*types.Struct)
		if !isStruct || !field.Embedded() {
			typeLog := fieldLog.Debug("Analysing field type '%s'...", fieldType.String())
			if !typeIdentifierIsAccessible(fieldType, fromPackagePath, typeLog) {
				return false
			}

			continue
		}

		embeddedLog := fieldLog.Debug("Recursively analysing embedded struct field type '%s'...", fieldType.String())
		if !StructFieldsFullyAccessible(fieldUnderlyingType, fromPackagePath, embeddedLog) {
			return false
		}
	}

	log.Debug("Accessible: all fields (and it's types) are exported or in the same package")
	return true
}
