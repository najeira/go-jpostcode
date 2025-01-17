# go-jpostcode

[![Go Reference](https://pkg.go.dev/badge/github.com/syumai/go-jpostcode.svg)](https://pkg.go.dev/github.com/syumai/go-jpostcode)

- go-jpostcode is a Go package to find Japanese address data from Japanese postal code.
  - This package was created to provide data from https://github.com/kufu/jpostcode-data
- This package requires Go 1.16+.

## Usage

```go
// Find an address
address, err := jpostcode.Find("0010928")

// Search addresses (Some addresses have same postal code)
addresses, err := jpostcode.Search("1138654")

// Print address as a JSON
addressJSON, err := address.ToJSON()
if err != nil { // error handling }
fmt.Println(addressJSON)
```

### Example

- [HTTP server example](https://github.com/syumai/go-jpostcode/blob/master/example/server/main.go)

```console
$ go run example/server/main.go
$ curl http://localhost:8090/0010928
```

### Building

go-jpostcode doesn't use `cgo` directly, but its dependency `badger` depends on it. If you want to run build without cgo, please add `cgo_enabled=0` to your build command.

## Install a CLI tool to get address from postcode

- A CLI tool is given as [jpost](https://github.com/syumai/go-jpostcode/blob/master/cmd/jpost).

### Installation

```
go install github.com/syumai/go-jpostcode/cmd/jpost@latest
```

### Usage of jpost

- To get address, **just give postal code** as argument.

```
# Get address from postal code: 0010928.
$ jpost 0010928
{"postcode":"0010928","prefecture":"北海道",...
```

## Update jpostcode-data

```console
$ make update
```

## License

- MIT

## Author

- [syumai](https://github.com/syumai)

## Original data

- [jpostcode-data](https://github.com/kufu/jpostcode-data)
