package finfree_http_client

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	baseURL := "https://3bcf-95-70-132-33.eu.ngrok.io"

	client := New(baseURL)
	response, err := client.Get("/users", nil, NewQuery("name", "john"), NewQuery("lastname", "wick"))
	if err != nil {
		log.Println("Error on get request. Err:", err)
		t.FailNow()
	}

	log.Println("[GET] Response ->", response)

}

func TestPost(t *testing.T) {
	baseURL := "https://3bcf-95-70-132-33.eu.ngrok.io"

	client := New(baseURL)

	requestBody := map[string]any{
		"name":     "John",
		"lastname": "Wick",
	}

	response, err := client.Post("/user", &requestBody, nil)
	if err != nil {
		log.Println("Error on post request. Err ->", err)
		t.FailNow()
	}

	log.Println("Response ->", response)

}

func TestPut(t *testing.T) {
	baseURL := "https://3bcf-95-70-132-33.eu.ngrok.io"

	client := New(baseURL)

	requestBody := map[string]any{
		"name":     "John",
		"lastname": "Wick",
	}

	response, err := client.Put("/user", &requestBody, nil)
	if err != nil {
		log.Println("Error on post request. Err ->", err)
		t.FailNow()
	}

	log.Println("Response ->", response)
}

func TestPatch(t *testing.T) {
	baseURL := "https://3bcf-95-70-132-33.eu.ngrok.io"

	client := New(baseURL)

	requestBody := map[string]any{
		"name":     "John",
		"lastname": "Wick",
	}

	response, err := client.Patch("/user", &requestBody, nil)
	if err != nil {
		log.Println("Error on post request. Err ->", err)
		t.FailNow()
	}

	log.Println("Response ->", response)

}

func TestDelete(t *testing.T) {
	baseURL := "https://3bcf-95-70-132-33.eu.ngrok.io"

	client := New(baseURL)

	response, err := client.Delete("/user", nil, NewQuery("name", "John"), NewQuery("lastname", "Wick"))
	if err != nil {
		log.Println("Error on post request. Err ->", err)
		t.FailNow()
	}

	log.Println("Response ->", response)
}
