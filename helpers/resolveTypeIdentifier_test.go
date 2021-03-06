package helpers

import (
	"fmt"
	"go/types"
	"testing"
)

func TestResolveTypeIdentifier(t *testing.T) {
	t.Run("Should return the literal name of the Basic types", func(t *testing.T) {
		id, e := ResolveTypeIdentifier(types.Typ[1], nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != types.Typ[1].Name() {
			t.Fatalf("Expected to return the literal typeName")
		}
	})
	t.Run("For pointers, should recursively call ResolveTypeIdentifier for the element type and return it with the '*' prefix", func(t *testing.T) {
		type_ := types.NewPointer(types.Typ[1])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("*%s", types.Typ[1].Name()) {
			t.Fatalf("Expected to return '*typeName'")
		}
	})
	t.Run("For pointers, return errors when recursively resolving the element", func(t *testing.T) {
		type_ := types.NewPointer(&fakeType{})
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error is not the expected one")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Arrays, should recursively call ResolveTypeIdentifier for the element type and return it with the '[n]' prefix", func(t *testing.T) {
		length := int64(5)
		type_ := types.NewArray(types.Typ[1], length)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("[%d]%s", length, types.Typ[1].Name()) {
			t.Fatalf("Expected to return '[n]typeName'")
		}
	})
	t.Run("For Arrays, return errors when recursively resolving the element type", func(t *testing.T) {
		length := int64(5)
		type_ := types.NewArray(&fakeType{}, length)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Slices, should recursively call ResolveTypeIdentifier for the element type and return it with the '[]' prefix", func(t *testing.T) {
		type_ := types.NewSlice(types.Typ[1])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("[]%s", types.Typ[1].Name()) {
			t.Fatalf("Expected to return '[]typeName'")
		}
	})
	t.Run("For Slices, return errors when recursively resolving the element type", func(t *testing.T) {
		type_ := types.NewSlice(&fakeType{})
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Maps, should recursively call ResolveTypeIdentifier for the element and key types, returning 'map[<keyType>]<elemType>'", func(t *testing.T) {
		type_ := types.NewMap(types.Typ[1], types.Typ[2])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("map[%s]%s", types.Typ[1].Name(), types.Typ[2].Name()) {
			t.Fatalf("Expected to return 'map[keyType]elemType'")
		}
	})
	t.Run("For Maps, return errors when recursively resolving the element type", func(t *testing.T) {
		type_ := types.NewMap(&fakeType{}, types.Typ[2])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Maps, return errors when recursively resolving the key type", func(t *testing.T) {
		type_ := types.NewMap(types.Typ[1], &fakeType{})
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For SendOnly Channels, should recursively call ResolveTypeIdentifier for the element type, returning 'chan<- <elemType>'", func(t *testing.T) {
		type_ := types.NewChan(types.SendOnly, types.Typ[1])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("chan<- %s", types.Typ[1].Name()) {
			t.Fatalf("Expected to return 'chan<- elemType'")
		}
	})
	t.Run("For SendOnly Channels, return errors when recursively resolving the element type", func(t *testing.T) {
		type_ := types.NewChan(types.SendOnly, &fakeType{})
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For RecvOnly Channels, should recursively call ResolveTypeIdentifier for the element, returning '<-chan <elemType>'", func(t *testing.T) {
		type_ := types.NewChan(types.RecvOnly, types.Typ[1])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("<-chan %s", types.Typ[1].Name()) {
			t.Fatalf("Expected to return '<-chan elemType'")
		}
	})
	t.Run("For SendRecv Channels, should recursively call ResolveTypeIdentifier for the element, returning 'chan <elemType>'", func(t *testing.T) {
		type_ := types.NewChan(types.SendRecv, types.Typ[1])
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("chan %s", types.Typ[1].Name()) {
			t.Fatalf("Expected to return 'chan elemType'")
		}
	})
	t.Run("For empty Structs, should return 'struct{}'", func(t *testing.T) {
		id, e := ResolveTypeIdentifier(types.NewStruct(nil, nil), nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "struct{}" {
			t.Fatalf("Expected to return 'struct{}'")
		}
	})
	t.Run("For Structs, should recursively call ResolveTypeIdentifier for every field type, returning 'struct{ <fields>; }'", func(t *testing.T) {
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}, nil)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("struct{fieldA %s; fieldB %s; fieldC %s}", types.Typ[1].Name(), types.Typ[2].Name(), types.Typ[3].Name()) {
			t.Fatalf("Expected to return 'struct{ <fields>; }'")
		}
	})
	t.Run("For Structs, return errors when recursively resolving every field type", func(t *testing.T) {
		type_ := types.NewStruct([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", &fakeType{}),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}, nil)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For empty Tuples, should return an empty string", func(t *testing.T) {
		type_ := types.NewTuple()
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("Expected to return an empty string")
		}
	})
	t.Run("For empty Tuples, return errors when recursively resolving the types", func(t *testing.T) {
		type_ := types.NewTuple(types.NewVar(0, nil, "field", &fakeType{}))
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Tuples, should recursively call ResolveTypeIdentifier for every elem, returning '<elemA>, <elemB>, <elemC>'", func(t *testing.T) {
		type_ := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("%s, %s, %s", types.Typ[1].Name(), types.Typ[2].Name(), types.Typ[3].Name()) {
			t.Fatalf("Expected to return '<elemA>, <elemB>, <elemC>'")
		}
	})
	t.Run("For Function Signatures without params and return value, should return 'func()'", func(t *testing.T) {
		type_ := types.NewSignature(nil, nil, nil, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "func()" {
			t.Fatalf("Expected to return 'func()'")
		}
	})
	t.Run("For Function Signatures without params with single return type, should return 'func() <type>'", func(t *testing.T) {
		tuples := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
		}...)
		type_ := types.NewSignature(nil, nil, tuples, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("func() %s", types.Typ[1]) {
			t.Fatalf("Expected to return 'func() <type>'")
		}
	})
	t.Run("For Function Signatures without params with single return type, return errors resolving the return type", func(t *testing.T) {
		tuples := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", &fakeType{}),
		}...)
		type_ := types.NewSignature(nil, nil, tuples, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Function Signatures without params with multiple return types, should return 'func() (<typeA>, <typeB>)'", func(t *testing.T) {
		tuples := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, nil, tuples, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("func() (%s, %s, %s)", types.Typ[1], types.Typ[2], types.Typ[3]) {
			t.Fatalf("Expected to return 'func() (<typeA>, <typeB>)'")
		}
	})
	t.Run("For Function Signatures with params and without return type, should return 'func(<paramA>, <paramB>)'", func(t *testing.T) {
		params := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, params, nil, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("func(%s, %s, %s)", types.Typ[1], types.Typ[2], types.Typ[3]) {
			t.Fatalf("Expected to return 'func(<typeA>, <typeB>)'")
		}
	})
	t.Run("For Function Signatures with params and without return type, return errors resolving params types", func(t *testing.T) {
		params := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", &fakeType{}),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, params, nil, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Function Signatures with params and return type, should return 'func(<paramA>, <paramB>) (<typeA>, <typeB>)'", func(t *testing.T) {
		params := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		returnTypes := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, params, returnTypes, false)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("func(%s, %s, %s) (%s, %s, %s)", types.Typ[1], types.Typ[2], types.Typ[3], types.Typ[1], types.Typ[2], types.Typ[3]) {
			t.Fatalf("Expected to return 'func(<typeA>, <typeB>) (<typeA>, <typeB>)'")
		}
	})
	t.Run("For Function Signatures with params (plus variadic) and return type, should return 'func(<paramA>, ...<variadicType>) (<typeA>, <typeB>)'", func(t *testing.T) {
		params := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
			types.NewVar(0, nil, "fieldD", types.NewSlice(types.Typ[4])),
		}...)
		returnTypes := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, params, returnTypes, true)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("func(%s, %s, %s, ...%s) (%s, %s, %s)", types.Typ[1], types.Typ[2], types.Typ[3], types.Typ[4], types.Typ[1], types.Typ[2], types.Typ[3]) {
			t.Fatalf("Expected to return 'func(<typeA>, <typeB>, ...<variadicType>) (<typeA>, <typeB>)'")
		}
	})
	t.Run("For Function Signatures with params (plus variadic) and return type, return errors when resolving params types (not the variadic one)", func(t *testing.T) {
		params := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", &fakeType{}),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
			types.NewVar(0, nil, "fieldD", types.NewSlice(types.Typ[4])),
		}...)
		returnTypes := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, params, returnTypes, true)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Function Signatures with params (plus variadic) and return type, return errors when resolving the variadic type", func(t *testing.T) {
		params := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[1]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
			types.NewVar(0, nil, "fieldD", types.NewSlice(&fakeType{})),
		}...)
		returnTypes := types.NewTuple([]*types.Var{
			types.NewVar(0, nil, "fieldA", types.Typ[1]),
			types.NewVar(0, nil, "fieldB", types.Typ[2]),
			types.NewVar(0, nil, "fieldC", types.Typ[3]),
		}...)
		type_ := types.NewSignature(nil, params, returnTypes, true)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Named types, should return the typeName if its a builtin type", func(t *testing.T) {
		typeName := types.NewTypeName(0, nil, "someName", types.Typ[1])
		type_ := types.NewNamed(typeName, nil, nil)
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != typeName.Name() {
			t.Fatalf("Expected to return the typeName")
		}
	})
	t.Run("For Named types that need import, should return 'pkgPath.typeName'", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath", "pkgName")
		typeName := types.NewTypeName(0, pkg, "someName", types.Typ[1])
		type_ := types.NewNamed(typeName, nil, nil)
		m := &mockGoImports{
			mockNeedImport: func(path string) bool {
				if path != pkg.Path() {
					t.Fatalf("Wrong package path")
				}
				return true
			},
			mockAliasFromPath: func(path string) string {
				if path != pkg.Path() {
					t.Fatalf("Wrong package path")
				}
				return pkg.Path()
			},
		}
		id, e := ResolveTypeIdentifier(type_, m)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != fmt.Sprintf("%s.%s", pkg.Path(), typeName.Name()) {
			t.Fatalf("Expected to return the 'pkgPath.typeName'")
		}
	})
	t.Run("For Named types that don't need import, should return 'typeName'", func(t *testing.T) {
		pkg := types.NewPackage("pkgPath", "pkgName")
		typeName := types.NewTypeName(0, pkg, "someName", types.Typ[1])
		type_ := types.NewNamed(typeName, nil, nil)
		m := &mockGoImports{
			mockNeedImport: func(path string) bool {
				if path != pkg.Path() {
					t.Fatalf("Wrong package path")
				}
				return false
			},
		}
		id, e := ResolveTypeIdentifier(type_, m)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != typeName.Name() {
			t.Fatalf("Expected to return the 'typeName'")
		}
	})
	t.Run("For empty Interfaces, should return 'interface{}'", func(t *testing.T) {
		type_ := types.NewInterfaceType(nil, nil)
		type_.Complete()
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "interface{}" {
			t.Fatalf("Expected to return 'interface{}'")
		}
	})
	t.Run("For Interfaces, should recursively call ResolveTypeIdentifier for every method, returning 'interface{<methodA>; <methodB>}'", func(t *testing.T) {
		methods := []*types.Func{
			types.NewFunc(0, nil, "methodA", types.NewSignature(nil, nil, nil, false)),
			types.NewFunc(0, nil, "methodB", types.NewSignature(nil, nil, nil, false)),
			types.NewFunc(0, nil, "methodC", types.NewSignature(nil, nil, nil, false)),
		}
		type_ := types.NewInterfaceType(methods, nil)
		type_.Complete()
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "interface{methodA(); methodB(); methodC()}" {
			t.Fatalf("Expected to return 'interface{<methodA>; <methodB>}'")
		}
	})
	t.Run("For Interfaces, return any errors when recursively resolving method types", func(t *testing.T) {
		methods := []*types.Func{
			types.NewFunc(0, nil, "methodA", types.NewSignature(nil, nil, nil, false)),
			types.NewFunc(0, nil, "methodB", types.NewSignature(nil, types.NewTuple(types.NewVar(0, nil, "fieldB", &fakeType{})), nil, false)),
			types.NewFunc(0, nil, "methodC", types.NewSignature(nil, nil, nil, false)),
		}
		type_ := types.NewInterfaceType(methods, nil)
		type_.Complete()
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
	t.Run("For Interfaces with embedded, should consider the embedded methods", func(t *testing.T) {
		embeddedMethods := []*types.Func{
			types.NewFunc(0, nil, "methodA", types.NewSignature(nil, nil, nil, false)),
			types.NewFunc(0, nil, "methodB", types.NewSignature(nil, nil, nil, false)),
		}
		methods := []*types.Func{
			types.NewFunc(0, nil, "methodC", types.NewSignature(nil, nil, nil, false)),
			types.NewFunc(0, nil, "methodD", types.NewSignature(nil, nil, nil, false)),
			types.NewFunc(0, nil, "methodE", types.NewSignature(nil, nil, nil, false)),
		}
		type_ := types.NewInterfaceType(methods, []types.Type{types.NewInterfaceType(embeddedMethods, nil)})
		type_.Complete()
		id, e := ResolveTypeIdentifier(type_, nil)
		if e != nil {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "interface{methodA(); methodB(); methodC(); methodD(); methodE()}" {
			t.Fatalf("Expected to return 'interface{<methodA>; <methodB>}'")
		}
	})
	t.Run("For unrecognized types, should return error", func(t *testing.T) {
		id, e := ResolveTypeIdentifier(&fakeType{}, nil)
		if e != UnexpectedTypeError {
			t.Fatalf("Error was expected to be nil")
		}
		if id != "" {
			t.Fatalf("When there's errors, the string is expected to be the zero value")
		}
	})
}
