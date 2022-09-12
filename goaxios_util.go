package goaxios

import (
	"errors"
	"log"
)

func (ga *GoAxios) ValidateBeforeRequest() error {

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
