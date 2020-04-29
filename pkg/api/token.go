package api

import (
	"context"
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type TokenProvider interface {
	VerifyIdentityToken(ctx context.Context, jwt string) (*SessionInfo, error)
	GetIdentityToken(ctx context.Context) (*IdentityToken, error)
}

//interface only for partner tokens
type PartnerTokenProvider interface {
	GetJWTToken(ctx context.Context) (*JwtToken, error)
}

//VerifyIdentityToken ...
func (cli *erplyClient) VerifyIdentityToken(ctx context.Context, jwt string) (*SessionInfo, error) {
	method := VerifyIdentityTokenMethod
	params := map[string]string{
		//params.Add("request", method)
		//params.Add("clientCode", cli.clientCode)
		//params.Add("setContentType", "1")
		"jwt": jwt,
	}
	resp, err := cli.sendRequest(ctx, method, params)
	if err != nil {
		return nil, err
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
func (cli *erplyClient) GetIdentityToken(ctx context.Context) (*IdentityToken, error) {
	method := GetIdentityToken

	resp, err := cli.sendRequest(ctx, method, map[string]string{})
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

//only for partnerClient
func (cli *erplyClient) GetJWTToken(ctx context.Context) (*JwtToken, error) {

	resp, err := cli.sendRequest(ctx, GetJWTTokenMethod, map[string]string{})
	if err != nil {
		return nil, err
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
