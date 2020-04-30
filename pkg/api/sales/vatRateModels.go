package sales

import (
	"github.com/erply/api-go-wrapper/pkg/common"
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
		Status   common.Status `json:"status"`
		VatRates []VatRate     `json:"records"`
	}
)
