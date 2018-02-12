package goSalesforce

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Do a query to salesforce
// Need the URL and the sessionId from Sfauth
// I will return a json content
func SfQuery(urlSf string, sessionId string, query string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlSf+"/services/data/v41.0/query/", bytes.NewReader(nil))
	if err != nil {
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "OAuth "+sessionId)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
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
		os.Exit(1)
	}
	req.Header.Add("Authorization", "OAuth "+sessionId)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return contents
}
