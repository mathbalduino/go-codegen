---
sidebar_position: 1
---

# Introduction

`go-codegen` is a library that acts as a wrapper around `golang.org/x/tools`, providing an API
that you can use to parse and generate `go` code. It's part of my personal stack, and I use it
to ease code generation inside another libraries.

This library is something like a compilation of the functions, abstractions, etc, all the code
that I usually need when generating code. Feel free to contribute, if you want to.

:::info
I don't want to use this library to handle the _content_ of the generated files, just to provide
tools that ease the process of code generation (code parsing, type inspection, file abstraction, etc).

Don't expect to see code related to template generation, etc. This is intended to be implemented by 
yourself when generating your code
:::

:::tip
I strongly recommend that you read [this document about GO types](https://github.com/golang/example/blob/master/gotypes/go-types.md),
to get a deeper understanding about what we'll talk about in the next chapters. I'll assume will have read.
:::

## Getting Started

You will need to install `go-codegen` before starting. To do it, execute the following command (inside your `go module`
folder):

```sh test
go get github.com/mathbalduino/go-codegen
```

As already said, this library acts as a wrapper around `golang.org/x/tools`, so you will need to create a new
`GoParser` instance (that calls `golang.org/x/tools/go/packages.Load`) in order to use its utility methods. Example:

```go
package main

import (
	"github.com/mathbalduino/go-codegen"
	"go/types"
)

func main() {
	config := parser.Config{
		// ... your configuration ...
		//
	}
	goParser, e := parser.NewGoParser("<your_pattern>", config)
	if e != nil {
		panic(e)
	}
	
	e = goParser.IterateStructs(func(struct_ *types.TypeName, logger parser.LoggerCLI) error { 
		// This method will be called once for every struct inside the parsed GO code
		// Use the given 'struct_' param to generate your code

		// If you want to stop the iteration, return a non-nil error below
		// This error will be forwarded to the caller of 'IterateStructs'
		return nil
	})
	if e != nil {
		panic(e)
	}
}
```

After you get a working instance of `*GoParser`, you can call methods that will iterate over the
parsed code. With this information, you can generate new code. Currently, you can iterate over `structs`
and `interfaces`. For more info, see the [*GoParser API](go-parser-api.md) section.

Don't worry about the configuration, string pattern, etc. All these concepts will be explored in details in the
next chapters.

## Files abstraction

If you want to generate `go` or `typescript` code, the library comes with builtin support for this kind
of files, abstracting the file import list, code formatting and persistence (you can use the same API for
`ts` files). Example:

```go
package main

import (
	"github.com/mathbalduino/go-codegen/goFile"
)

func main() {
	// ... your goParser instance creation
	
	f := goFile.New("filename", "packageName", "destination/package/import/path")
	e := goParser.IterateStructs(func(struct_ *types.TypeName, logger parser.LoggerCLI) error {
		generatedCode := generateCodeUsingTemplate(struct_) // or, whatever you want
		f.AddCode(generatedCode)
		return nil
	})
	if e != nil {
		panic(e)
	}
	
	// <title> usually is the name of the library that
	// generated the code
	// Example: "Code generated by library/import/path v1.2.3"
	e = f.Save("<title>", "save/to/folder/x")
	if e != nil {
		panic(e)
	}
	// If there's no error, the file has been written to disk
}
```

## Focus

The `*GoParser`, by default, will iterate over the entire parsed code. You can control this behavior by giving
a `*Focus` to the `Config`, at the `*GoParser` creation, that tells to the `*GoParser`
to iterate only over some `typeName`, `filePath`, etc (for details, see the subsection [Focus](./go-parser-api.md#focus)):

```go
func FocusPackagePath(packagePath string) *Focus { ... }
func FocusTypeName(typeName string) *Focus { ... }
// ...
```

Keep reading the docs for details...