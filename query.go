package goSalesforce

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// Do a query to salesforce
// Need the URL and the sessionId from Sfauth
// I will return a json content
func SfQuery(urlSf string, sessionId string, query string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlSf+"/services/data/v41.0/query/", bytes.NewReader(nil))
	if err != nil {
		log.Println("Can't send request : %s", err)
	}
	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "OAuth "+sessionId)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Can't read response request : %s", err)
	}
	return contents

}

// Download the content on the path
// Useful for downloading attachments
// Return the file content in byte
func SfDownload(urlSf string, sessionId string, path string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlSf+""+path, bytes.NewReader(nil))
	if err != nil {
		log.Println("Can't send request : %s", err)
	}
	req.Header.Add("Authorization", "OAuth "+sessionId)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Can't read response request : %s", err)
	}
	return contents
}
