package common

import (
	"context"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"net/url"
	"strings"
)

func IsJSONResponseOK(responseStatus *Status) bool {
	return strings.EqualFold(responseStatus.ResponseStatus, "ok")
}

func getHTTPRequest(cli *Client) (*http.Request, error) {
	req, err := http.NewRequest("POST", cli.url, nil)
	if err != nil {
		return nil, erro.NewFromError("failed to build HTTP request", err)

	}
	return req, err
}

func GetBaseURL(cc string) string {
	return fmt.Sprintf(BaseUrl, cc)
}

const (
	clientCode = "clientCode"
	sessionKey = "sessionKey"
)

func getMandatoryParameters(cli *Client, request string) url.Values {
	params := url.Values{}
	params.Add("request", request)
	params.Add("setContentType", "1")
	params.Add(sessionKey, cli.sessionKey)
	params.Add(clientCode, cli.clientCode)
	if cli.partnerKey != "" {
		params.Add("partnerKey", cli.partnerKey)
	}
	return params
}

func setParams(params url.Values, filters map[string]string) {
	for k, v := range filters {
		params.Add(k, v)
	}
}

func (cli *Client) SendRequest(ctx context.Context, apiMethod string, filters map[string]string) (*http.Response, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erro.NewFromError("failed to build http request", err)
	}
	req = req.WithContext(ctx)
	params := getMandatoryParameters(cli, apiMethod)
	setParams(params, filters)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erro.NewFromError(fmt.Sprintf("%v request failed", apiMethod), err)
	}
	return resp, nil
}
func doRequest(req *http.Request, cli *Client) (*http.Response, error) {
	resp, err := cli.httpClient.Do(req)
	return resp, err
}
