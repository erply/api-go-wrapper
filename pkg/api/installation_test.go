package api

import (
	"net/http"
	"testing"
	"time"
)

//works
func TestCreateInstallation(t *testing.T) {
	const (
		baseUrl    = ""
		partnerKey = ""
	)
	var (
		req = &InstallationRequest{
			CompanyName: "aaa",
			FirstName:   "aasd",
			LastName:    "asdasd",
			Phone:       "asd",
			Email:       "asd@asd.ee",
			SendEmail:   0,
		}
		cli = &http.Client{Timeout: 10 * time.Second}
	)

	resp, err := CreateInstallation(baseUrl, partnerKey, req, cli)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
