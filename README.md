ERPLY API Go SDK
--------
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/erply/api-go-wrapper/pkg/api?tab=doc)
[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://learn-api.erply.com/)

This SDK covers the [ERPLY API](https://erply.com/erply-api/) requests. 

Client Structure
------
Majority of the request wrappers are available through the client.
The client is described in `GoDoc` type `Client` and in `/pkg/api/client.go`. It is divided into sub-clients for each topic that the underlying API covers. 
For now not all the requests are mapped to topics. Such request wrappers are in `/pkg/api` directory. 
Some requests are accessible not from the client, but from the `auth` package of this SDK. They are covered in the example in `/examples` directory.

Install
-------
   `go get github.com/erply/api-go-wrapper@X.Y.Z`
   
   where X.Y.Z is your desired version.
   
Clients
--------
Ways of using the API: 
* One is you create a `Partner Client` that will always use the partner key with requests and will have access to the requests that require the partner key.
* You can use the simple `Client` that will work without the partner key as well.
* You can also create a client that can act like a partner client, normal one and it is possible to define the headers that will be added for every request on your own. For that one please use the `NewClientWithCustomHeaders` constructor.

You can find the example in the `/examples` directory for the client initialization process

Contributing
-------
This library is not in the final state, and it means for continuous development. Therefore, I would like to cover the entire ERPLY API and contributions are of course always welcome. The calling pattern is pretty well-established, so adding new methods is relatively straightforward. 

