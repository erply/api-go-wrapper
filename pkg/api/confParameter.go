package api

import (
	"context"
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

func (cli *erplyClient) GetConfParameters(ctx context.Context) (*ConfParameter, error) {

	resp, err := cli.sendRequest(ctx, GetConfParametersMethod, map[string]string{})
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

type (
	ConfParameter struct {
		Announcement         string `json:"invoice_announcement_eng"`
		InvoiceClientIsPayer string `json:"invoice_client_is_payer"`
	}
	//GetConfParametersResponse ...
	GetConfParametersResponse struct {
		Status         Status          `json:"status"`
		ConfParameters []ConfParameter `json:"records"`
	}

	ConfManager interface {
		GetConfParameters(ctx context.Context) (*ConfParameter, error)
	}
)
