package addresses

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"strconv"
)

type (
	//GetAddressesResponse ..
	Response struct {
		Status    common.Status `json:"status"`
		Addresses Addresses     `json:"records"`
	}

	//Addresses from getAddresses
	Addresses []Address

	//Address from getAddresses
	Address struct {
		AddressID  int         `json:"addressID"`
		OwnerID    int         `json:"ownerID"`
		TypeID     interface{} `json:"typeID"`
		Address2   string      `json:"address2"`
		Address    string      `json:"address"`
		Street     string      `json:"street"`
		PostalCode string      `json:"postalCode"`
		City       string      `json:"city"`
		State      string      `json:"state"`
		Country    string      `json:"country"`
	}

	AddressManager interface {
		GetAddresses(ctx context.Context, filters map[string]string) ([]Address, error)
		SaveAddress(ctx context.Context, filters map[string]string) ([]Address, error)
	}

	Client struct {
		*common.Client
	}
)

func NewClient(sk, cc, partnerKey string, httpCli *http.Client) *Client {

	cli := &Client{
		common.NewClient(sk, cc, partnerKey, httpCli),
	}
	return cli
}
func (cli *Client) GetAddresses(ctx context.Context, filters map[string]string) ([]Address, error) {
	resp, err := cli.SendRequest(ctx, "getAddresses", filters)
	if err != nil {
		return nil, erro.NewFromError("GetAddresses request failed", err)
	}

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("unmarshaling GetAddressesResponse failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Addresses, nil
}
func (cli *Client) SaveAddress(ctx context.Context, filters map[string]string) ([]Address, error) {
	method := "saveAddress"
	resp, err := cli.SendRequest(ctx, method, filters)
	if err != nil {
		return nil, erro.NewFromError(method+": request failed", err)
	}
	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError(method+": JSON unmarshal failed", err)
	}

	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Addresses) == 0 {
		return nil, erro.NewFromError(method+": no records in response", nil)
	}

	return res.Addresses, nil
}
