---
sidebar_position: 2
---

# *GoParser API

The library main type is the `*GoParser`. With this `struct`, you can call methods that iterate
over the `go` parsed code, allowing you to gather information and generate new code (be it `go`, `js`, 
`ts`, etc).

This type is a wrapper around the `golang.org/x/tools/go/packages` API, that at the time I used
was very hard to understand and use. After reading [this article](https://github.com/golang/example/blob/master/gotypes/go-types.md)
multiple times, I came with this idea.

:::danger
I strongly recommend you to read the above article, to fully understand what we'll talk about. I'll assume that you've
read
:::

## Creating a *GoParser instance

You can create a new `*GoParser` instance using the `NewGoParser` function, exported by the root package.
It takes a `pattern` `string` and a `Config` `struct`.

```go
func NewGoParser(pattern string, config Config) (*GoParser, error) { ... }
```

The `pattern` `string` is forwarded directly to the `packages.Load` function. If you want to see details
about what is this string, go to the official [packages docs](https://pkg.go.dev/golang.org/x/tools/go/packages#pkg-overview).

:::info
When you create the `*GoParser` instance (calling the `NewGoParser` function), the `packages.Load` function is called 
to parse the code. Depending on the size of the code to parse, this function call can be expensive.
:::

:::note
Any errors that occur when calling `packages.Load` will be directly returned to the caller of `NewGoParser`,
with a nil `*GoParser` reference.
:::

## Configuration

Before creating a new `*GoParser` instance, you will need to read about the `Config` struct, that 
acts as a wrapper around the `packages.Config`, plus some other info:

```go
type Config struct {
	Tests 		bool
	Dir 		string
	Env 		[]string
	Fset 		*token.FileSet
	BuildFlags 	[]string
	Focus 		*Focus
	Logger      LoggerCLI
	LogFlags 	uint64
}
```

The `Tests`, `Dir`, `Env`, `Fset` and `BuildFlags` are directly forwarded to the underlying `packages.Config`
struct. If you need to see more info about it, see the [packages.Config docs](https://pkg.go.dev/golang.org/x/tools/go/packages#Config)

:::note
You will notice that the `packages.Config` has many fields that aren't present inside the `go-codegen.Config`.
This is on purpose. If you need to use one of the excluded fields, please [let me know](https://github.com/mathbalduino/go-codegen/issues/new)
:::

The `Logger` field is used to allow the library user to pass an already created logger instance, instead of relying on
the one created by the library itself. This is extremely useful, since this library will usually be used to compose
other libraries, the logger instance usually already exists.

The `LogFlags` field is used to control the amount of information that the library will write to the `stdout`,
using the [LoggerCLI](https://mathbalduino.com.br/go-log/docs/advanced/logger_cli) (another library that belongs to
my personal stack). The flags will be directly forwarded to it. Note that since this lib uses the `LoggerCLI`, it's
possible to use the `beautify` package, from the `LoggerCLI` itself, to pretty-print the generated output. You can see 
the available flags at the [go-log](https://mathbalduino.com.br/go-log/docs/basic-concepts/configuration#lvlsenabled-usage)
official documentation (note that the `go-codegen` library re-exports these flags, so you don't need to point to `go-log`
directly).

:::tip
In addition to the original `go-log` flags, there's an extra one: `LogJSON`. This extra flag belongs to the `go-codegen`
itself, and is used to control whether the logs are converted (or not) to `json`, before being sent to the `stdout`
:::

:::note
If you set the value of the `Logger` field, the `LogFlags` field will be completely ignored, since it is only used when
the library needs to dynamically create a new `LoggerCLI` instance
:::

### Focus

Sometimes, you will want to parse the `go` code but iterate only over some specific thing. You can parse a 
package that contains many files, but just want to iterate over a specific file, for example. To do it, you 
must give a `*Focus` to the `Config`, at `*GoParser` creation time.

The `root` package exports three functions that you can use to create a new `*Focus`:

```go
func FocusPackagePath(packagePath string) *Focus { ... }
func FocusFilePath(filePath string) *Focus { ... }
func FocusTypeName(typeName string) *Focus { ... }
```

:::caution
It's very important to note that the string received by the above functions is considered a `RegExp`. It's *not* a simple
string. It means that if you pass the `"Input"` `string` to the `FocusTypeName` function, the following `typenames` will match:
`InputStruct`, `Input`, `MyInputStruct`, etc

If you want an exact match, use the `^` and `$` characters, that set the start/end of the string, respectively: `"^Input$"`. It is just
an example, but you can use all the features that you would use inside a `regexp.MatchString()` method call
:::

Example: if you want to parse the `codeToParse.go` file, that contains 3 `structs`, but iterate only over the
`StructB`, you can do something like this (consider `workdir` to be anything):

```go title="<workdir>/codeToParse.go"
package main

type StructA struct {
	FieldA string
	FieldB string
}

type StructB struct {
	FieldC string
	FieldD string
}

type StructC struct {
	FieldE string
	FieldF string
}
```

```go title="<workdir>/main.go"
package main

import "github.com/mathbalduino/go-codegen"

func main() {
	config := parser.Config{
		Focus: parser.FocusTypeName("StructB"),
	}
	// ...
}
```

With this `Config`, the iterator will skip every type name that's different from `"StructB"`.

You can combine more than one `*Focus` using the `MergeFocus` function, exported by the root package:

```go
func MergeFocus(f1 *Focus, f2 *Focus) *Focus { ... }
```

This function will copy the first argument to a third `*Focus` and override it with the values of the second one. This
way, you can have a `*Focus` that filters by `typename` and `package`, simultaneously, for example.

## Iterate interfaces

After the `*GoParser` instantiation (and code parsing), you can call the method below to iterate over `interfaces`:

```go
// just an alias
type InterfacesIterator = func(interface_ *types.TypeName, logger LoggerCLI) error

func (p *GoParser) IterateInterfaces(callback InterfacesIterator, optionalLogger ...LoggerCLI) error { ... }
```

:::note
Note that you have the option to give a new `LoggerCLI` instance, at the moment you call `IterateInterfaces`. It's very
useful when you're using the `beautify` package, from [go-log](https://mathbalduino.com.br/go-log), and want to customize
the log display

You can completely omit it (since it is a variadic), in which case the logger of the `*GoParser` instance will be used instead
:::

With this method, you pass a callback function that will be executed once for every `interface` type inside the
parsed code:

```go title="<workdir>/codeToParse.go"
package main

type StructA struct {
	FieldA string
	FieldB string
}

type InterfaceA interface {
	MethodA()
	MethodB()
}

// private
type interfaceB interface {
	methodC()
	methodD()
}
```

```go title="<workdir>/main.go"
package main

import (
	"fmt"
	"github.com/mathbalduino/go-codegen"
	"go/types"
)

func main() {
	config := parser.Config{}
	
	// Assuming the code is being executed at <workdir>
	p, e := parser.NewGoParser("./", config)
	if e != nil {
		panic(e)
	}
	
	// stdout:
	// 		InterfaceA
	// 		interfaceB
	e = p.IterateInterfaces(func(interface_ *types.TypeName, logger parser.LoggerCLI) error {
		fmt.Println(interface_.Name())
		return nil
	})
	if e != nil {
		panic(e)
	}
}
```

:::note
The example above will parse the `<workdir>/main.go` too, but it will be completely ignored, since there's no
`interfaces` inside it
:::

:::note
If you return some error from the callback function, this error will be returned to the caller of
`IterateInterfaces`, stopping its execution.
:::


## Iterate structs

After the `*GoParser` instantiation (and code parsing), you can call the method below to iterate over `structs`:

```go
// just an alias
type StructsIterator = func(struct_ *types.TypeName, logger LoggerCLI) error

func (p *GoParser) IterateStructs(callback StructsIterator, optionalLogger ...LoggerCLI) error { ... }
```

:::note
Note that you have the option to give a new `LoggerCLI` instance, at the moment you call `IterateInterfaces`. It's very
useful when you're using the `beautify` package, from [go-log](https://mathbalduino.com.br/go-log), and want to customize
the log display

You can completely omit it (since it is a variadic), in which case the logger of the `*GoParser` instance will be used instead
:::

With this method, you pass a callback function that will be executed once for every `struct` type inside the
parsed code:

```go title="<workdir>/codeToParse.go"
package main

type StructA struct {
	FieldA string
	FieldB string
}

// private
type structB interface {
	fieldA string
	fieldB string
}

type interfaceB interface {
	methodC()
	methodD()
}
```

```go title="<workdir>/main.go"
package main

import (
	"fmt"
	"github.com/mathbalduino/go-codegen"
	"go/types"
)

func main() {
	config := parser.Config{}
	
	// Assuming the code is being executed at <workdir>
	p, e := parser.NewGoParser("./", config)
	if e != nil {
		panic(e)
	}
	
	// stdout:
	// 		StructA
	// 		structB
	e = p.IterateStructs(func(struct_ *types.TypeName, logger parser.LoggerCLI) error {
		fmt.Println(struct_.Name())
		return nil
	})
	if e != nil {
		panic(e)
	}
}
```

:::note
The example above will parse the `<workdir>/main.go` too, but it will be completely ignored, since there's no
`structs` inside it
:::

:::note
If you return some error from the callback function, this error will be returned to the caller of 
`IterateStructs`, stopping its execution.
:::

## Other iterators

The library comes with other iterators, but I keep them private to reduce the API size. Take a look:

```go
func (p *GoParser) iterateTypeNames(callback typeNamesIterator) error { ... }
func (p *GoParser) iterateFiles(callback filesIterator) error { ... }
func (p *GoParser) iteratePackages(callback packagesIterator) error { ... }
```

Internally, these methods are used to compose the exported ones.

:::note
If you think it's interesting to export any of these methods, [let me know](https://github.com/mathbalduino/go-codegen/issues/new)
:::

## Some notes

As you've seemed, this library doesn't abstract the most primitive types from `go/types`. When you call `IterateInterfaces`
or `IterateStructs`, the `*types.TypeName` that the callback receives has all the necessary information about the 
underlying type (name, fields, methods, etc).

The library iterates over the parsed code using a `single-threaded` strategy. This can be improved in the future, but for
now I don't see the necessity. If you think I'm wrong, please [let me know](https://github.com/mathbalduino/go-codegen/issues/new).
