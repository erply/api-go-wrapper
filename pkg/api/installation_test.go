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
			CompanyName: "",
			FirstName:   "",
			LastName:    "",
			Phone:       "",
			Email:       "",
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
