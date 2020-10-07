package common

import (
	"context"
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type BulkInput struct {
	MethodName string
	Filters    map[string]interface{}
}

func IsJSONResponseOK(responseStatus *common.Status) bool {
	return strings.EqualFold(responseStatus.ResponseStatus, "ok")
}

func getHTTPRequest(cli *Client, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", cli.Url, body)
	if err != nil {
		return nil, erro.NewFromError("failed to build HTTP request", err)

	}
	return req, err
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
	req, err := getHTTPRequest(cli, nil)
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
	bulkRequest := make([]map[string]interface{}, 0, len(inputs))
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

	var params url.Values
	if cli.headersFunc != nil {
		params = cli.headersFunc("")
		params.Del("request")
	} else {
		params = make(url.Values)
	}

	setParams(params, filters)

	req, err := getHTTPRequest(cli, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, erro.NewFromError("failed to build http request", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req = req.WithContext(ctx)

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

func ConvertSourceToJsonStrIfPossible(source interface{}) string {
	data, err := json.Marshal(source)
	if err != nil {
		return fmt.Sprintf("%+v", source)
	}

	return string(data)
}

func Die(err error) {
	if err != nil {
		panic(err)
	}
}
