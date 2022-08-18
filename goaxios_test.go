package goaxios

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
	"time"
)

func TestGetMethod(t *testing.T) {

	t.Run("ContentType - application/json", func(t *testing.T) {
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

	t.Run("ContentType - text/html", func(t *testing.T) {
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
		h := map[string]interface{}{
			"cid":      "bafkreia4ruswe7ghckleh3lmpujo5asrnd7hrtu5r23zjk2robpcoend34",
			"duration": 190,
			"verified": true,
			"price":    0,
			"client":   "",
			"token": map[string]interface{}{
				"token": "testing",
			},
		}

		fmt.Println(h)

		a := GoAxios{
			Url:    "http://35.202.1.73/deal/make",
			Method: "POST",
			Body:   h,
		}
		_, b, d, _ := a.RunRest()
		fmt.Printf("%+v\n", d)

		var n struct {
			Status bool `json:"status"`
			Data   struct {
				ID             string      `json:"id" dynamobav:"id"`
				DealID         interface{} `json:"deal_id" dynamobav:"deal_id"`
				IsDeal         bool        `json:"is_deal" default:"true" dynamobav:"is_deal"`
				CID            string      `json:"cid" dynamobav:"cid"`
				ImportCID      string      `json:"import_cid" dynamobav:"import_cid"`
				DealCID        string      `json:"deal_cid" dynamobav:"deal_cid"`
				Price          int64       `json:"price" dynamobav:"price"`
				Duration       int         `json:"duration" dynamobav:"duration"`
				Created        time.Time   `json:"created" dynamobav:"created"`
				Updated        time.Time   `json:"updated" dynamobav:"updated"`
				Renewed        string      `json:"renewed" dynamobav:"renewed"`
				Expired        string      `json:"expired" dynamobav:"expired"`
				Status         string      `json:"status" dynamobav:"status"`
				Miner          string      `json:"miner" dynamobav:"miner"`
				Verified       bool        `json:"verified" dynamobav:"verified"`
				Address        string      `json:"address" dynamobav:"address"`
				RetrievalError []string    `json:"retrieval_error" dynamobav:"retrieval_error"`
				Client         string      `json:"client" dynamobav:"client" default:"lotus"`
			} `json:"data"`
		}

		_ = json.Unmarshal(b, &n)
		fmt.Printf("%+v\n", n)
	})
}
