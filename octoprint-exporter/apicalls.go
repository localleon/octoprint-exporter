package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func apiGetPrinterInfo() PrinterInfoJSON {
	// Call API for job information
	apiurl := "http://" + c.Octopi + "/api/printer?exclude=state"
	// Create Request with correct headers for authorization
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	req.Header.Add("X-Api-Key", c.Apikey)
	resp, err := client.Do(req)
	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// Create Struct to store json and parse the response
	var infos PrinterInfoJSON
	json.Unmarshal(body, &infos)
	return infos
}

// PrinterInfoJSON represents the JSON Response from this Octoprint API Endpoint
type PrinterInfoJSON struct {
	Sd struct {
		Ready bool `json:"ready"`
	} `json:"sd"`
	Temperature struct {
		Bed struct {
			Actual float64 `json:"actual"`
			Offset int     `json:"offset"`
			Target float64 `json:"target"`
		} `json:"bed"`
		Tool0 struct {
			Actual float64 `json:"actual"`
			Offset int     `json:"offset"`
			Target float64 `json:"target"`
		} `json:"tool0"`
	} `json:"temperature"`
}

func apiGetJobInfo() JobInfoJSON {
	// Call API for job information
	apiurl := "http://" + c.Octopi + "/api/job"
	// Create Request with correct headers for authorization
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	req.Header.Add("X-Api-Key", c.Apikey)
	resp, err := client.Do(req)
	// Check for errors
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// Create Struct to store json and parse the response
	var infos JobInfoJSON
	json.Unmarshal(body, &infos)
	return infos
}

// JobInfoJSON represents the JSON Response from this Octoprint API Endpoint
type JobInfoJSON struct {
	Job struct {
		AveragePrintTime   interface{} `json:"averagePrintTime"`
		EstimatedPrintTime float64     `json:"estimatedPrintTime"`
		Filament           struct {
			Tool0 struct {
				Length float64 `json:"length"`
				Volume float64 `json:"volume"`
			} `json:"tool0"`
		} `json:"filament"`
		File struct {
			Date    int    `json:"date"`
			Display string `json:"display"`
			Name    string `json:"name"`
			Origin  string `json:"origin"`
			Path    string `json:"path"`
			Size    int    `json:"size"`
		} `json:"file"`
		LastPrintTime interface{} `json:"lastPrintTime"`
	} `json:"job"`
	Progress struct {
		Completion          float64 `json:"completion"`
		Filepos             int     `json:"filepos"`
		PrintTime           int     `json:"printTime"`
		PrintTimeLeft       int     `json:"printTimeLeft"`
		PrintTimeLeftOrigin string  `json:"printTimeLeftOrigin"`
	} `json:"progress"`
	State string `json:"state"`
}
