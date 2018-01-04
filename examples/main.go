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
	sessionId := goSalesforce.SfAuth(url, login, password, token)
	result := goSalesforce.SfQuery(url, sessionId, "SELECT CaseNumber from CASE")
	fmt.Println(string(result))
}
