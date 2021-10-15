---
sidebar_position: 1
---

# Introduction

`go-codegen` is a library that acts as a wrapper around `golang.org/x/tools`, providing an API
that you can use to parse and generate `go` code.

:::tip
I strongly recommend that you read [this document about GO types](https://github.com/golang/example/blob/master/gotypes/go-types.md),
to get a deeper understanding about it
:::

## Getting Started

You will need to install go-codegen before starting. To do it, execute the following command:

```sh test
go get github.com/mathbalduino/go-codegen
```

As already said, this library acts as a wrapper around `golang.org/x/tools`, so you will need to create a new
`GoParser` instance (that calls `golang.org/x/tools/go/packages.Load`) in order to use it's utility methods:

```go
package main

import (
	"github.com/mathbalduino/go-codegen"
)

func main() {
	
}
```

[ ] Falar sobre IterateInterfaces
[ ] Falar sobre IterateStructs
[ ] Falar sobre Config
[ ] Mostrar como criar um novo parser (colocar referencia para o packages.Load)
[ ] Falar sobre os Focus

[ ] Falar sobre a abstração sobre GO files
[ ] Abstração sobre TS files
[ ] Falar sobre cada um dos helpers

