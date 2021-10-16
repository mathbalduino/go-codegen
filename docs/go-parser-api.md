---
sidebar_position: 2
---

# *GoParser API

The library main type is the `*GoParser`. With this `struct`, you can call methods that iterate
over the `go` parsed code, allowing you to gather information and generate new code (be it `go`, `js`, 
`ts`, etc).

This type is a wrapper around the `golang.org/x/tools/go/packages` API, that at the time I used
was very hard to understand and use. After reading [this article](https://github.com/golang/example/blob/master/gotypes/go-types.md)
multiple times, I came with this idea. I strongly recommend you to read the same article, to fully
understand what we'll talk about (I'll assume that you've read).

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
	Focus 		*ParserFocus
	LogFlags 	uint
}
```

The `Tests`, `Dir`, `Env`, `Fset` and `BuildFlags` are directly forwarded to the underlying `packages.Config`
struct. If you need more info about it, see the [packages.Config docs](https://pkg.go.dev/golang.org/x/tools/go/packages#Config)

:::note
You will notice that the `packages.Config` has many fields that aren't present inside the `go-codegen.Config`.
This is on purpose. If you need to use one of the excluded fields, please [let me know](https://github.com/mathbalduino/go-codegen/issues/new)
:::

The `LogFlags` field is used to control the amount of information that the library will write to the `stdout`,
using the [LoggerCLI](https://mathbalduino.com.br/go-log/docs/advanced/logger_cli). The flags passed to 

FALAR AQUI SOBRE AS FLAGS DO LOGGERCLI, JUNTO COM A QUESTAO DO JSON + BEAUTIFY

### Focus

Sometimes, you will want to parse the `go` code but iterate only over some specific section. For example, you can
parse a package that contains many files, but just want to iterate over a specific file. To do it, you must give a
`*ParserFocus` to the `Config`, at `*GoParser` creation time.

The `root` package exports many functions that you can use to create a new `*ParserFocus`:

```go

```
