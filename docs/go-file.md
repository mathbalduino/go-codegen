---
sidebar_position: 4
---

# *GoFile

If you want to generate `go` files, I strongly recommend that you use the `goFile` package. This package handles the import
list of the file automatically (trust me, this is very welcome), and some minor stuff (code formatting, persistence).

Every `*GoFile` has an embedded `*GoImports`, allowing you to manipulate the import list directly, if you want to. The 
`*GoFile` can be created using the exported functions from the `goFile` package, while the `*GoImports` can be created
using the exported functions from the `goFile.goImports` package.

## *GoImports

This is the most important component of the `*GoFile` struct, so you need to understand it before going to the `*GoFile`
API (that is, basically, just `get`/`set` methods). You will find the `*GoImports` code under the [goImports package](https://github.com/mathbalduino/go-codegen/tree/main/goFile/goImports).

The `*GoImports` struct represents a list of imports made by some package. This list can be the import list inside some 
file (that belongs to a package, naturally), for example. The important thing is: `*GoImports` stores information about
the imports being made by some package.

The `*GoImports` struct is composed by two fields: `packagePath string` and `imports map[string]string`. 
More details [here](https://github.com/mathbalduino/go-codegen/blob/main/goFile/goImports/new.go).

:::note
I could've just used the `golang.org/x/tools/imports.Process()` function to insert the missing imports, but sometimes I've
had some issues, so I decided to implement it myself (I still use the `Process()` function, but just to format the code)
:::

### New

You can create a new `*GoImports` instance using the `New` function, described below:

```go
func New(packagePath string) *GoImports { ... }
```

This function takes the package path of the package that owns the import list (represented as the `*GoImports` itself). 
If you're using `go modules`, the string can be something like `<repository_url>/<path>/<to>/<pkg>`. 

Example: if we're generating code that will belong to the `go-codegen` library, the `packagePath` argument will be 
equal to `"github.com/mathbalduino/go-codegen"` (for the root package), or something like 
`"github.com/mathbalduino/go-codegen/<pkg_name>"` (for another package inside the lib).

### AddImport

```go
func (i *GoImports) AddImport(suggestedAlias, packagePath string) string { ... }
```

When adding imports to any `go file`, you can use aliases to help distinguish imports with the same root package name.
This method receives a string representing the desired alias and the imported package path. Note that the alias is just
a suggestion, since it will be changed if there's a clash with another alias in the list.

The returned string will be the final alias used to identify the given import. If the import is already in the list, its
alias will be returned.

:::note
The aliases will always be used when generating code, even if they're not required (`import fmt "fmt"`, for example)
:::

### AliasFromPath

```go
func (i *GoImports) AliasFromPath(packagePath string) string { ... }
```

If you want to query the alias used to represent some imported package, you can give its package path to this method. If
the returned `string` is empty, there's no import with the given package path inside the import list.

### MergeImports

```go
func (i *GoImports) MergeImports(other *GoImports) map[string]string { ... }
```

If you have two import lists and want to merge them, this method will do it for you. Note that this method will just
iterate over the given import list and try to add each one, using the `AddImport` method.

The returned map represents the aliases that changed when being added, in case of clashes. The map is in the form of 
`originalAlias -> newAlias`. The returned map can be empty if there's no clashes.

:::caution
If you have two import lists that belong to two different package paths, you can get a `panic` call. Example:

- Say that you have two import lists: package A and package B
- The package B imports the package A
- You merge the import list of the package B into the import list of the package A
- Since the package A cannot import itself, panic will be called 

More details [here](https://github.com/mathbalduino/go-codegen/blob/20cc90dac2de869cd647272abfabf5333e692553/goFile/goImports/addImport.go#L20)
:::

### NeedImport

```go
func (i *GoImports) NeedImport(otherPackagePath string) bool { ... }
```

Sometimes, you will need to test if some package needs to be imported in order to be accessible from the package that owns
the import list. This method will just compare the strings (the one given to the `New` function and the one given to the
method itself).

### PackagePath

```go
func (i *GoImports) PackagePath() string { ... }
```

Just a getter to the string that represents the package path of the package that owns the import list.

### SourceCode

```go
func (i *GoImports) SourceCode() string { ... }
```

When you have finished adding imports to the import list, you can use this method to generate valid `go` code that 
can be attached to some file.

:::note
The returned string is just the "import" section, not the entire file, so you cannot compile it right away
:::

### Example

We will create an import list of some file inside the `example` package (so, the generated code will be put inside the
`example` package):

```go
package main

import (
	"fmt"
	"github.com/mathbalduino/go-codegen/goFile/goImports"
)

func main() {
	importList := goImports.New("example")
	importList.AddImport("fmt", "fmt")
	importList.AddImport("time", "time")
	importList.AddImport("logger", "github.com/mathbalduino/go-log")

	anotherList := goImports.New("anotherPkg")
	anotherList.AddImport("sync", "sync")
	anotherList.AddImport("fmt", "fmt")
	// anotherList.Add("example", "example") // If enabled, causes panic

	importList.MergeImports(anotherList)
	fmt.Println(importList.SourceCode())
	// stdout (order not guaranteed):
	//
	// import (
	//    fmt "fmt"
	//    time "time"
	//    logger "github.com/mathbalduino/go-log"
	//    sync "sync"
	// )
}
```

## *GoFile

The `*GoFile` struct holds information about the `name`, the `packageName` (the name, not the path), the `sourceCode` (
the body of the file) and the `importList` (as an embedded `*GoImports` struct). You will find the `*GoFile` code under
the [goFile package](https://github.com/mathbalduino/go-codegen/tree/main/goFile).

This struct will have all the methods from `*GoImports`, plus the ones describe below.

### New

```go
func New(filename, packageName, packagePath string) *GoFile { ... }
```

This function will create a new `go` file instance. Note that the `filename` will receive some suffix (copyright and
file extension), so you don't need to pass `"myfile.go"`, pass just `"myfile"` instead.

The `packageName` and `packagePath` are the name and the import path of the package that the file will belong to, after
persisted. If you're creating a file for the `go-codegen` library root package, the arguments will be:

```go
New("exampleFile", "parser", "github.com/mathbalduino/go-codegen")
```

### AddCode

```go
func (f *GoFile) AddCode(newSourceCode string) { ... }
```

This method will just append the given string to the file body

:::caution
Don't use this method to add code related to the package information (`package <pkgName>` keyword), or code related to
the file import list (use the `*GoImports` API to do this). These things will be handled automatically
:::

### Name

```go
func (f *GoFile) Name() string { ... }
```

Just a getter to the name of the file. Note that the returned string will contain the copyright and the file extension.

### PackageName

```go
func (f *GoFile) PackageName() string { ... }
```

Just a getter to the package that file belongs to.

### Save

```go
func (f *GoFile) Save(headerTitle, folder string) error { ... }
```

This is one of the most important method of the `*GoFile` API. This method is the responsible for building the content of
the file and persisting it to the `filesystem`.

Note that this method will call the `SourceCode` method, in order to build the file content, and use the `os.Create`, 
`os.Write` and `os.Close`.

The `headerTitle` is just a string that will be put inside the file copyright comment section (right at the header). 
Usually, I use something like `"<library>/<import>/<path> v1.0.0"`, but you can use anything.

### SourceCode

```go
func (f *GoFile) SourceCode(headerTitle, filepath string) ([]byte, error) { ... }
```

This method will build the content of the file, letting it ready to be compiled.

The `headerTitle` arg is just a string that will be put inside the file copyright comment section (right at the header).
Usually, I use something like `"<library>/<import>/<path> v1.0.0"`, but you can use anything.

The `filepath` arg is the absolute `filesystem` path that the file will be persisted. It is used to process the import
list of the file and format it (the underlying build tool needs this to get the context).

### Example

```go
package main

import (
	"github.com/mathbalduino/go-codegen/goFile"
)

func main() {
	newFile := goFile.New("exampleFile", "parser", "github.com/mathbalduino/go-log")
	newFile.AddImport("fmt", "fmt")
	newFile.AddImport("time", "time")
	
	newFile.AddCode("func exampleFn() { fmt.Printf(\"Hello world: %s\\n\", exampleConstant) }")
	newFile.AddCode("const exampleConstant = \"go-codegen\"")

	newFile.Save("docs example", "/") // root folder
	// File content:
	//    package parser
	//
	//    import (
	//       fmt "fmt"
	//       time "time"
	//    )
	//
	//    func exampleFn() { fmt.Printf("Hello world: %s\n", exampleConstant) }
	//
	//    const exampleConstant = "go-codegen"
}
```

