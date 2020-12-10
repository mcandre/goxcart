# goxcart: Go port all the things!

![goxcart-logo](https://raw.githubusercontent.com/mcandre/goxcart/master/goxcart.png)

# EXAMPLE

```console
$ cd test
$ goxcart -output bin -banner "hello-0.0.1" -commands ./cmd/...

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

See `goxcart -help` for more details.

# ABOUT

goxcart helps Go developers build programs for more platforms. While Go and gox already provide a lot of support for basic cross-platform builds, goxcart takes this further, enabling more target tuples than are currently available for host-run builds. For example, goxcart can build Go binaries targeting:

* dragonfly
* linux-musl
* nacl
* plan9
* solaris

goxcart does this by building applications inside reusable Docker containers configured for these additional targets.

Note that some hosts may require manual `chown`, `chmod` corrections to build artifacts. This is a known issue with [Docker](https://github.com/moby/moby/issues/39441) on GNU/Linux hosts.

Note that any C / cgo dependencies are likely to present challenges to cross-platform builds. In this case, we recommend creating dedicated build VM's, one per build target a la [tonixxx](https://github.com/mcandre/tonixxx).

# DOWNLOAD

https://github.com/mcandre/goxcart/releases

# RUNTIME REQUIREMENTS

* [Docker](https://www.docker.com/)

## Recommended

* [tree](https://linux.die.net/man/1/tree)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
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

# SEE ALSO

* [xgo](https://github.com/karalabe/xgo) automates cross-compiling Go applications, including cgo apps with native dependencies.
* [mcandre/docker-gox](https://github.com/mcandre/docker-gox) provides base images featuring the gox porting utility.
* [CloudABI](https://nuxi.nl/) provides a C-level abstraction for applications to run on many different platforms.
* [tonixxx](https://github.com/mcandre/tonixxx) can port projects written in many more programming languages
