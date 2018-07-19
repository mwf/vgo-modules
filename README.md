# vgo-modules
Yet another vgo test project

## Upgrading modules, which dropped some dependencies

[golang/go#26474](https://github.com/golang/go/issues/26474)

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


## Downgrading modules, which updated some dependencies

[golang/go#26481](https://github.com/golang/go/issues/26481)

Consider the following flow:

C(v0.1.0) uses A(v0.1.0):
```
module github.com/mwf/vgo-modules/c

require github.com/mwf/vgo-modules/a v0.1.0
```

Then A(v0.2.0) is released and C(v0.2.0) depending on new A(v0.2.0).

We got a project `main` which uses C(v0.1.0) and we want to upgrade C to v0.2.0


```
cd `mktemp -d`
git clone https://github.com/mwf/vgo-modules .
cd c_user
echo "\n$(cat go.mod)"
echo "\n$(cat go.sum)"
go get github.com/mwf/vgo-modules/c@v0.2.0
echo "\n$(cat go.mod)"
```

Output:
```
module github.com/mwf/vgo-modules/c_user

require github.com/mwf/vgo-modules/c v0.1.0

github.com/mwf/vgo-modules/a v0.1.0 h1:QD+ijrXwAYk4CMTOqEAzcogJoF4zLfit2alZVmT80EM=
github.com/mwf/vgo-modules/a v0.1.0/go.mod h1:XGJvSKC62ygHgRNmDf4RqXyp0zngXIJMLc6BaSfyTfI=
github.com/mwf/vgo-modules/c v0.1.0 h1:db5PK5vzxlLviLcUsd1YCaLjG3+35hw6OMKSdVzqPMo=
github.com/mwf/vgo-modules/c v0.1.0/go.mod h1:qfiW4sL7v2Bnu6kBwvwxeVylzuPa6cxOldlsx2NCIRI=

module github.com/mwf/vgo-modules/c_user

require github.com/mwf/vgo-modules/c v0.2.0
```

All seems OK, than we run tests and understand that C(v0.2.0) is broken - `go test .`
```
--- FAIL: TestC (0.00s)
    main_test.go:12: c.CA: got CAaaaa, expected: CA
FAIL
FAIL    github.com/mwf/vgo-modules/c_user   0.014s
```

We decide to downgrade C back to v0.1.0 used before:

```
go get github.com/mwf/vgo-modules/c@v0.1.0
echo "\n$(cat go.mod)"
go test .
```

After the downgrade we get unexpected `github.com/mwf/vgo-modules/a v0.2.0 // indirect` and tests still fail.

```
module github.com/mwf/vgo-modules/c_user

require (
    github.com/mwf/vgo-modules/a v0.2.0 // indirect
    github.com/mwf/vgo-modules/c v0.1.0
)

--- FAIL: TestC (0.00s)
    main_test.go:12: c.CA: got CAaaaa, expected: CA
FAIL
FAIL    github.com/mwf/vgo-modules/c_user   0.018s
```
