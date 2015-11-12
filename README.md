# wildp
wildp handles wildcard easily.

## Usage
```go
fmt.Printf("%v", wildp.Args)
```

```shell
$ dir /B
README.md
wildp.go
wildp_test.go
```

```shell
$ example.exe a b c wild?*
[example.exe a b c wildp.go wildp_test.go]
```

```shell
$ example.exe a b c w[a-z]ld?_*
[example.exe a b c wildp_test.go]
```

```shell
$ example.exe a b c [!w]ildp.go
[example.exe a b c]
```

## Installation
```
$ go get github.com/donke/wildp
```

## License
MIT

## Author
Ken Kudo (aka.kudoken@gmail.com)
