package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

//VerifyUser will give you session key
func VerifyUser(username, password, clientCode string, client *http.Client) (string, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	params.Add("username", username)
	params.Add("clientCode", clientCode)
	params.Add("password", password)
	params.Add("request", "verifyUser")

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return "", erro.NewFromError("failed to build HTTP request", err)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return "", erro.NewFromError("failed to build VerifyUser request", err)
	}

	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", erro.NewFromError("failed to decode VerifyUserResponse", err)
	}
	if len(res.Records) < 1 {
		return "", erro.NewFromError("VerifyUser: no records in response", nil)
	}
	return res.Records[0].SessionKey, nil
}

//pass filters (including clientCode and sessionKey), pass client code, context and http client
func VerifyUserV2(ctx context.Context, filters map[string]string, clientCode string, cli *http.Client) (string, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	for k, v := range filters {
		params.Add(k, v)
	}
	params.Add("request", "verifyUser")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl, nil)
	if err != nil {
		return "", erro.NewFromError("failed to build HTTP request", err)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")
	resp, err := cli.Do(req)

	if err != nil {
		return "", erro.NewFromError("failed to build VerifyUser request", err)
	}
	res := &VerifyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", erro.NewFromError("failed to decode VerifyUserResponse", err)
	}
	if len(res.Records) < 1 {
		return "", erro.NewFromError("VerifyUser: no records in response", nil)
	}
	return res.Records[0].SessionKey, nil
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//GetSessionKeyUser returns user information for the used session key
func GetSessionKeyUser(sessionKey string, clientCode string, client HttpClient) (*SessionKeyUser, error) {
	requestUrl := fmt.Sprintf(common.BaseUrl, clientCode)
	params := url.Values{}
	params.Add("sessionKey", sessionKey)
	params.Add("doNotGenerateIdentityToken", "1")
	params.Add("request", "getSessionKeyUser")
	params.Add("clientCode", clientCode)

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return nil, erro.NewFromError("failed to build HTTP request", err)
	}

	req.URL.RawQuery = params.Encode()
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, erro.NewFromError("failed to call getSessionKeyUser request", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body := []byte{}
		if resp.Body != nil {
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				body = []byte{}
			}
		}

		return nil, fmt.Errorf("wrong response status code: %d, body: %s", resp.StatusCode, string(body))
	}

	res := &SessionKeyUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to decode SessionKeyUserResponse", err)
	}
	if len(res.Records) < 1 {
		return nil, erro.NewFromError("getSessionKeyUser: no records in response", nil)
	}
	return &res.Records[0], nil
}
