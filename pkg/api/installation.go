package api

import (
	"encoding/json"
	"errors"
	"fmt"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
	"net/http"
	"net/url"
)

type InstallationRequest struct {
	CompanyName string `json:"companyName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	SendEmail   int    `json:"sendEmail"`
}
type InstallationResponse struct {
	ClientCode int    `json:"clientCode"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
}

func CreateInstallation(baseUrl, partnerKey string, filters map[string]string, httpCli *http.Client) (*InstallationResponse, error) {

	if httpCli == nil {
		return nil, errors.New("no http cli provided")
	}

	params := url.Values{}
	for k, v := range filters {
		params.Add(k, v)
	}
	params.Add("request", createInstallationMethod)
	params.Add("partnerKey", partnerKey)

	req, err := http.NewRequest("POST", baseUrl, nil)
	if err != nil {
		return nil, erro.NewFromError("failed to build HTTP request", err)

	}
	req.URL.RawQuery = params.Encode()
	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, erro.NewFromError("CreateInstallation: error sending POST request", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, erro.NewFromError(fmt.Sprintf("CreateInstallation: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  common2.Status
		Records []InstallationResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erro.NewFromError("CreateInstallation: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, erro.NewFromError(fmt.Sprintf("CreateInstallation: API error %s", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return nil, erro.NewFromError("CreateInstallation: no records in response", nil)
	}

	return &respData.Records[0], nil
}
