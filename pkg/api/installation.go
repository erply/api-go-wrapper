package api

import (
	"encoding/json"
	"errors"
	"fmt"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
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
		return nil, sharedCommon.NewFromError("failed to build HTTP request", err, 0)

	}
	req.URL.RawQuery = params.Encode()
	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, sharedCommon.NewFromError("CreateInstallation: error sending POST request", err, 0)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CreateInstallation: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	var respData struct {
		Status  sharedCommon.Status
		Records []InstallationResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, sharedCommon.NewFromError("CreateInstallation: error decoding JSON response body", err, 0)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("CreateInstallation: API error %s", respData.Status.ErrorCode), nil, respData.Status.ErrorCode)
	}
	if len(respData.Records) < 1 {
		return nil, sharedCommon.NewFromError("CreateInstallation: no records in response", nil, respData.Status.ErrorCode)
	}

	return &respData.Records[0], nil
}
