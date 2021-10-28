---
sidebar_position: 5
---

# *TsFile

The abstraction for `typescript` files is very similar to the one used for `go` files. There's only three differences:

- `typescript` supports named imports
- default imports too
- and there's no "package" associated, only `filesystem` path

:::note
Why `ts`? Well, I was building a mobile app and needed to generate code based on some `go` structs (from the server-side)
:::

## AddNamedImport

```go
func (i *TsImports) AddNamedImport(namedImport string, path string) error { ... }
```

Use this method to add named imports to our `ts` file. Note that this method returns an `error` interface, not a string.
Currently, you cannot use aliases to rename the named imports (if you need it, please [let me know](https://github.com/mathbalduino/go-codegen/issues/new)).

If this method returns an error, the given import cannot be added to the `ts` file. If returns nil, the given named import
will **never** be changed.

Named imports are in the form of: 

```ts
import { NamedImport } from 'some/file/path'
```

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
