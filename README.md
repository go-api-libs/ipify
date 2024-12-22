# 🌐 ipify API
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/ipify.svg)](https://pkg.go.dev/github.com/go-api-libs/ipify/pkg/ipify)
[![Official Documentation](https://img.shields.io/badge/docs-API-blue)](https://www.ipify.org/)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/ipify)](https://goreportcard.com/report/github.com/go-api-libs/ipify)
![Code Coverage](https://img.shields.io/badge/coverage-26%25-red)
![API Health](https://img.shields.io/badge/API_health-95%25-brightgreen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

A simple public IP address API, easy to integrate into any application in seconds. ([Source](https://freepublicapis.com/ipify-api))

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/ipify/pkg/ipify
```

## Usage

### Example: 

```go
package main

import (
	"context"

	"github.com/go-api-libs/ipify/pkg/ipify"
)

func main() {
	c, err := ipify.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	getOkJSONResponse, err := c.Get(ctx, &ipify.GetParams{Format: "json"})
	if err != nil {
		panic(err)
	}

	// Use getOkJSONResponse object
}

```

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/go-api-libs/ipify/pkg/ipify): The Go reference documentation for the client package.
- [**Official Documentation**](https://www.ipify.org/): The official API documentation.
- [**OpenAPI Specification**](./api/openapi.json): The OpenAPI 3.1.0 specification.
- [**Go Report Card**](https://goreportcard.com/report/github.com/go-api-libs/ipify): Check the code quality report.

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/go-api-libs/ipify).

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
