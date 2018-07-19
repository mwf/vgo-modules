# vgo-modules
Yet another vgo test project

## Refactoring modules

[go#26250](https://github.com/golang/go/issues/26250) - **fixed**

Let's assume we are in active development phase with versions `v0.x`.

At some point we decide to move some package to a standalone module, to split the dependencies and to have better semantic structure.

Thus `cmd` package is introduced to a separate module `github.com/mwf/vgo-modules/cmd`.

But it turns out you can't use both `github.com/mwf/vgo-modules/cmd` and `github.com/mwf/vgo-modules` at the same time.
Have a look at [example](./example) project.

```
cd ./example && vgo get github.com/mwf/vgo-modules/cmd

vgo: import "github.com/mwf/vgo-modules/cmd": found in both github.com/mwf/vgo-modules v0.0.2 and github.com/mwf/vgo-modules/cmd v0.0.1
```
