package goSalesforce

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Envelope struct {
	XmlResponse XmlResponse `xml:"Body"`
}
type XmlResponse struct {
	LoginResponse XmlResult `xml:"loginResponse"`
}
type XmlResult struct {
	Result XmlAnswer `xml:"result"`
}
type XmlAnswer struct {
	SessionId string `xml:"sessionId"`
}

// Will return a Oauth SessionID
// The SessionId will be use for queries
func SfAuth(url string, login string, password string, token string) (string, error) {

	loginXml := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\" ?>"+
		"<env:Envelope xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:env=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:urn=\"urn:partner.soap.sforce.com\">"+
		"  <env:Header>"+
		"    <urn:CallOptions>"+
		"      <urn:client>RestForce</urn:client>"+
		"      <urn:defaultNamespace>sf</urn:defaultNamespace>"+
		"    </urn:CallOptions>"+
		"  </env:Header>"+
		"  <env:Body>"+
		"    <n1:login xmlns:n1=\"urn:partner.soap.sforce.com\">"+
		"      <n1:username>%s</n1:username>"+
		"      <n1:password>%s</n1:password>"+
		"    </n1:login>"+
		"  </env:Body>"+
		"</env:Envelope>", login, password+token)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url+"/services/Soap/u/41.0", bytes.NewReader([]byte(loginXml)))
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("SOAPAction", "login")
	req.Header.Add("Content-Type", "text/xml; charset=UTF-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	v := Envelope{}
	err = xml.Unmarshal(contents, &v)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		err = errors.New("Auth failed  : " + string(contents))
	}

	return v.XmlResponse.LoginResponse.Result.SessionId, err
}
