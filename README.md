ERPLY API Go SDK
--------
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/erply/api-go-wrapper/pkg/api?tab=doc)
[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://learn-api.erply.com/)

This SDK covers the [ERPLY API](https://erply.com/erply-api/) requests. Majority of the available requests can be checked from [IClient](https://pkg.go.dev/github.com/erply/api-go-wrapper/pkg/api?tab=doc#IClient) file, but some are also [here](https://pkg.go.dev/github.com/erply/api-go-wrapper/pkg/api?tab=doc#pkg-index) and [here](https://pkg.go.dev/github.com/erply/api-go-wrapper/pkg/api?tab=doc#CreateInstallation).

Install
-------
   `go get github.com/erply/api-go-wrapper`
   
Clients
--------
There are 2 ways of using the API. 
* One is you create a `Partner Client` that will always use the partner key with requests and will have access to the requests that require the partner key.
* You can use the simple `Client` that will work without the partner key also.

You can find the example in the `examples` directory for the client initialization process
Example usage as a service
-------
```go
import (
	"github.com/pkg/errors"
	"github.com/erply/api-go-wrapper/pkg/api"
	"strconv"
	"strings"
)

type erplyApiService struct {
	api.IClient
}

func NewErplyApiService(sessionKey, clientCode string) *erplyApiService {
	return &erplyApiService{api.NewClient(sessionKey, clientCode, nil)}
}

//getPointsOfSale erply API request
func (s *erplyApiService) getPointsOfSale(posID string) (string, error) {
	res, err := s.GetPointsOfSaleByID(posID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(res.WarehouseID), nil
}

//verifyIdentityToken erply API request
func (s *erplyApiService) verifyIdentityToken(jwt string) (string, error) {
	res, err := s.VerifyIdentityToken(jwt)
	if err != nil {
		if strings.Contains(err.Error(), "1000") {
			return "", errors.New("jwt expired")
		}
	}
	return res.SessionKey, nil
}

//getIdentityToken erply API request
func (s *erplyApiService) getIdentityToken() (string, error) {
	res, err := s.GetIdentityToken()
	if err != nil {
		if strings.Contains(err.Error(), "1054") {
			return "", errors.New("API session key expired")
		}
	}
	return res.Jwt, nil
}
```

Contributing
-------
I would like to cover the entire ERPLY API and contributions are of course always welcome. The calling pattern is pretty well established, so adding new methods is relatively straightforward. 

Authors
-------
[David Zingerman](https://github.com/Dysar)
