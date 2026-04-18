# Module 01 - Introduction & Setup

## What is Go?

Go (a.k.a. **Golang**) is a compiled, statically typed programming language created at Google in 2007 by **Robert Griesemer, Rob Pike, and Ken Thompson**. First public release: 2009.

### Why Go?
- **Simple & minimal syntax** - ~25 keywords
- **Fast compilation** - compiles to a single native binary
- **Built-in concurrency** - goroutines & channels
- **Strong standard library** - HTTP, JSON, crypto, testing out of the box
- **Garbage collected** - no manual memory management
- **Cross-platform** - build for Linux/macOS/Windows from one machine

### Who uses Go?
Docker, Kubernetes, Terraform, Prometheus, Grafana, Uber, Netflix, Cloudflare, Twitch.

---

## Installation

**Linux / macOS:**
```bash
# Download from https://go.dev/dl/ then:
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

# Add Go to PATH (append to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

**Windows:** run the `.msi` installer from go.dev/dl/.

**Verify installation:**
```bash
go version
# go version go1.22.0 linux/amd64
```

---

## The `go` Command - Complete Reference

Go ships with ONE tool - the `go` command - that does everything: compile, test, format, manage dependencies, install binaries. Run `go help` to see every subcommand.

### 1. `go version` - show the installed Go version

```bash
go version
```

### 2. `go env` - show / set Go environment variables

```bash
go env                      # show all
go env GOPATH GOROOT        # specific ones
go env -w GOPROXY=direct    # permanently set
```

Key variables:
| Variable | Meaning |
|----------|---------|
| `GOROOT` | Where Go itself is installed (`/usr/local/go`) |
| `GOPATH` | Your Go workspace (`~/go`) |
| `GOBIN`  | Where `go install` puts binaries |
| `GOOS` / `GOARCH` | Target OS / CPU - used for cross-compilation |
| `GOPROXY` | Module proxy (default `https://proxy.golang.org`) |

### 3. `go run` - compile and execute in one step

```bash
go run hello.go              # run a single file
go run .                     # run all files in current package
go run ./cmd/server          # run a specific package
```

Great for quick experiments. Produces no binary.

### 4. `go build` - compile to an executable

```bash
go build                     # build current package → produces ./<dir-name>
go build -o myapp            # choose output name
go build ./...               # build every package recursively
go build -v                  # verbose (prints packages compiled)
```

The binary is **fully self-contained** - no runtime required on the target machine.

**Cross-compilation** - build for any OS/CPU from your machine:
```bash
GOOS=linux   GOARCH=amd64 go build -o app-linux
GOOS=windows GOARCH=amd64 go build -o app.exe
GOOS=darwin  GOARCH=arm64 go build -o app-mac
```

### 5. `go install` - build AND install to `$GOBIN`

```bash
go install                                        # install current package
go install github.com/user/tool@latest           # install a remote CLI tool
```

Installed binaries live in `$GOBIN` (default `~/go/bin`) - make sure it's in your `PATH`.

### 6. `go fmt` - auto-format source code

Go has ONE official style. Never debate formatting again.

```bash
go fmt ./...     # format every file in the module
gofmt -d file.go # show what would change (diff mode)
gofmt -s -w .    # simplify code too (-s), write in place (-w)
```

Pro tip: install `goimports` - formats AND manages imports automatically.
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

### 7. `go vet` - static analysis (lint)

Finds suspicious code: shadowed variables, wrong Printf verbs, unreachable code, bad struct tags, etc.

```bash
go vet ./...
```

Run it **before every commit**. Most editors do it automatically.

### 8. `go test` - run tests and benchmarks

```bash
go test                    # tests in current package
go test ./...              # every package recursively
go test -v                 # verbose - list every test
go test -run TestAdd       # run tests whose name matches "TestAdd"
go test -cover             # show coverage %
go test -coverprofile=c.out && go tool cover -html=c.out  # HTML coverage report
go test -race              # enable data-race detector (use this!)
go test -bench=.           # run benchmarks
```

### 9. `go mod` - dependency / module management

Every Go project since 1.16 is a **module**. A module is a folder with a `go.mod` file.

