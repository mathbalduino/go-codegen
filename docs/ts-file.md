---
sidebar_position: 5
---

# *TsFile

The abstraction for `typescript` files is very similar to the one used for `go` files. There's only three differences:

- `typescript` supports named imports
- default imports too
- and there's no "package" associated, only `filesystem`/`transpiler` path

:::note
Why `ts`? Well, I was building a mobile app and needed to generate `client` code based on some `go` structs (from the 
server-side)
:::

:::note
I will show only the documentation for the differences between `*TsFile` and `*GoFile`. If you want details about something
that is not describe below, look at the [*GoFile](./go-file.md) docs
:::

## *TsImports

## AddNamedImport

```go
func (i *TsImports) AddNamedImport(namedImport string, path string) error { ... }
```

Use this method to add named imports to our `ts` file. Note that this method returns an `error` interface, not a string.

If this method returns an error, the given import cannot be added to the `ts` file. If returns nil, the given named import
will be used and **never** changed.

Named imports are in the form of: 

```ts
import { NamedImport } from 'some/file/path'
```

:::note
Currently, it's no possible to use aliases in the named imports. If you need this feature, please [let me know](https://github.com/mathbalduino/go-codegen/issues/new)
:::

## AddDefaultImport

```go
func (i *TsImports) AddDefaultImport(defaultImport string, path string) error { ... }
```

Use this method to add a new default import to your `ts` file. Since default imports can be imported using any name, you
don't need aliases.

If this method returns an error, the given import cannot be added to the `ts` file. If returns nil, the given default import
will **never** be changed.

Default imports are in the form of:

```ts
import DefaultImport from 'some/file/path'
```
