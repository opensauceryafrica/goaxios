package goaxios

import (
	"net/http"
	"time"
)

type Interceptor struct {
	Request  func(req *GoAxios) *GoAxios
	Response func(resp *http.Response) *http.Response
}

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
	Interceptor Interceptor
}

type Methods interface {
	ValidateBeforeRequest() error

	RunRest() (*http.Response, []byte, interface{}, error)

	RunGraphQL() (*http.Response, []byte, interface{}, error)

	PerformResponseMarshalling(string, interface{}, []byte, []byte, error, *http.Response) (*http.Response, []byte, interface{}, error)
}
