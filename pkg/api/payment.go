package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	PaymentAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeType  string `json:"attributeType"`
		AttributeValue string `json:"attributeValue"`
	}
	PaymentStatus string
	PaymentType   string

	PaymentInfo struct {
		DocumentID   int    `json:"documentID"` // Invoice ID
		Type         string `json:"type"`       // CASH, TRANSFER, CARD, CREDIT, GIFTCARD, CHECK, TIP
		Date         string `json:"date"`
		Sum          string `json:"sum"`
		CurrencyCode string `json:"currencyCode"` // EUR, USD
		Info         string `json:"info"`         // Information about the payer or payment transaction
		Added        uint64 `json:"added"`
	}
	PaymentManager interface {
		SavePayment(ctx context.Context, filters map[string]string) (int64, error)
		GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error)
	}
)

func (cli *erplyClient) SavePayment(ctx context.Context, filters map[string]string) (int64, error) {

	/*	params.Add("documentID", strconv.Itoa(in.DocumentID))
		params.Add("type", in.Type)
		params.Add("currencyCode", in.CurrencyCode)
		params.Add("date", in.Date)
		params.Add("sum", in.Sum)
		params.Add("info", in.Info)
	*/

	resp, err := cli.sendRequest(ctx, savePaymentMethod, filters)
	if err != nil {
		return 0, erplyerr("SavePayment: error sending POST request", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return 0, erplyerr(fmt.Sprintf("SavePayment: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []struct {
			PaymentID int64 `json:"paymentID"`
		} `json:"records"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return 0, erplyerr("SavePayment: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return 0, erplyerr(fmt.Sprintf("SavePayment: API error %d", respData.Status.ErrorCode), nil)
	}
	if len(respData.Records) < 1 {
		return 0, erplyerr("SavePayment: no records in response", nil)
	}

	return respData.Records[0].PaymentID, nil
}

func (cli *erplyClient) GetPayments(ctx context.Context, filters map[string]string) ([]PaymentInfo, error) {
	resp, err := cli.sendRequest(ctx, GetPaymentsMethod, filters)
	if err != nil {
		return nil, erplyerr("GetPayments: error sending request", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, erplyerr(fmt.Sprintf("GetPayments: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  Status
		Records []PaymentInfo
	}

	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, erplyerr("GetPayments: error decoding JSON response body", err)
	}
	if respData.Status.ErrorCode != 0 {
		return nil, erplyerr(fmt.Sprintf("GetPayments: API error %d", respData.Status.ErrorCode), nil)
	}

	return respData.Records, nil
}
