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
		//Will change the case with id your-caseid to priority normal
		fmt.Println(goSalesforce.SfUpdate(url, sessionId, "Case", "your-caseid", "Priority", "Normal"))
	} else {
		fmt.Println(err)
	}
}
