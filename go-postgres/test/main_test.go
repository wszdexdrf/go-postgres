package main_test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
)

type Response struct {
	Message string
}

// to run tests, first run the server using `npm run start`
// from go-postgres directory
func TestCreate(t *testing.T) {
	resp, err := http.Get("http://localhost:3001/create")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := Response{}
	json.Unmarshal([]byte(body), &data)
	if data.Message != "Create Success" {
		log.Fatalf("Message received: %s", data.Message)
	}
}

func TestRead(t *testing.T) {
	resp, err := http.Get("http://localhost:3001/read")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := Response{}
	json.Unmarshal([]byte(body), &data)
	if data.Message != "Read Success" {
		log.Fatalf("Message received: %s", data.Message)
	}
}

func TestUpdate(t *testing.T) {
	resp, err := http.Get("http://localhost:3001/update")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := Response{}
	json.Unmarshal([]byte(body), &data)
	if data.Message != "Update Success" {
		log.Fatalf("Message received: %s", data.Message)
	}
}

func TestDelete(t *testing.T) {
	resp, err := http.Get("http://localhost:3001/delete")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := Response{}
	json.Unmarshal([]byte(body), &data)
	if data.Message != "Delete Success" {
		log.Fatalf("Message received: %s", data.Message)
	}
}