```bash
go mod init github.com/you/myapp    # start a new module
go mod tidy                          # add missing deps, remove unused
go mod download                      # pre-fetch dependencies
go mod graph                         # show dep graph
go mod why <package>                 # explain why a dep is needed
go mod verify                        # check cached deps haven't been tampered with
go mod edit -replace old=new         # swap a dep with a local path
```

`go.mod` (auto-generated):
```
module github.com/you/myapp

go 1.22

require (
    github.com/google/uuid v1.6.0
)
```

`go.sum` - checksums of every dependency version for reproducible builds. **Commit both.**

### 10. `go get` - add / upgrade dependencies

```bash
go get github.com/google/uuid              # latest version
go get github.com/google/uuid@v1.6.0       # specific version
go get -u ./...                            # upgrade every dep
go get github.com/foo/bar@none             # remove a dep
```

### 11. `go doc` - view documentation offline

```bash
go doc fmt               # package overview
go doc fmt.Println       # specific symbol
go doc -all fmt          # everything in the package
```

Or browse locally: `godoc -http=:6060` then open `http://localhost:6060`.

### 12. `go work` - multi-module workspaces (Go 1.18+)

Work on multiple related modules at once without publishing them.

```bash
go work init ./moduleA ./moduleB
go work use ./moduleC
```

### 13. `go clean` - remove build artifacts

```bash
go clean               # remove object files from current package
go clean -cache        # wipe the build cache
go clean -modcache     # wipe the module cache (~/go/pkg/mod)
```

### 14. `go tool` - low-level helpers

```bash
go tool cover -html=c.out       # visualize test coverage
go tool pprof cpu.out           # performance profiling
go tool objdump myapp           # disassemble
go tool dist list               # list all supported GOOS/GOARCH combinations
```

### 15. `go generate` - run code-generation directives

```bash
go generate ./...
```

Executes `//go:generate <cmd>` comments embedded in source files.

### 16. `go help` - discover everything

```bash
go help                 # list all subcommands
go help build           # detailed help for one
go help modules         # conceptual guide
go help gopath          # legacy workspace info
```

---

## Typical Developer Workflow

```bash
# 1. Create a new project
mkdir myapp && cd myapp
go mod init github.com/you/myapp

# 2. Write main.go, then:
go fmt ./...     # format
go vet ./...     # lint
go test ./...    # test
go build         # compile
./myapp          # run

# 3. Add a dependency
go get github.com/google/uuid
go mod tidy
```

---

## Anatomy of a Go Program

```go
package main              // every executable starts with package main

import "fmt"              // import standard library packages

func main() {             // entry point
    fmt.Println("Hello")  // call a function from fmt
}
```

Every `.go` file has three sections:
1. **Package declaration** - required (first non-comment line)
2. **Imports** - optional
3. **Declarations** - functions, types, variables, constants

---

## Files in this module

- [01_hello_world.go](01_hello_world.go) - first program
- [02_program_structure.go](02_program_structure.go) - anatomy of a Go file
- [03_comments.go](03_comments.go) - comment styles
- [04_fmt_basics.go](04_fmt_basics.go) - Println, Printf, Sprintf, Scanln
- [05_go_tool_cheatsheet.md](05_go_tool_cheatsheet.md) - printable one-page cheat sheet
- [06_first_module_demo/](06_first_module_demo/) - walk a complete `go mod init → build → run` workflow

---

## Exercises

1. Install Go and run `go version`, `go env GOPATH`, `go env GOOS`.
2. Create a folder `hello/`, run `go mod init example.com/hello`, write a Hello-World program, then `go run .` and `go build`.
3. Cross-compile your program for Windows: `GOOS=windows GOARCH=amd64 go build -o hello.exe`.
4. Intentionally mis-format a file (extra spaces, wrong indent) - run `go fmt` and observe the fix.
5. Write a program that uses a Printf verb incorrectly (e.g., `%d` with a string). Run `go vet` and read the warning.
6. Add `github.com/google/uuid` as a dependency. Import it, print a UUID, then run `go mod tidy`.
