---
sidebar_position: 4
---

# *GoFile

If you want to generate GO files, I strongly recommend that you use the `goFile` package. This package handles the import
list of the file automatically (trust me, this is very welcome), and some minor stuff (code formatting, persistence).

Every `*GoFile` has an embedded `*GoImports`, allowing you to manipulate the import list directly, if you want to. The 
`*GoFile` can be created using the exported functions from the `goFile` package, while the `*GoImports` can be created
using the exported functions from the `goFile.goImports` package.

## *GoImports

This is the most important component of the `*GoFile` struct, so you need to understand it before going to the `*GoFile`
API (that is, basically, just `get`/`set` methods).

### New

You can create a new `*GoImports` instance using the `New` function, described below:

```go
func New(packagePath string) *GoImports { ... }
```

This function takes the package path that will hold the import list (represented as the `*GoImports` itself). If you're
using `go modules`, the string will be something like `<repository_url>/<path>/<to>/<pkg>`. 

Example: if we're generating code that will belong to the `go-codegen` library, the `packagePath` argument will be 
equal to `github.com/mathbalduino/go-codegen` (for the root package), or something like 
`github.com/mathbalduino/go-codegen/<sub_pkg>` (for another package inside the lib).

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
`suggestedAlias -> newAlias`. The returned map can be empty if there's no clashes.

### NeedImport


