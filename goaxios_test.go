package goaxios

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestGetMethod(t *testing.T) {

	t.Run("ContentType - text/html", func(t *testing.T) {
		type Todo struct {
			UserId    int    `json:"userId"`
			Id        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		a := GoAxios{
			Url: "https://jsonplaceholder.typicode.com/:path/",
			Params: map[string]string{
				"path": "todos",
			},
			Method:         "GET",
			ResponseStruct: &[]Todo{},
		}

		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
		}

		v := res.rawRes.(*[]Todo)
		if (*v)[0].Title == "" {
			t.Errorf("expected: %v, got: %v", "delectus aut autem", (*v)[0].Title)
		}
	})

	t.Run("ContentType - text/plain", func(t *testing.T) {
		a := GoAxios{
			Url:    "https://type.fit/api/quotes",
			Method: "GET",
			ResponseStruct: &[]struct {
				Text   string `json:"text"`
				Author string `json:"author"`
			}{},
		}
		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
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
		type Travelers struct {
			XMLName             xml.Name              `xml:"travelers"`
			TravelerInformation []TravelerInformation `xml:"Travelerinformation"`
		}
		a := GoAxios{
			Url:    "http://restapi.adequateshop.com/api/Traveler?page=1",
			Method: "GET",
			ResponseStruct: &[]struct {
				XMLName     xml.Name    `xml:"TravelerinformationResponse"`
				Page        string      `xml:"page"`
				PerPage     string      `xml:"per_page"`
				TotalRecord string      `xml:"totalrecord"`
				TotalPages  string      `xml:"total_pages"`
				Travelers   []Travelers `xml:"travelers"`
			}{},
		}

		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
		}
	})
}

func TestPostMethod(t *testing.T) {

	t.Run("ContentType - application/json", func(t *testing.T) {
		// build body
		a := GoAxios{
			Url:    "https://reqres.in/api/users",
			Method: "POST",
			Body: map[string]string{
				"name": "morpheus",
				"job":  "leader",
			},
		}
		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
		}
	})

	t.Run("ContentType - multipart/form-data", func(t *testing.T) {
		a := GoAxios{
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

		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
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

		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
		}
		fmt.Println(res.rawRes)
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

		res := a.RunRest()
		if res.err != nil {
			t.Errorf("err: %v", res.err)
		}
	})
}
