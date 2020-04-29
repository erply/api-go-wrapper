package api

import (
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type (
	//GetAddressesResponse ..
	GetAddressesResponse struct {
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
)

func (cli *erplyClient) GetAddresses(filters map[string]string) ([]Address, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("failed to build GetAddresses request", err)
	}

	params := getMandatoryParameters(cli, GetAddressesMethod)
	for fk, fv := range filters {
		params.Add(fk, fv)
	}

	req.URL.RawQuery = params.Encode()
	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("GetAddresses request failed", err)
	}

	res := &GetAddressesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("unmarshaling GetAddressesResponse failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	return res.Addresses, nil
}
func (cli *erplyClient) SaveAddress(in *AddressRequest) (int, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return 0, erplyerr("SaveAddress: failed to build request", err)
	}
	params := getMandatoryParameters(cli, saveAddressMethod)
	params.Add("addressID", strconv.Itoa(in.AddressID))
	params.Add("typeID", strconv.Itoa(in.TypeID))
	params.Add("ownerID", strconv.Itoa(in.OwnerID))
	params.Add("street", in.Street)
	params.Add("postalCode", in.PostalCode)
	params.Add("city", in.City)
	params.Add("state", in.State)
	params.Add("country", in.Country)

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return 0, erplyerr("SaveAddress: request failed", err)
	}

	var res struct {
		Status  Status
		Records []Address
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, erplyerr("SaveAddress: JSON unmarshal failed", err)
	}

	if !isJSONResponseOK(&res.Status) {
		return 0, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}

	if len(res.Records) == 0 {
		return 0, erplyerr("SaveAddress: no records in response", nil)
	}

	return res.Records[0].AddressID, nil
}
