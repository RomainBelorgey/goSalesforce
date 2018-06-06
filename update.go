package goSalesforce

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Update a record on salesforce
// Need the URL and the sessionId from Sfauth
// See the example to understand this function
// Will return a bool if the change was done
func SfUpdate(urlSf string, sessionId string, typeUpdate string, idUpdate string, contentName string, contentValue string) bool {

	contentValue = strings.Replace(contentValue, `"`, `\"`, -1)
	content := fmt.Sprintf("{"+
		"\"%s\" : \"%s\""+
		"}", contentName, contentValue)

	return sfUpdateBack(urlSf, sessionId, typeUpdate, idUpdate, content)
}

// Update a record on salesforce with a boolean as value
// Need the URL and the sessionId from Sfauth
// See the example to understand this function
// Will return a bool if the change was done
func SfUpdateBool(urlSf string, sessionId string, typeUpdate string, idUpdate string, contentName string, contentValue bool) bool {

	content := fmt.Sprintf("{"+
		"\"%s\" : %v"+
		"}", contentName, contentValue)

	return sfUpdateBack(urlSf, sessionId, typeUpdate, idUpdate, content)
}

func sfUpdateBack(urlSf string, sessionId string, typeUpdate string, idUpdate string, content string) bool {

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", urlSf+"/services/data/v41.0/sobjects/"+typeUpdate+"/"+idUpdate, bytes.NewReader([]byte(content)))
	if err != nil {
		log.Println("Can't prepare request : %s", err)
	}
	req.Header.Add("Authorization", "OAuth "+sessionId)
	//Don't auto-assign case
	req.Header.Add("Sforce-Auto-Assign", "FALSE")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Can't send request : %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	} else {
		return false
	}
}
