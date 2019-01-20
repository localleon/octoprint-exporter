package main

import (
	"encoding/json"
	"fmt"
)

func apiJobProgress() float64 {
	apiGetJobInfo()
	return 1.0
}

func apiGetJobInfo() map[string]interface{} {
	// Call API for job information
	apiurl := "http://" + c.Octopi + "/api/job"
	resp := makeGetRequest(apiurl)
	// Parse Json
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	mapString := make(map[string]string)
	for key, value := range result {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)

		mapString[strKey] = strValue
	}
	fmt.Printf("New parser")
	fmt.Printf("%#v", mapString["job"])
	return result
}
