package api

import (
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/url"
	"strconv"
)

func (cli *erplyClient) GetConfParameters() (*ConfParameter, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetConfParameters request", err)
	}

	params := url.Values{}
	params.Add("request", GetConfParametersMethod)
	params.Add(sessionKey, cli.sessionKey)
	params.Add(clientCode, cli.clientCode)
	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetConfParameters request failed", err)
	}
	res := &GetConfParametersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetConfParametersResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.ConfParameters) == 0 {
		return nil, erplyerr(fmt.Sprint("Conf Parameters were not found", nil), err)
	}

	return &res.ConfParameters[0], nil
}

type ConfParameter struct {
	Announcement         string `json:"invoice_announcement_eng"`
	InvoiceClientIsPayer string `json:"invoice_client_is_payer"`
}
