<div align="center">
  <br>
  <h1>GOAxios</h1>
  <strong>A lightweight package for sending send Rest and GraphQL requests in Golang</strong>
</div>


## Table of Contents

  - [Features](#features)
  - [Installing](#installing)
  - [Example](#example)
  - [Improvements](#improvements)


## Features
- Make HTTP request
- Make Graphql requests

## Installation
```sh
  go get github.com/samperfect/goaxios
```



## Example 
```go
  import github.com/samperfect/goaxios
```
Performing a `GET` request
```go
    client := goaxios.GoAxios{
      Url:    "https://jsonplaceholder.typicode.com/todos/",
      Method: "GET",
      ResponseStruct: []struct {
        UserId    string `json:"userId"`
        Id        string `json:"id"`
        Title     string `json:"title"`
        Completed bool   `json:"completed"`
      }{},
    }

	_, _, res, err := client.RunRest()
	if err != nil {
		log.Fatal(err)
	}
	// do something with res
	log.Println(res)
```

## Improvements

- [x] Basic configuration for REST requests
- [ ] JavaScript `Promise.all()` style to run multiple requests
- [ ] Download file to a destination
- [ ] Upload file from a source
