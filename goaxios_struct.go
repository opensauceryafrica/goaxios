package goaxios

import (
	"net/http"
	"time"
)

type GoAxios struct {
	Url            string
	Method         string
	Body           interface{}
	Form           interface{}
	Query          map[string]interface{}
	BearerToken    string
	ResponseStruct interface{}
	Headers        map[string]string
	IsMultiPart    bool // if true, then the body is a multipart form
	Timeout        time.Duration
	Methods
}

type Methods interface {
	ValidateBeforeRequest() error

	RunRest() (*http.Response, []byte, interface{}, error)

	RunGraphQL() (*http.Response, []byte, interface{}, error)

	PerformResponseMarshalling(string, interface{}, []byte, []byte, error, *http.Response) (*http.Response, []byte, interface{}, error)
}
