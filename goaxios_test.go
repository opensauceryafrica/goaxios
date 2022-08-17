package goaxios

import (
	"fmt"
	"testing"
)

func TestGetMethod(t *testing.T) {
	a := GoAxios{
		Url:    "https://type.fit/api/quotes",
		Method: "GET",
		ResponseStruct: []struct {
			Text   string `json:"text"`
			Author string `json:"author"`
		}{},
	}
	_, d, _ := a.RunRest()
	fmt.Println(d)
}
