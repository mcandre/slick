# DEVELOPMENT GUIDE

slick follows standard, cargo based operations for compiling and unit testing Go code.

For advanced operations, such as linting, we further supplement with some software industry tools.

# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* Provision additional dev tools with `make`

## Recommended

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)

# AUDIT

```console
$ mage audit
```

# INSTALL

```console
$ mage install
```

# UNINSTALL

```console
$ mage uninstall
```

# LINT

```console
$ mage lint
```

# TEST

```console
$ mage test
```

# CLEAN

```console
$ mage clean
```
