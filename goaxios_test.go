package goaxios

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"
)

func TestGetMethod(t *testing.T) {

	t.Run("ContentType - text/html", func(t *testing.T) {
		a := GoAxios{
			Url:    "https://jsonplaceholder.typicode.com/todos/",
			Method: "GET",
			ResponseStruct: []struct {
				UserId    string `json:"userId"`
				Id        string `json:"id"`
				Title     string `json:"title"`
				Completed bool   `json:"completed"`
			}{},
		}

		_, _, _, err := a.RunRest()
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})

	t.Run("ContentType - text/plain", func(t *testing.T) {
		a := GoAxios{
			Url:    "https://type.fit/api/quotes",
			Method: "GET",
			ResponseStruct: []struct {
				Text   string `json:"text"`
				Author string `json:"author"`
			}{},
		}
		_, _, _, err := a.RunRest()
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})

	t.Run("ContentType - application/xml", func(t *testing.T) {
		type TravelerInformation struct {
			XMLName   xml.Name `xml:"Travelerinformation"`
			Id        string   `xml:"id"`
			Name      string   `xml:"name"`
			Email     string   `xml:"email"`
			Adderes   string   `xml:"adderes"`
			Createdat string   `xml:"createdat"`
		}
		type Travellers struct {
			XMLName             xml.Name              `xml:"travelers"`
			TravelerInformation []TravelerInformation `xml:"Travelerinformation"`
		}
		a := GoAxios{
			Url:    "http://restapi.adequateshop.com/api/Traveler?page=1",
			Method: "GET",
			ResponseStruct: []struct {
				XMLName     xml.Name `xml:"TravelerinformationResponse"`
				Page        string   `xml:"page"`
				PerPage     string   `xml:"per_page"`
				TotalRecord string   `xml:"totalrecord"`
				TotalPages  string   `xml:"total_pages"`
				Travellers  []Travellers
			}{},
		}

		_, _, _, err := a.RunRest()
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}

func TestPostMethod(t *testing.T) {

	t.Run("ContentType - application/json", func(t *testing.T) {
		// build body
		// http://34.67.216.167/api/v0/dag/stat?arg=bafybeihhmzinrglpc6isvfcsqg2edduechay46fusfh7yu47yfv77zy7mu&
		a := GoAxios{
			Url:    "http://34.67.216.167/api/v0/dag/stat",
			Method: "POST",
			Query: map[string]interface{}{
				"arg": "bafybeihhmzinrglpc6isvfcsqg2edduechay46fusfh7yu47yfv77zy7mu",
			},
		}
		_, _, _, err := a.RunRest()
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}

func TestRequestInterceptor(t *testing.T) {
	t.Run("reqeuest interceptor", func(t *testing.T) {
		a := GoAxios{
			Url:    "http://localhost:3000/check-header",
			Method: "POST",
			Interceptor: Interceptor{
				Request: func(req *GoAxios) *GoAxios {
					// Modify the request as needed
					req.BearerToken = "token"
					req.Headers = map[string]string{
						"Content-Type": "application/json",
						"hello":        "world",
					}
					req.Body = map[string]string{
						"key": "value",
					}
					return req
				},
			},
		}

		_, _, response, err := a.RunRest()
		if err != nil {
			t.Errorf("err: %v", err)
		}
		fmt.Println(response)
	})
}

func TestResponseInterceptor(t *testing.T) {
	t.Run("response interceptor", func(t *testing.T) {
		a := GoAxios{
			Url:         "http://localhost:3000",
			Method:      "PATCH",
			BearerToken: "token",
			Body: map[string]string{
				"key": "value",
			},
			Interceptor: Interceptor{
				Response: func(resp *http.Response) *http.Response {
					if resp.StatusCode != 200 {
						panic("not OK")
					}
					return resp
				},
			},
		}

		_, _, _, err := a.RunRest()
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}
