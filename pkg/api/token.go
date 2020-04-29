package api

import (
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/url"
	"strconv"
)

//VerifyIdentityToken ...
func (cli *erplyClient) VerifyIdentityToken(jwt string) (*SessionInfo, error) {
	method := VerifyIdentityTokenMethod
	params := url.Values{}
	params.Add("request", method)
	params.Add("clientCode", cli.clientCode)
	params.Add("setContentType", "1")
	params.Add("jwt", jwt)
	req, err := newPostHTTPRequest(cli, params)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}

	res := &verifyIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Result, nil
}

//GetIdentityToken ...
func (cli *erplyClient) GetIdentityToken() (*IdentityToken, error) {
	method := GetIdentityToken

	params := getMandatoryParameters(cli, method)
	queryParams := getMandatoryParameters(cli, method)

	req, err := newPostHTTPRequest(cli, params)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}
	req.URL.RawQuery = queryParams.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}
	res := &getIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Result, nil
}

func (cli *erplyClient) GetJWTToken(partnerKey string) (*JwtToken, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, fmt.Errorf("error building GetJWTToken request: %v", err)
	}

	params := getMandatoryParameters(cli, GetJWTTokenMethod)
	params.Set("partnerKey", partnerKey)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("error making request for GetJWTToken", err)
	}

	var res JwtTokenResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, erplyerr("error decoding GetJWTToken response", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Records, nil
}
