package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type (
	//GetAddressesResponse ..
	AddressesResponse struct {
		Status    Status    `json:"status"`
		Addresses Addresses `json:"records"`
	}

	//Addresses from getAddresses
	Addresses []Address

	AddressRequest struct {
		AddressID  int    `json:"addressID"`
		OwnerID    int    `json:"ownerID"`
		TypeID     int    `json:"typeID"`
		Address2   string `json:"address2"`
		Address    string `json:"address"`
		Street     string `json:"street"`
		PostalCode string `json:"postalCode"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
	}

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
		SaveAddress(ctx context.Context, in *AddressRequest) ([]Address, error)
	}
)

func (cli *erplyClient) GetAddresses(ctx context.Context, filters map[string]string) ([]Address, error) {

	resp, err := cli.sendRequest(ctx, GetAddressesMethod, filters)
	if err != nil {
		return nil, erplyerr("GetAddresses request failed", err)
	}

	res := &AddressesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetAddressesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Addresses, nil
}
func (cli *erplyClient) SaveAddress(ctx context.Context, in *AddressRequest) ([]Address, error) {
	filters := map[string]string{
		"addressID":  strconv.Itoa(in.AddressID),
		"typeID":     strconv.Itoa(in.TypeID),
		"ownerID":    strconv.Itoa(in.OwnerID),
		"street":     in.Street,
		"postalCode": in.PostalCode,
		"city":       in.City,
		"state":      in.State,
		"country":    in.Country,
	}

	resp, err := cli.sendRequest(ctx, saveAddressMethod, filters)
	if err != nil {
		return nil, erplyerr(saveAddressMethod+": request failed", err)
	}
	res := &AddressesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr(saveAddressMethod+": JSON unmarshal failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Addresses) == 0 {
		return nil, erplyerr(saveAddressMethod+": no records in response", nil)
	}

	return res.Addresses, nil
}
