package api

import (
	"net/http"
	"testing"
)

//works
func TestVerifyUser(t *testing.T) {
	const (
		username = ""
		password = ""
		cc       = ""
	)
	sk, err := VerifyUser(username, password, cc, &http.Client{})

	if err != nil {
		t.Error(err)
		return
	}
	if sk == "" {
		t.Error("did not get sk")
		return
	}
}
