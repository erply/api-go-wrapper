package api

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(partnerKey string, timestamp int64, request, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(fmt.Sprintf("%s%s%d", partnerKey, request, timestamp)))
	return hex.EncodeToString(h.Sum(nil))

}
func isJSONResponseOK(status *Status) bool {
	return strings.EqualFold(status.ResponseStatus, responseStatusOK)
}

func getHTTPRequest(cli *erplyClient) (*http.Request, error) {
	req, err := http.NewRequest("GET", cli.url, nil)
	if err != nil {
		return nil, erplyerr("failed to build HTTP request", err)

	}
	return req, err
}

func newPostHTTPRequest(cli *erplyClient, params url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", cli.url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, erplyerr("failed to build HTTP request", err)

	}
	return req, err
}

func getMandatoryParameters(cli *erplyClient, request string) url.Values {
	params := url.Values{}
	params.Add("request", request)
	params.Add("setContentType", "1")
	if cli.sessionKey != "" && cli.clientCode != "" {
		params.Add(sessionKey, cli.sessionKey)
		params.Add(clientCode, cli.clientCode)
	}
	if cli.partnerKey != "" && cli.secret != "" {
		now := time.Now().Unix()
		params.Add(applicationKey, GenerateToken(cli.partnerKey, now, request, cli.secret))
		params.Add(clientCode, cli.clientCode)
		params.Add("partnerKey", cli.partnerKey)
		params.Add("timestamp", strconv.Itoa(int(now)))
	}
	return params
}

func erplyerr(msg string, err error) *erro.ErplyError {
	if err != nil {
		return erro.NewErplyError("Error", errors.Wrap(err, msg).Error())
	}
	return erro.NewErplyError("Error", msg)
}

func setParams(params url.Values, filters map[string]string) {
	for k, v := range filters {
		params.Add(k, v)
	}
}

func (cli *erplyClient) sendRequest(ctx context.Context, apiMethod string, filters map[string]string) (*http.Response, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build http request", err)
	}
	req = req.WithContext(ctx)
	params := getMandatoryParameters(cli, apiMethod)
	setParams(params, filters)
	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%v request failed", apiMethod), err)
	}
	return resp, nil
}
func doRequest(req *http.Request, cli *erplyClient) (*http.Response, error) {
	resp, err := cli.httpClient.Do(req)
	return resp, err
}
