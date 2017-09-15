# slick: a CI-ready shell language syntax checker

# EXAMPLES

```console
$ slick -n examples; echo "$?"
2017/09/14 17:46:09 examples/apples.bash:2:8: arrays are a bash feature
2017/09/14 17:46:09 examples/hello.sh:2:6: reached EOF without closing quote '
1

$ slick -help
  -help
        Show usage information
  -n    Validate syntax
  -version
        Show version information
```

# ABOUT

slick provides an alternative to `sh -n`, which is problematic for a number of minor reasons:

* `sh` is hardly ever a bare bones POSIX sh interpreter on most *nix systems, but usually soft linked to `bash`, `ksh`, `ash`, or even stranger things. So anyone genuinely interested in vetting their `#!/bin/sh` scripts for compliance risks getting false negative scans for scripts that actually contain bashisms, kshisms, and so on. By contrast, `slick` guarantees pure POSIX parsing, so that scripts are scanned consistently regardless of the particular environment configuration.
* `sh` is difficult to obtain in Windows. Cygwin-like environments are themselves difficult to setup. Should a unix, Linux, Windows, or other system desire syntax checking, `slick` is easy to obtain by gox ports, or through the wonderfully cross-platform Go toolchain.

# DOWNLOAD

https://github.com/mcandre/slick/releases

# DEVELOPMENT REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/slick/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/slick.git $GOPATH/src/github.com/mcandre/slick
$ cd $GOPATH/src/github.com/mcandre/slick
$ git submodule update --init --recursive
$ sh -c 'cd cmd/slick && go install'
```
