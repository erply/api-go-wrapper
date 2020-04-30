package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type (
	verifyIdentityTokenResponse struct {
		Status common.Status `json:"status"`
		Result SessionInfo   `json:"records"`
	}

	SessionInfo struct {
		SessionKey string `json:"sessionKey"`
	}

	getIdentityTokenResponse struct {
		Status common.Status `json:"status"`
		Result IdentityToken `json:"records"`
	}
	IdentityToken struct {
		Jwt string `json:"identityToken"`
	}
	JwtTokenResponse struct {
		Status  common.Status `json:"status"`
		Records JwtToken      `json:"records"`
	}
	JwtToken struct {
		Token string `json:"token"`
	}

	Provider interface {
		VerifyIdentityToken(ctx context.Context, jwt string) (*SessionInfo, error)
		GetIdentityToken(ctx context.Context) (*IdentityToken, error)
	}

	//interface only for partner tokens
	PartnerTokenProvider interface {
		GetJWTToken(ctx context.Context) (*JwtToken, error)
	}

	Client        struct{ *common.Client }
	PartnerClient struct{ *common.Client }
)

func NewClient(client *common.Client) *Client {

	cli := &Client{
		client,
	}
	return cli
}

func NewPartnerClient(client *common.Client) *PartnerClient {
	return &PartnerClient{client}
}

//VerifyIdentityToken ...
func (cli *Client) VerifyIdentityToken(ctx context.Context, jwt string) (*SessionInfo, error) {
	method := "verifyIdentityToken"
	params := map[string]string{
		//params.Add("request", method)
		//params.Add("clientCode", cli.clientCode)
		//params.Add("setContentType", "1")
		"jwt": jwt,
	}
	resp, err := cli.SendRequest(ctx, method, params)
	if err != nil {
		return nil, err
	}
	res := &verifyIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Result, nil
}

//GetIdentityToken ...
func (cli *Client) GetIdentityToken(ctx context.Context) (*IdentityToken, error) {
	method := "getIdentityToken"

	resp, err := cli.SendRequest(ctx, method, map[string]string{})
	if err != nil {
		return nil, erro.NewFromError(fmt.Sprintf("%s request failed", method), err)
	}
	res := &getIdentityTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError(fmt.Sprintf("unmarshaling %s response failed", method), err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Result, nil
}

//only for partnerClient
func (cli *PartnerClient) GetJWTToken(ctx context.Context) (*JwtToken, error) {

	resp, err := cli.SendRequest(ctx, "getJwtToken", map[string]string{})
	if err != nil {
		return nil, err
	}
	var res JwtTokenResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, erro.NewFromError("error decoding GetJWTToken response", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return &res.Records, nil
}
