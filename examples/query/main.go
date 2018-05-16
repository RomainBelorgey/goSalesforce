package main

import (
	"fmt"
	"github.com/RomainBelorgey/goSalesforce"
)

const (
	url      = "your-salesforce-url"
	login    = "your-login"
	password = "your-password"
	token    = "your-token"
)

func main() {
	sessionId, err := goSalesforce.SfAuth(url, login, password, token)
	if err != nil {
		result := goSalesforce.SfQuery(url, sessionId, "SELECT CaseNumber from CASE")
		fmt.Println(string(result))
	} else {
		fmt.Println(err)
	}
}
