package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
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

func CreateInstallation(baseUrl, partnerKey string, in *InstallationRequest, httpCli *http.Client) (*InstallationResponse, error) {

	if httpCli == nil {
		return nil, errors.New("no http cli provided")
	}

	params := url.Values{}
	params.Add("request", createInstallationMethod)
	params.Add("partnerKey", partnerKey)
	params.Add("companyName", in.CompanyName)
	params.Add("firstName", in.FirstName)
	params.Add("lastName", in.LastName)
	params.Add("phone", in.Phone)
	params.Add("email", in.Email)
	params.Add("sendEmail", strconv.Itoa(in.SendEmail))

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
		Status  common.Status
		Records []InstallationResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erro.NewFromError("CreateInstallation: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		fmt.Println(respData.Status.ErrorField)
		return nil, erro.NewFromError(fmt.Sprintf("CreateInstallation: API error %d", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return nil, erro.NewFromError("CreateInstallation: no records in response", nil)
	}

	return &respData.Records[0], nil
}
