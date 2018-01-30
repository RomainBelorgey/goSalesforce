package goSalesforce

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// Update a record on salesforce
// Need the URL and the sessionId from Sfauth
// See the example to understand this function
// Will return a bool if the change was done
func SfUpdate(urlSf string, sessionId string, typeUpdate string, idUpdate string, contentName string, contentValue string) bool {

	content := fmt.Sprintf("{"+
		"\"%s\" : \"%s\""+
		"}", contentName, contentValue)

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", urlSf+"/services/data/v41.0/sobjects/"+typeUpdate+"/"+idUpdate, bytes.NewReader([]byte(content)))
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("Authorization", "OAuth "+sessionId)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	} else {
		return false
	}
}