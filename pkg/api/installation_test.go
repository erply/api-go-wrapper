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
		filters = map[string]string{
			"companyName": "",
			"firstName":   "",
			"lastName":    "",
			"phone":       "",
			"email":       "@.",
			"sendEmail":   "0",
		}
		cli = &http.Client{Timeout: 10 * time.Second}
	)

	resp, err := CreateInstallation(baseUrl, partnerKey, filters, cli)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
