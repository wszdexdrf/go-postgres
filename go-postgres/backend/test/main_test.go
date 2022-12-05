package main_test

import (
	"io"
	"log"
	"net/http"
	"testing"
)

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
	if string(body) != "Create Success" {
		log.Fatal(string(body))
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
	if string(body) != "Read Success" {
		log.Fatal(string(body))
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
	if string(body) != "Update Success" {
		log.Fatal(string(body))
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
	if string(body) != "Delete Success" {
		log.Fatal(string(body))
	}
}
