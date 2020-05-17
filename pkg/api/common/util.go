package common

import (
	"context"
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"net/http"
	"net/url"
	"strings"
)

type BulkInput struct {
	MethodName string
	Filters    map[string]string
}

func IsJSONResponseOK(responseStatus *Status) bool {
	return strings.EqualFold(responseStatus.ResponseStatus, "ok")
}

func getHTTPRequest(cli *Client) (*http.Request, error) {
	req, err := http.NewRequest("POST", cli.Url, nil)
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

func (cli *Client) getDefaultMandatoryHeaders(request string) url.Values {
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
	params := cli.headersFunc(apiMethod)
	setParams(params, filters)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erro.NewFromError(fmt.Sprintf("%v request failed", apiMethod), err)
	}
	return resp, nil
}

func (cli *Client) SendRequestBulk(ctx context.Context, inputs []BulkInput, filters map[string]string) (*http.Response, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erro.NewFromError("failed to build http request", err)
	}

	bulkRequest := make([]map[string]string, 0, len(inputs))
	for _, input := range inputs {
		bulkItemFilters := input.Filters
		bulkItemFilters["requestName"] = input.MethodName

		bulkRequest = append(bulkRequest, bulkItemFilters)
	}

	jsonRequests, err := json.Marshal(bulkRequest)
	if err != nil {
		return nil, erro.NewFromError("failed to build requests payload", err)
	}

	filters["requests"] = string(jsonRequests)

	req = req.WithContext(ctx)
	params := cli.headersFunc("")
	params.Del("request")
	setParams(params, filters)

	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erro.NewFromError("Bulk request failed", err)
	}
	return resp, nil
}

func doRequest(req *http.Request, cli *Client) (*http.Response, error) {
	resp, err := cli.httpClient.Do(req)
	return resp, err
}
