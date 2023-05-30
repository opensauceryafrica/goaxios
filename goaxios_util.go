package goaxios

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (ga *GoAxios) validateBeforeRequest() error {

	if ga.Url == "" {
		return errors.New("url is required")
	}

	if ga.Method == "" {
		return errors.New("method is required")
	}

	if ga.Body != nil || ga.Form != nil {
		if ga.Method == "GET" {
			return errors.New("body is not allowed for GET request")
		}

		if ga.Method == "DELETE" {
			log.Default().Println("body may not be allowed for DELETE requests")
		}
	}

	return nil
}

// marshalls the response body based on the content type and user-defined struct, if any.
func (ga *GoAxios) performResponseMarshalling(contentType string, response interface{}, data, body []byte, err error, res *http.Response) (*http.Response, []byte, interface{}, error) {
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
			err = xml.Unmarshal(data, &response)
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
