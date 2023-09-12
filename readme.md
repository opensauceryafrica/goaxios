# goaxios - inspired by the popular JavaScript axios

A lightweight package that makes it easier to send Rest and GraphQL requests in Golang.
For every request you make, Goaxios returns the http response object, the raw response body in bytes, the parsed response body in a struct or interface, and an error object.

## Features

- [x] Create and run REST HTTP requests
- [x] Basic configuration for REST HTTP requests
- [x] Validate Goaxios Struct before running request
- [x] Interceptors for before Request and after Response
- [x] Multipart form data requests
- [ ] Create and run GraphQL HTTP requests
- [ ] Basic configuration for GraphQL HTTP requests
- [x] Download file to a location
- [x] Download file to a writer
- [x] Upload file from a source
- [ ] Upload and download progress
- [ ] JavaScript `Promise.all()` style to run multiple requests
- [x] Auto build path from path parameters like `/users/:id`

## Installation

```bash
go get github.com/opensaucerer/goaxios
```

## Usage for REST HTTP requests

### Run a simple GET request

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)
func main() {
    a := goaxios.GoAxios{
        Url:    "https://anapioficeandfire.com/api/houses/1",
        Method: "GET",
    }
    res := a.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }
    fmt.Printf("Response Object: ", res.Response)
    fmt.Printf("Raw Body in Bytes: ", res.Bytes)
    fmt.Printf("Parsed Body: ", res.Body)
}
```

### POST request with query and path parameters

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

func main() {
    a := goaxios.GoAxios{
        Url:    "https://anapioficeandfire.com/api/:id",
        Params: map[string]string{
            "id": "houses",
        },
        Method: "POST",
        Query: map[string]string{
            "name": "House Stark",
            "region": "The North",
        },
    }
    res := a.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }
}
```

### POST request with body

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

func main() {
    a := goaxios.GoAxios{
        Url:    "https://anapioficeandfire.com/api/houses",
        Method: "POST",
        Body: map[string]string{
            "name": "House Stark",
            "region": "The North",
        },
    }
    res := a.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }
}
```

### GET request with custom response struct

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

type House struct {
    Name    string `json:"name"`
    Region  string `json:"region"`
    Words   string `json:"words"`
    Seats   []string `json:"seats"`
}

func main() {
    a := goaxios.GoAxios{
        Url:    "https://anapioficeandfire.com/api/houses/1",
        Method: "GET",
        ResponseStruct: &House{},
    }
    res := a.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }
    houseData,_  := res.Body.(*House)
    
    fmt.Println(houseData.Name)
}
```

### Request With bearer token

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

type ResponseStruct struct {
    Login            string `json:"login"`
    Id               int    `json:"id"`
    NodeID           string `json:"node_id"`
    URL              string `json:"url"`
    ReposURL         string `json:"repos_url"`
    EventsURL        string `json:"events_url"`
    HooksURL         string `json:"hooks_url"`
    IssuesURL        string `json:"issues_url"`
    MembersURL       string `json:"members_url"`
    PublicMembersURL string `json:"public_members_url"`
    AvatarURL        string `json:"avatar_url"`
    Description      string `json:"description"`
}

func main() {
    token := ""
    a := goaxios.GoAxios{
        Url:            "https://api.github.com/user/orgs",
        Method:         "GET",
        ResponseStruct: &[]ResponseStruct{},
        BearerToken:    token,
    }
    res := a.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }

    responseData,_  := res.Body.(*[]ResponseStruct)
    fmt.Println(responseData.MembersURL)
}
```

### URL encoded POST request

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

func main() {
    r := goaxios.GoAxios{
        Url:     "https://api.twitter.com/2/oauth2/token",
        Method:  "POST",
        Headers: map[string]string{
            // needs to be empty to prevent goaxios from setting content-type to application/json
        },
        Query: map[string]interface{}{
            "grant_type":    "refresh_token",
            "refresh_token": refreshToken,
            "client_id":     config.Env.TwitterClientID,
        },
    }
    res := r.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }
    fmt.Println(res.Body)
}
```

### Multipart form data POST request

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

func main() {
    r := goaxios.GoAxios{
        Url:    "https://api.pinata.cloud/pinning/pinFileToIPFS",
        Method: "POST",
        Form: &Form{
            Files: []FormFile{
                {
                    Name: "somefile.json",
                    Path: os.Getenv("LOCATION"),
                    Key:  "file",
                },
            },
        },
        BearerToken: os.Getenv("TOKEN"),
    }
    res := r.RunRest()
    if res.Error != nil {
        fmt.Printf("err: %v", err)
    }
    fmt.Println(res.Body)
}
```

### Multipart form data POST request with an in-memory file

```go
package main

import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

func main() {
    f, _ := os.Open(os.Getenv("LOCATION"))

    r := GoAxios{
        Url:    "https://api.pinata.cloud/pinning/pinFileToIPFS",
        Method: "POST",
        Form: &Form{
            Files: []FormFile{
                {
                    Name: "somefile.json",
                    Handle: f,
                    Key:  "file",
                },
            },
        },
        BearerToken: os.Getenv("TOKEN"),
    }
    res := r.RunRest()
    if res.Error != nil {
        log.Fatalf("err: %v", res.Error)
    }
    fmt.Println(res.Body)
}
```

### Download file to a location

```go
package main

import (
	"log"

	"github.com/opensaucerer/goaxios"
)

func main() {

	a := goaxios.GoAxios{
		Url:        "https://media.publit.io/file/wm_22a67238/castorr.webm",
		Method:     "GET",
		IsDownload: true,
		DownloadDestination: goaxios.Destination{
			Location: "test.webm",
		},
	}

	res := a.RunRest()
	if res.Error != nil {
		log.Fatalf("err: %v", res.Error)
	}
}
```

### Download file to a writer

```go
package main

import (
	"log"
	"os"

	"github.com/opensaucerer/goaxios"
)

func main() {

	// create a file
	w, err := os.Create("castorr.webm")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer w.Close()

	a := goaxios.GoAxios{
		Url:        "https://media.publit.io/file/wm_22a67238/castorr.webm",
		Method:     "GET",
		IsDownload: true,
		DownloadDestination: goaxios.Destination{
			Writer: w,
		},
	}

	res = a.RunRest()
	if res.Error != nil {
		log.Fatalf("err: %v", res.Error)
	}
}
```

## Interceptors

### Request

```go
package main
import (
    "fmt"
    "github.com/opensaucerer/goaxios"
)

a := goaxios.GoAxios{
  Url:    "https://api.twitter.com/2/oauth2/token",
  Method: "POST",
  Interceptor: goaxios.Interceptor{
    Request: func(req *goaxios.GoAxios) *goaxios.GoAxios {
      // modify the request as needed
      req.BearerToken = "token"
      req.Headers = map[string]string{
        "Content-Type": "application/json",
      }
      req.Body = map[string]string{
        "key": "value",
      }
      return req
    },
    Response: func(resp *http.Response) *http.Response {
      // do something with the response - logging/error handling or something
      if resp.StatusCode != 200 {
        panic("not OK")
      }
      return resp
    },
  },
}

res := a.RunRest()
if res.Error != nil {
    t.Errorf("err: %v", res.Error)
}
fmt.Println(res.Body)
```

## Usage for GraphQL HTTP requests

## Contributing

Contributions are welcome. Please open a [pull request](https://github.com/opensaucerer/goaxios/pulls) or open an [issue](https://github.com/opensaucerer/goaxios/issues) to discuss the change you wish to make.
