package api

import (
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

func (cli *erplyClient) GetServiceEndpoints() (*ServiceEndpoints, error) {
	method := "getServiceEndpoints"
	params := url.Values{}
	params.Add("clientCode", cli.clientCode)
	params.Add("request", method)

	req, err := newPostHTTPRequest(cli, params)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("failed to build %s request", method), err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr(fmt.Sprintf("%s request failed", method), err)
	}

	res := &getServiceEndpointsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, "failed to decode VerifyUserResponse")
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	if len(res.Records) < 1 {
		return nil, errors.New("no records in response")
	}
	return &res.Records[0], nil
}

type getServiceEndpointsResponse struct {
	Status  Status
	Records []ServiceEndpoints `json:"records"`
}

type ServiceEndpoints struct {
	Cafa        Endpoint `json:"cafa"`
	Pim         Endpoint `json:"pim"`
	Wms         Endpoint `json:"wms"`
	Promotion   Endpoint `json:"promotion"`
	Reports     Endpoint `json:"reports"`
	Json        Endpoint `json:"json"`
	Assignments Endpoint `json:"assignments"`
}
type Endpoint struct {
	Url           string `json:"url"`
	Documentation string `json:documentation""`
}
