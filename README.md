# goxcart: Go port all the things!

![goxcart-logo](https://raw.githubusercontent.com/mcandre/goxcart/master/goxcart.png)

# EXAMPLE

```console
$ cd test
$ goxcart -output bin -banner "0.0.1" -commands ./cmd/...

$ tree bin
bin
└── hello-0.0.1
    ├── android
    │   └── 386
    │       └── hello
    ├── darwin
    │   ├── 386
    │   │   └── hello
    │   └── amd64
    │       └── hello
    ├── dragonfly
    │   └── amd64
    │       └── hello
    ├── freebsd
    │   ├── 386
    │   │   └── hello
    │   ├── amd64
    │   │   └── hello
    │   └── arm
    │       └── hello
    ├── linux-glibc
    │   ├── 386
    │   │   └── hello
    │   ├── amd64
    │   │   └── hello
    │   ├── arm
    │   │   └── hello
    │   ├── arm64
    │   │   └── hello
    │   ├── mips
    │   │   └── hello
    │   ├── mips64
    │   │   └── hello
    │   ├── mips64le
    │   │   └── hello
    │   ├── mipsle
    │   │   └── hello
    │   ├── ppc64
    │   │   └── hello
    │   ├── ppc64le
    │   │   └── hello
    │   └── s390x
    │       └── hello
    ├── linux-musl
    │   └── amd64
    │       └── hello
    ├── nacl
    │   ├── amd64p32
    │   │   └── hello
    │   └── arm
    │       └── hello
    ├── netbsd
    │   ├── 386
    │   │   └── hello
    │   ├── amd64
    │   │   └── hello
    │   └── arm
    │       └── hello
    ├── openbsd
    │   ├── 386
    │   │   └── hello
    │   └── amd64
    │       └── hello
    ├── plan9
    │   ├── 386
    │   │   └── hello
    │   ├── amd64
    │   │   └── hello
    │   └── arm
    │       └── hello
    ├── solaris
    │   └── amd64
    │       └── hello
    └── windows
        ├── 386
        │   └── hello.exe
        └── amd64
            └── hello.exe
```

# ABOUT

goxcart helps Go developers build programs for more platforms. While Go and gox already provide a lot of support for basic cross-platform builds, goxcart takes this further, enabling more target tuples than are currently available for host-run builds. For example, goxcart can build Go binaries targeting:

* android
* dragonfly
* linux-musl
* nacl
* plan9
* solaris

goxcart does this by building applications inside reusable Docker containers configured for these additional targets.

# RUNTIME REQUIREMENTS

* [Docker](https://www.docker.com/)

## Recommended

* [tree](https://linux.die.net/man/1/tree)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.9+
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/goxcart/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/goxcart.git $GOPATH/src/github.com/mcandre/goxcart
$ cd $GOPATH/src/github.com/mcandre/goxcart
$ git submodule update --init --recursive
$ go install ./...
```

# UNIT TEST

```console
$ go test
```

# INTEGRATION TEST

```console
$ mage integrationTest
```

# UNIT + INTEGRATION TEST

```console
$ mage test
```

# LINT

```console
$ mage lint
```

# PORT

```console
$ mage port
```

# CLEAN ALL ARTIFACTS

```console
$ mage clean; mage uninstall; mage -clean
```
