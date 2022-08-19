package goaxios

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
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
}

// a wrapper around Go's *http.Request ojbect to make it faster to run REST http requests.
// It returns the *http.Response object, the response body as byte, the unmarshalled response body and an error object (if any or nil)
func (ga *GoAxios) RunRest() (*http.Response, []byte, interface{}, error) {

	// TODO: validate before request

	url := ga.Url + "?"

	// parse query params
	l := len(ga.Query)
	i := 0
	for k, v := range ga.Query {
		if i == 0 && l > 1 {
			url = url + k + "=" + v.(string) + "&"
		} else if i == l-1 {
			url = url + k + "=" + v.(string)
		} else {
			url = url + k + "=" + v.(string) + "&"
		}
		i++
	}

	// fake http response
	var fail *http.Response
	// fake response body
	var body []byte

	// response body
	var response interface{}
	if ga.ResponseStruct != nil {
		response = ga.ResponseStruct
	}

	// parse body
	// reqBody := strings.NewReader(ga.Body)
	reqBody, err := json.Marshal(ga.Body)
	if err != nil {
		return fail, body, response, err
	}

	client := &http.Client{
		Timeout: ga.Timeout,
	}

	req, err := http.NewRequest(strings.ToUpper(ga.Method), url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fail, body, response, err
	}

	// add headers
	if ga.Headers != nil {
		for k, v := range ga.Headers {
			req.Header.Add(k, v)
		}
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	// add bearer token
	if ga.BearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+ga.BearerToken)
	}

	res, err := client.Do(req)
	if err != nil {
		return res, body, response, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res, body, response, err
	}

	// unmarshall
	contentType := res.Header.Get("Content-Type")

	return ga.PerformResponseMarshalling(contentType, response, data, body, err, res)
}

// a wrapper around Go's *http.Request object to make it faster to run GraphQL http requests.
// It returns the *http.Response object, the response body as byte, the unmarshalled response body and an error object (if any or nil)
func (ga *GoAxios) RunGraphQL() (*http.Response, []byte, interface{}, error) {

	return new(http.Response), *new([]uint8), new(interface{}), nil
}

func (ga *GoAxios) PerformResponseMarshalling(contentType string, response interface{}, data, body []byte, err error, res *http.Response) (*http.Response, []byte, interface{}, error) {
	switch true {
	case strings.Contains(contentType, "text/plain"):
		if ga.ResponseStruct != nil {
			err = json.Unmarshal(data, &response)
			if err != nil {
				return res, body, response, err
			}
		} else {
			response = string(data)
		}
	case strings.Contains(contentType, "application/xml"):
		if ga.ResponseStruct != nil {
			err = xml.NewDecoder(res.Body).Decode(response)
			if err != nil {
				return res, body, response, err
			}
		} else {
			response = string(data)
		}
	default:
		err = json.Unmarshal(data, &response)
		if err != nil {
			if ga.ResponseStruct != nil {
				return res, body, response, err
			} else {
				err = nil
				response = string(data)
			}
		}
	}
	return res, data, response, err
}
