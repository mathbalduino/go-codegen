---
sidebar_position: 3
---

# Helpers

The library comes with some builtin functions that you can use to ease your code generation:

```go
func CallbackOnNamedType(fieldType types.Type, callback func(obj *types.Named)) error { ... }
func ResolveTypeIdentifier(t types.Type, pkgImports GoImports) (string, error) { ... }
func IdentifierToAsciiTypeName(typeIdentifier string) string { ... }
func IdentifierToTypeName(typeIdentifier string) string { ... }
func ObjectIsAccessible(obj types.Object, fromPackagePath string, logger LoggerCLI) bool { ... }
func TypeIsAccessible(t types.Type, fromPackagePath string) (bool, error) { ... }
```

These functions are a selection of functions that I always implement when I'm generating code in another libraries.
Feel free to [suggest new ones](https://github.com/mathbalduino/go-codegen/issues/new).

:::tip
Note that almost every helper function expects to receive a logger instance. These helpers are intended to be used inside
the `*GoParser` iteration methods (`IterateInterfaces` or `IterateStructs`), so you can just forward the logger received
via the callback arguments
:::

## CallbackOnNamedType

```go
func CallbackOnNamedType(fieldType types.Type, callback func(obj *types.Named)) error { ... }
```

This function will call the `callback` argument whenever it encounters some `NamedType` (if you don't know what it is,
read [this article](https://github.com/golang/example/blob/master/gotypes/go-types.md)).

:::note
If the `fieldType` argument has a `struct` or `interface` as it's underlying type, this function will recursively 
iterate over its `fields`/`methods`. The `callback` can be called more than just once
:::

## ResolveTypeIdentifier

```go
func ResolveTypeIdentifier(type_ types.Type, pkgImports GoImports) (string, error) { ... }
```

This function will recursively iterate over the `type_` argument, building its `string` identifier. At the end of its
execution, you will have something like this: `map[string]uint`, `struct{someField string}`, etc.

:::note
Since the function takes a list of package imports (the `pkgImports` argument), it can calculate when it's necessary
to include the name of the package from which some type comes from.

The final type identifier can be both `*someType` or `*pkgName.someType`, depending on the given `pkgImports` argument
:::


## IdentifierToAsciiTypeName

```go
func IdentifierToAsciiTypeName(typeIdentifier string) string { ... }
```

Sometimes, when generating code, you will need to name the generated types. Some symbols used to build type identifiers
cannot be used in type names. The `map[string]int` cannot be converted to a valid type name without modifying some chars
(the "[" and "]" chars are forbidden).

This function will take the type identifier and convert it to a valid type name. The above `map[string]int` example will
be converted to something like `mapstringint`. The final result is not that readable, but it works.

If you want some readable generated type name, take a look at the next function.

## IdentifierToTypeName

```go
func IdentifierToTypeName(typeIdentifier string) string { ... }
```

This function is very similar to the `IdentifierToAsciiTypeName` in the sense that it converts a type identifier into a
valid type name. The difference is that it generates type names with custom unicode symbols that resemble the replaced
chars (I tried **_very_** hard to keep the generated type names readable).

The type identifier `*int`, for example, will be converted to `ᕽint` (using the custom [Canadian Syllabics Hk](https://unicode-table.com/en/157D/)).

:::caution
If you need that your types contain only ASCII type names, don't use this function, use the `IdentifierToAsciiTypeName`
instead
:::

## ObjectIsAccessible

```go
func ObjectIsAccessible(obj types.Object, fromPackagePath string, logger LoggerCLI) bool { ... }
```

If you want to test some `types.Object` to see if you can access it from some specific package, you can give it to this
function, with some package information.

The object `pkgA.TypeA` is accessible within the package `pkgB`, because it is exported. The `pkgA.typeB` object is 
accessible within the package `pkgA`, even being private, because it is the same package. And so on...

## TypeIsAccessible

```go
func TypeIsAccessible(t types.Type, fromPackagePath string, logger LoggerCLI) (bool, error) { ... }
```

Very similar to the `ObjectIsAccessible` function, but with `types.Type` instead of `types.Object`.
