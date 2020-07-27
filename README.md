# JSON Errors for Go

[![GoDoc](https://godoc.org/github.com/erdaltsksn/jerr?status.svg)](https://godoc.org/github.com/erdaltsksn/jerr)
![Go](https://github.com/erdaltsksn/jerr/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/jerr)](https://goreportcard.com/report/github.com/erdaltsksn/jerr)

`jerr` provides a simple error handling for your rest applications.

## Features

- Compatible with built-in `error` interface
- Error wrapping
- JSON escaping

## Getting Started

```sh
go get github.com/erdaltsksn/jerr
touch main.go
```

**main.go:**

```go
package main

import (
	"fmt"

	"github.com/erdaltsksn/jerr"
)

func main() {
	err := someFunc()
	if err != nil {
		wrapped := jerr.Wrap(err, "Message about error")
		fmt.Println(wrapped.Error())
	}
}

func someFunc() error {
	return jerr.New("nope")
}
```

```sh
go run main.go
```

**Output:**

```json
{"message":"Message about error","details":{"message":"nope"}}
```

Check out [examples](examples/simple) directory for more.

## Installation

```sh
go get github.com/erdaltsksn/jerr
```

## Updating / Upgrading

```sh
go get -u github.com/erdaltsksn/jerr
```

## Usage

```go
// Simple error.
errSimple := jerr.New("Error Message")

// An error that wraps another error.
err := SomeFunc()
if err != nil {
    fmt.Print(jerr.Wrap(err, "Message about error"))
}
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](docs/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](docs/SECURITY.md).

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](docs/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
