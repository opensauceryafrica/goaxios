package goaxios

import "net/http"

// type GoAxios struct {
// 	Url    string
// 	Method string
// 	Body   string
// 	Query  map[string]interface{}
// 	Token  string
// }

type Axios interface {
	ValidateBeforeRequest() error

	RunRest() (*http.Response, []byte, interface{}, error)

	RunGraphQL() (*http.Response, []byte, interface{}, error)

	PerformResponseMarshalling(string, interface{}, []byte, []byte, error, *http.Response) (*http.Response, []byte, interface{}, error)
}
