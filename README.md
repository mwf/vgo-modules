# vgo-modules
Yet another vgo test project

## Updating modules, which dropped some dependencies

Consider the following flow:

A(v0.0.1) uses B(v0.0.1):
```
module github.com/mwf/vgo-modules/a

require github.com/mwf/vgo-modules/b v0.0.1
```

Then A(v0.1.0) drops the dependency.

Here we got a project `main` which uses A(v0.0.1) and we want to upgrade A to v0.1.0


```
cd `mktemp -d`
git clone https://github.com/mwf/vgo-modules .
echo "\n$(cat go.mod)"
echo "\n$(cat go.sum)"
go get github.com/mwf/vgo-modules/a@v0.1.0
echo "\n$(cat go.mod)"
```

Output:
```
module github.com/mwf/vgo-modules

require github.com/mwf/vgo-modules/a v0.0.1

github.com/mwf/vgo-modules/a v0.0.1 h1:BaK3DzgdA4LxVXd42qID7F9G+KFxQ4SL69hm+BNcjA0=
github.com/mwf/vgo-modules/a v0.0.1/go.mod h1:nK31MtN2feKpYgPBywpeTy/Om+1x3OxQHQAsyi95AQ0=
github.com/mwf/vgo-modules/b v0.0.1 h1:aEPPx6pEuQ1/Wu9SzeLkiYltlir2h1fDJmKRv1ckarY=
github.com/mwf/vgo-modules/b v0.0.1/go.mod h1:7kI8RJ5+IPwKWU1MrE/kTA9AdWBNzMmucdpAiWXATOs=

module github.com/mwf/vgo-modules

require (
    github.com/mwf/vgo-modules/a v0.1.0
    github.com/mwf/vgo-modules/b v0.0.1 // indirect
)
```

We get an unexpected indirect dependency `github.com/mwf/vgo-modules/b v0.0.1 // indirect` in `go.mod`, it seems to be taken from go.sum as a side-effect.
