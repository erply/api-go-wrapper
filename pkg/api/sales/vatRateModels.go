package sales

import (
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	VatRate struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Rate   string `json:"rate"`
		Code   string `json:"code"`
		Active string `json:"active"`
		//Added        string `json:"added"`
		LastModified string `json:"lastModified"`
		//IsReverseVat int    `json:"isReverseVat"`
		//ReverseRate int `json:"reverseRate"`
	}

	VatRates []VatRate

	NetTotalsByTaxRate struct {
		VatrateID int     `json:"vatrateID"`
		Total     float64 `json:"total"`
	}

	//GetVatRatesResponse ...
	getVatRatesResponse struct {
		Status   common2.Status `json:"status"`
		VatRates []VatRate      `json:"records"`
	}
)
