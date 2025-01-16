# Go vs TinyGo

Versions:

```bash
$ go version
go version go1.23.3 linux/amd64

$ tinygo version
tinygo version 0.35.0 linux/amd64 (using go version go1.23.3 and LLVM version 18.1.2)
```

## Go package

```bash
$ go build -o hello-go main.go ; ll -h hello-go

-rwxr-xr-x 1 gitpod gitpod 1.5M Jan 16 17:03 hello-go*
```

## TinyGo package

```bash
$ tinygo build -o hello-tinygo main.go ; ll -h hello-tinygo

-rwxr-xr-x 1 gitpod gitpod 114K Jan 16 17:03 hello-tinygo*
```
