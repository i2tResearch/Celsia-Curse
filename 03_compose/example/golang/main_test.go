// main_test.go

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/gorilla/mux" // used to get the params from the route
	"gopkg.in/resty.v1"
)


func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.ServeHTTP(rr, req)


	return rr
}

func TestAddSong(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().
			SetBody([]byte(`{
				"name": "Dakiti",
				"singer":"Bad Bunny",
				"genre":"Reggaeton"
			}
			`)).
			Post("http://localhost:8080/api/newsong")

	fmt.Printf("%+v\n", resp)
}

func TestGetSongs(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://localhost:8080/api/songs")
	//fmt.Printf("%+v\n", resp)

	bodyBytes := resp.Body()


	var m []map[string]interface{}
    json.Unmarshal([]byte(bodyBytes), &m)


	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

    if m[0]["name"] != "Dakiti" {
        t.Errorf("Expected name  to be 'Dakiti'. Got '%v'", m[0]["name"])
    }

    if m[0]["singer"] != "Bad Bunny" {
        t.Errorf("Expected singer to be 'Bad Bunny'. Got '%v'", m[0]["singer"])
    }

	if m[0]["genre"] != "Reggaeton" {
        t.Errorf("Expected genre to be 'Reggaeton'. Got '%v'", m[0]["genre"])
    }

}

func TestEditSong(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().
			SetBody([]byte(`{
				"name": "Flaca",
				"singer":"Andres Calamaro",
				"genre":"Rock"
			}
			`)).
			Put("http://localhost:8080/api/song/1")



	cl := resty.New()
	response, _ := cl.R().Get("http://localhost:8080/api/songs")
	//fmt.Printf("%+v\n", response)
		
	bodyBytes := response.Body()
		
		
	var m []map[string]interface{}
	json.Unmarshal([]byte(bodyBytes), &m)

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

    if m[0]["name"] != "Flaca" {
        t.Errorf("Expected name to be 'Flaca'. Got '%v'", m[0]["name"])
    }

    if m[0]["singer"] != "Andres Calamaro" {
        t.Errorf("Expected singer to be 'Andres Calamaro'. Got '%v'", m[0]["singer"])
    }

	if m[0]["genre"] != "Rock" {
        t.Errorf("Expected genre to be 'Rock'. Got '%v'", m[0]["genre"])
    }
	
}

func TestDeleteSong(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().
				Delete("http://localhost:8080/api/deletesong/1")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}

	cl := resty.New()
	response, _ := cl.R().Get("http://localhost:8080/api/songs")
		
	bodyBytes := response.Body()
		
		
	var m []map[string]interface{}
	json.Unmarshal([]byte(bodyBytes), &m)
	//fmt.Printf("%+v\n", len(m))



	if len(m) > 0 {
		t.Errorf("List not empty")	
	}	
}

