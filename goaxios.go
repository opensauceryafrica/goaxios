package goaxios

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GoAxios struct {
	Url            string
	Method         string
	Body           string
	Query          map[string]interface{}
	Token          string
	ResponseStruct interface{}
	Headers        map[string]string
}

// a wrapper around Go's *http.Request ojbect to make it faster to run REST http requests.
// It returns the *http.Response object and the response body as a map[string]interface{} and error (if any or nil)
func (ga *GoAxios) RunRest() (*http.Response, interface{}, error) {

	url := ga.Url + "?"

	// parse query params
	for k, v := range ga.Query {
		url = url + k + "=" + v.(string) + "&"
	}

	// parse body
	reqBody := strings.NewReader(ga.Body)

	client := &http.Client{}

	// fake http response
	var fail *http.Response

	// response body
	var response interface{}
	if ga.ResponseStruct != nil {
		response = ga.ResponseStruct
	}

	req, err := http.NewRequest(ga.Method, url, reqBody)
	if err != nil {
		return fail, response, err
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
	if ga.Token != "" {
		req.Header.Add("Authorization", "Bearer "+ga.Token)
	}

	res, err := client.Do(req)
	if err != nil {
		return res, response, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return res, response, err
	}

	// unmarshall
	contentType := res.Header.Get("Content-Type")
	fmt.Println("CONTENT TYPE: ", contentType)
	if strings.Contains(contentType, "application/json") {
		err = json.Unmarshal(data, &response)
		if err != nil {
			fmt.Println(err)
			return res, response, err
		}
	} else {
		switch contentType {
		case "text/html":
			response = string(data)
		case "text/plain":
			if ga.ResponseStruct != nil {
				err = json.Unmarshal(data, &response)
				if err != nil {
					fmt.Println(err)
					return res, response, err
				}
			} else {
				response = string(data)
			}
		case "application/xml":
			if ga.ResponseStruct != nil {
				err = xml.NewDecoder(res.Body).Decode(response)
			} else {
				response = string(data)
			}
		}

	}

	return res, response, err
}

// a wrapper around Go's *http.Request object to make it faster to run GraphQL http requests.
// It returns the *http.Response object and the response body as a map[string]interface{} and error (if any or nil)
func (ga *GoAxios) RunGraphQL() (*http.Response, map[string]interface{}, error) {

	return new(http.Response), *new(map[string]interface{}), nil
}
