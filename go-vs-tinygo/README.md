# Go vs TinyGo

Versions:

```bash
$ go version
go version go1.23.1 linux/amd64

$ tinygo version
tinygo version 0.33.0 linux/amd64 (using go version go1.23.1 and LLVM version 18.1.2)
```

## Go package

```bash
$ go build -o hello-go main.go ; ll -h hello-go

-rwxr-xr-x 1 aurelievache staff 1.5M Oct 15 13:36 hello-go*
```

## TinyGo package

```bash
$ tinygo build -o hello-tinygo main.go ; ll -h hello-tinygo

-rwxr-xr-x 1 aurelievache staff 108K Oct 15 13:37 hello-tinygo*
```
