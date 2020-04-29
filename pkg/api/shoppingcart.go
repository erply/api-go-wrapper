package api

import (
	"encoding/json"
	"fmt"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"net/http"
	"strconv"
)

type ShoppingCartTotals struct {
	Rows     []ShoppingCartProduct `json:"rows"`
	NetTotal float64               `json:"netTotal"`
	VATTotal float64               `json:"vatTotal"`
	Total    float64               `json:"total"`
}

type ShoppingCartProduct struct {
	ProductID            string  `json:"productID"`
	Amount               string  `json:"amount"`
	OriginalPrice        float64 `json:"originalPrice"`
	OriginalPriceWithVAT float64 `json:"originalPriceWithVAT"`
	FinalPrice           float64 `json:"finalPrice"`
	FinalPriceWithVAT    float64 `json:"finalPriceWithVAT"`
	RowNetTotal          float64 `json:"rowNetTotal"`
	RowTotal             float64 `json:"rowTotal"`
	Discount             float64 `json:"discount"`
}

func (cli *erplyClient) CalculateShoppingCart(in *DocumentData) (*ShoppingCartTotals, error) {
	req, err := getHTTPRequest(cli)
	if err != nil {
		return nil, erplyerr("CalculateShoppingCart: failed to build request", err)
	}

	params := getMandatoryParameters(cli, calculateShoppingCartMethod)
	params.Add("customerID", strconv.FormatUint(uint64(in.CustomerId), 10))

	for i, prod := range in.ProductRows {
		params.Add(fmt.Sprintf("productID%d", i), prod.ProductID)
		params.Add(fmt.Sprintf("amount%d", i), prod.Amount)
	}

	req.URL.RawQuery = params.Encode()

	resp, err := doRequest(req, cli)
	if err != nil {
		return nil, erplyerr("CalculateShoppingCart: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, erplyerr(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []*ShoppingCartTotals
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, erplyerr("CalculateShoppingCart: unmarshaling response failed", err)
	}
	if !isJSONResponseOK(&respData.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(respData.Status.ErrorCode), respData.Status.Request+": "+respData.Status.ResponseStatus)
	}
	if len(respData.Records) < 1 {
		return nil, erplyerr("CalculateShoppingCart: no records in response", nil)
	}

	return respData.Records[0], nil
}
