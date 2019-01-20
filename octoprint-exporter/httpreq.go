package main

import (
	"log"
	"net/http"
)

func makeGetRequest(url string) *http.Response {
	// Create Request with correct headers for authorization
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Api-Key", c.Apikey)
	resp, err := client.Do(req)
	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
