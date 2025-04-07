# Go MAIB ECommerce SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/NikSays/go-maib-ecomm.svg)](https://pkg.go.dev/github.com/NikSays/go-maib-ecomm)

    Copyright (C) 2025 Nejintev Nicolai

    This library is free software: you can redistribute it and/or modify
    it under the terms of the GNU Lesser General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This library is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Lesser General Public License for more details.

    You should have received a copy of the GNU Lesser General Public License
    along with this library.  If not, see <https://www.gnu.org/licenses/>.

This library provides tools to interact with MAIB ECommerce in a type safe way.

This library is **UNOFFICIAL**, inspired by public Integration Guide and [implementations](https://github.com/maibank) 
shared by MAIB.

Unlike other implementations, you don't need to run any `openssl` commands to change the certificate's format.
Only the default `.pfx` pkcs#12 file is needed. 

## Usage

To install, run the following:
```shell
go get github.com/NikSays/go-maib-ecomm/v2
```

Then use the package `maib`:
```go
package main
import "github.com/NikSays/go-maib-ecomm/v2"

func main()  {
    conf := maib.Config{/* ... */}
    // ...
}
```

## Documentation

Documentation and examples are available at [Go Reference](https://pkg.go.dev/github.com/NikSays/go-maib-ecomm).

## Testing

If you need to regenerate the testing certificates, run the `./genCrt.sh` script
in the ./testdata folder.

To test the package just run `go test ./... `.
