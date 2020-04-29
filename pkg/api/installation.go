package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func CreateInstallation(baseUrl, partnerKey string, in *InstallationRequest, cli *http.Client) (*InstallationResponse, error) {

	params := url.Values{}
	params.Add("request", createInstallationMethod)
	params.Add("partnerKey", partnerKey)
	params.Add("companyName", in.CompanyName)
	params.Add("firstName", in.FirstName)
	params.Add("lastName", in.LastName)
	params.Add("phone", in.Phone)
	params.Add("email", in.Email)
	params.Add("sendEmail", strconv.Itoa(in.SendEmail))

	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, erplyerr("failed to build HTTP request", err)

	}
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, &erplyClient{httpClient: cli})
	if err != nil {
		return nil, erplyerr("CreateInstallation: error sending POST request", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, erplyerr(fmt.Sprintf("CreateInstallation: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []InstallationResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erplyerr("CreateInstallation: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		fmt.Println(respData.Status.ErrorField)
		return nil, erplyerr(fmt.Sprintf("CreateInstallation: API error %d", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return nil, erplyerr("CreateInstallation: no records in response", nil)
	}

	return &respData.Records[0], nil
}
