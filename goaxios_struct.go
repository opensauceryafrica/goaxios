package goaxios

import (
	"io"
	"time"
)

type GoAxios struct {
	// Url is the url to send the request to.
	Url string
	// Method is the http method to use for the request. This can be GET, POST, PUT, PATCH, DELETE, etc.
	Method string
	// Body to pass to the request. This can be a string, []byte, or a struct
	Body interface{}
	// Form is only required when dealing with multi-part/formdata
	Form *Form
	/*Query represents the parameters to add to the url.
	Example:
	`/users?name=John&age=30` */
	Query map[string]string
	/*Params is the path parameters to replace in the url.
	Example:
	`/users/:name/:age` */
	Params map[string]string
	// BearerToken is the bearer token to use for the request. This will be added to the Authorization header in the form `Bearer <token>`
	BearerToken string
	// ResponseStruct is the struct to use for marshalling the response body. This is optional.
	ResponseStruct interface{}
	/*Headers is a map of headers to add to the request. A default
	`Content-Type: application/json`*/
	// is added if no headers are passed. To prevent this, pass an empty map.
	Headers map[string]string
	// Timeout is the timeout to use for the request. This is optional.
	Timeout time.Duration
	// IsMultiPart is a flag to indicate if the request is a multipart form.
	IsMultiPart bool
	// IsDownload is a flag to indicate if the request is a download request.
	// if true, a download destination must be provided.
	IsDownload bool
	// DownloadDestination is the path/writer to write the download to.
	// This is required if IsDownload is true.
	DownloadDestination Destination
}

type Destination struct {
	Location string
	Writer   io.Writer
}

// Form is the struct used to pass parameters to request methods.
type Form struct {
	// Files is a list of files to upload.
	Files []FormFile
	// Data is a list of data to upload along with the files.
	Data []FormData
}

// FormData is the struct for uploading data along with files in a multipart request.
type FormData struct {
	// Key is the key to use for the data.
	Key string
	// Value is the value to use for the data.
	Value string
}

// FormFile is the struct for uploading a single file in a multipart request.
type FormFile struct {
	// Name is the name of the file.
	Name string
	// Path is the path to the file. This has an xor relationship with Handle.
	Path string
	// Key is the key to use for the file.
	Key string
	// Handle is an io.ReadCloser to use instead of a file path. This has an xor relationship with Path. GoAxios will close this handle after reading it. If you handle does not implement io.Closer, then you can use io.NopCloser to wrap it.
	Handle io.ReadCloser
}
