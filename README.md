# slick: a CI-ready shell language syntax checker

[![Docker Pulls](https://img.shields.io/docker/pulls/n4jm4/slick)](https://hub.docker.com/r/n4jm4/slick) [![GitHub Downloads](https://img.shields.io/github/downloads/mcandre/slick/total?logo=github)](https://github.com/mcandre/slick/releases) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/slick) [![Test](https://github.com/mcandre/slick/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/slick/actions/workflows/test.yml) [![Test-Futureproof-Dependencies](https://github.com/mcandre/slick/actions/workflows/test-futureproof-dependencies.yml/badge.svg)](https://github.com/mcandre/slick/actions/workflows/test-futureproof-dependencies.yml) [![Test-Futureproof-Language](https://github.com/mcandre/slick/actions/workflows/test-futureproof-language.yml/badge.svg)](https://github.com/mcandre/slick/actions/workflows/test-futureproof-language.yml) [![Test-Futureproof-OS](https://github.com/mcandre/slick/actions/workflows/test-futureproof-os.yml/badge.svg)](https://github.com/mcandre/slick/actions/workflows/test-futureproof-os.yml) [![license](https://img.shields.io/badge/license-BSD-3)](LICENSE.md) [![Donate](https://img.shields.io/badge/GUMROAD-36a9ae?style=flat&logo=gumroad&logoColor=white)](https://mcandre.gumroad.com/)

![squirtle squad](squirtle-squad.png)

# SUMMARY

slick provides a reliable syntax checker for pure POSIX `sh` scripts, as `sh` itself is often aliased with superset languages from `ash` to `zsh`.

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

* `sh` is hardly ever a bare bones POSIX sh interpreter on most UNIX systems, but usually soft linked to `bash`, `ksh`, `ash`, or even stranger things. So anyone genuinely interested in vetting their `#!/bin/sh` scripts for compliance risks getting false negative scans for scripts that actually contain bashisms, kshisms, and so on. By contrast, `slick` guarantees pure POSIX parsing, so that scripts are scanned consistently regardless of the particular environment configuration.
* `sh` is difficult to obtain in Windows. Cygwin-like environments are themselves difficult to setup. Should a unix, Linux, Windows, or other system desire syntax checking, `slick` is easy to obtain by gox ports, or through the wonderfully cross-platform Go toolchain.

# INSTALLATION

See [INSTALL.md](INSTALL.md).

# SEE ALSO

* [mvdan/sh](https://github.com/mvdan/sh) for POSIX sh parsing
* [stank](https://github.com/mcandre/stank) for more capable validation
* [todolint](https://github.com/mcandre/todolint) for finding software bugs

🌀
