package main

import (
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
)

func main() {
	const (
		username   = ""
		password   = ""
		clientCode = ""
		partnerKey = ""
	)
	httpCli := common.GetDefaultHTTPClient()

	//get the session key
	sessionKey, err := auth.VerifyUser(username, password, clientCode, httpCli)
	if err != nil {
		panic(err)
	}

	//GetSessionKeyInfo allows you to get more information about the session if needed
	sessInfo, err := auth.GetSessionKeyInfo(sessionKey, clientCode, httpCli)
	if err != nil {
		panic(err)
	}
	fmt.Println(sessInfo)

	//GetSessionKeyUser
	info, _ := auth.GetSessionKeyUser(sessionKey, clientCode, httpCli)
	fmt.Println(info)
}
