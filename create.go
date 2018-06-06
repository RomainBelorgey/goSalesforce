package goSalesforce

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// Create a comment on a case on salesforce
// Need the URL and the sessionId from Sfauth
// Need to provide the caseId related, the comment text and if it need to be published (public)
// See the example to understand this function
// Will return a bool if the creation was done
func SfCreateComment(urlSf string, sessionId string, caseId string, comment string, isPublished bool) bool {
	content := fmt.Sprintf("{"+
		"\"ParentId\" : \"%s\","+
		"\"CommentBody\" : %q,"+
		"\"IsPublished\" : %v"+
		"}", caseId, comment, isPublished)

	return sfCreateBack(urlSf, sessionId, "CaseComment", content)
}

func sfCreateBack(urlSf string, sessionId string, typeCreate string, content string) bool {

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlSf+"/services/data/v39.0/sobjects/"+typeCreate+"/", bytes.NewReader([]byte(content)))
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
