# GoAxios

A lightweight package that makes it easier to send Rest and GraphQL requests in Golang.
For every request you make, Goaxios returns the http response object, the raw response body in bytes, the parse response body in a struct or interface, and an error object.

## Features

- [x] Create and run REST HTTP requests
- [x] Basic configuration for REST HTTP requests
- [x] Validate Goaxios Struct before running request
- [ ] Create and run GraphQL HTTP requests
- [ ] Basic configuration for GraphQL HTTP requests
- [ ] Download file to a destination
- [ ] Upload file from a source
- [ ] Upload and download progress
- [ ] JavaScript `Promise.all()` style to run multiple requests
- [ ] Auto build path from path parameters like `/users/{id}`

## Installation

```bash
go get github.com/samperfect/goaxios
```

## Usage for REST HTTP requests

Run a simple GET request

```go
package main

import (
    "fmt"
    "github.com/samperfect/goaxios"
)
a := GoAxios{
    Url:    "https://anapioficeandfire.com/api/houses/1",
    Method: "GET",
}
r, b, d, err := a.RunRest()
if err != nil {
    fmt.Printf("err: %v", err)
}
fmt.Printf("Response Object: ", r)
fmt.Printf("Raw Body in Bytes: ", b)
fmt.Printf("Parsed Body: ", d)
```

Post request with query parameters

```go
package main

import (
    "fmt"
    "github.com/samperfect/goaxios"
)

a := GoAxios{
    Url:    "https://anapioficeandfire.com/api/houses",
    Method: "POST",
    Query: map[string]string{
        "name": "House Stark",
        "region": "The North",
    },
}
_, _, _, err := a.RunRest()
if err != nil {
    fmt.Printf("err: %v", err)
}
```

Post request with body

```go
package main

import (
    "fmt"
    "github.com/samperfect/goaxios"
)

a := GoAxios{
    Url:    "https://anapioficeandfire.com/api/houses",
    Method: "POST",
    Body: map[string]string{
        "name": "House Stark",
        "region": "The North",
    },
}
_, _, _, err := a.RunRest()
if err != nil {
    fmt.Printf("err: %v", err)
}
```

Get request with custom response struct

```go
package main

import (
    "fmt"
    "github.com/samperfect/goaxios"
)

type House struct {
    Name    string `json:"name"`
    Region  string `json:"region"`
    Words   string `json:"words"`
    Seats   []string `json:"seats"`
}
a := GoAxios{
    Url:    "https://anapioficeandfire.com/api/houses/1",
    Method: "GET",
    ResponseStruct: House{},
}
_, _, _, err := a.RunRest()
if err != nil {
    fmt.Printf("err: %v", err)
}
```

## Usage for GraphQL HTTP requests

## Contributing

Contributions are welcome. Please open a [pull request](https://github.com/samperfect/goaxios/pulls) an [issue](https://github.com/samperfect/goaxios/issues) to discuss the change you wish to make.
