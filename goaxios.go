package goaxios

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// RunRest is a tiny wrapper around Go's *http.Request object to make it quicker to run REST http requests.
// It returns the *http.Response object, the response body as byte, the unmarshalled response body and an error object (if any or nil)
func (ga *GoAxios) RunRest() (Response) {
  
	if ga.Interceptor.Request != nil {
		ga.Interceptor.Request(ga)
	}

	// TODO: improve validate before request
	err := ga.validateBeforeRequest()
	if err != nil {
		return Response {
			Response: nil,
			Bytes: nil,
			Body: nil,
			Error: err,
		}
	}

	// replace path params
	for k, v := range ga.Params {
		ga.Url = strings.Replace(ga.Url, ":"+k, v, -1)
	}

	// parse query params
	for k, v := range ga.Query {
		if strings.HasSuffix(ga.Url, "?") {
			ga.Url += k + "=" + v + "&"
		} else if strings.HasSuffix(ga.Url, "&") {
			ga.Url += k + "=" + v + "&"
		} else {
			ga.Url += "?" + k + "=" + v + "&"
		}
	}
	ga.Url = strings.TrimSuffix(ga.Url, "&")

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
		return Response{
			Response: fail,
			Bytes: body,
			Body: response,
			Error: err,
		}
	}

	client := &http.Client{
		Timeout: ga.Timeout,
	}

	req, err := http.NewRequest(strings.ToUpper(ga.Method), ga.Url, nil)
	if err != nil {
		return Response{
			Response: fail,
			Bytes: body,
			Body: response,
			Error: err,
		}
	}

	// add headers
	for k, v := range ga.Headers {
		req.Header.Add(k, v)
	}

	// add body
	if ga.IsMultiPart || ga.Form != nil {
		r, w := io.Pipe()
		writer := multipart.NewWriter(w)

		go func() {
			defer w.Close()
			defer writer.Close()

			for _, pf := range ga.Form.Files {
				var file io.ReadCloser
				var err error
				// open file
				if pf.Path != "" && pf.Handle == nil {
					file, err = os.Open(pf.Path)
					if err != nil {
						return
					}
				} else {
					file = pf.Handle
				}
				// close file
				defer file.Close()

				part, err := writer.CreateFormFile(pf.Key, pf.Name)
				if err != nil {
					return
				}
				_, err = io.Copy(part, file)
				if err != nil {
					return
				}
			}

			for _, pd := range ga.Form.Data {
				_ = writer.WriteField(pd.Key, pd.Value)
			}
		}()

		req.Body = r
		req.Header.Add("Content-Type", writer.FormDataContentType())
	} else {
		if ga.Body != nil {
			closerBody := ioutil.NopCloser(bytes.NewReader(reqBody))
			req.Body = closerBody
		}
		if ga.Headers == nil {
			req.Header.Add("Content-Type", "application/json")
		}
	}

	// add bearer token
	if ga.BearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+ga.BearerToken)
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{
			Response: res,
			Bytes: body,
			Body: response,
			Error: err,
		}
	}

	if ga.Interceptor.Response != nil {
		ga.Interceptor.Response(res)
	}

	defer res.Body.Close()

	// handle download
	if ga.IsDownload {
		if ga.DownloadDestination.Location != "" {
			out, err := os.Create(ga.DownloadDestination.Location)
			if err != nil {
				return Response {
					Response: res,
					Bytes: body,
					Body: response,
					Error: err,
				}
			}
			defer out.Close()
			_, err = io.Copy(out, res.Body)
			if err != nil {
				return Response {
					Response: res,
					Bytes: body,
					Body: response,
					Error: err,
				}
			}
		} else if ga.DownloadDestination.Writer != nil {
			_, err = io.Copy(ga.DownloadDestination.Writer, res.Body)
			if err != nil {
				return Response {
					Response: res,
					Bytes: body,
					Body: response,
					Error: err,
				}
			}
		} else {
			return Response {
				Response: res,
				Bytes: body,
				Body: response,
				Error: errors.New("download destination not provided"),
			}
		}

		return Response {
			Response: res,
			Bytes: body,
			Body: response,
			Error: nil,
		}
	} else {

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return Response {
				Response: res,
				Bytes: body,
				Body: response,
				Error: err,
			}
		}

		// unmarshall
		contentType := res.Header.Get("Content-Type")

		return ga.performResponseMarshalling(contentType, response, data, body, err, res)
	}
}

// RunGraphQL is a wrapper around Go's *http.Request object to make it faster to run GraphQL http requests.
// It returns the *http.Response object, the response body as byte, the unmarshalled response body and an error object (if any or nil)
func (ga *GoAxios) RunGraphQL() (*http.Response, []byte, interface{}, error) {

	return new(http.Response), *new([]uint8), new(interface{}), nil
}

// URL is a getter for the URL property
func (ga *GoAxios) URL() string {
	return ga.Url
}
