package sales

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
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
	GetVatRatesResponse struct {
		Status   sharedCommon.Status `json:"status"`
		VatRates []VatRate           `json:"records"`
	}

	GetVatRatesBulkItem struct {
		Status   sharedCommon.StatusBulk `json:"status"`
		VatRates []VatRate               `json:"records"`
	}

	GetVatRatesResponseBulk struct {
		Status    sharedCommon.Status   `json:"status"`
		BulkItems []GetVatRatesBulkItem `json:"requests"`
	}

	SaveVatRateResult struct {
		VatRateID int `json:"vatRateID"`
	}

	SaveVatRateResultResponse struct {
		Status            sharedCommon.Status `json:"status"`
		SaveVatRateResult []SaveVatRateResult `json:"records"`
	}

	SaveVatRateBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SaveVatRateResult     `json:"records"`
	}
	SaveVatRateResponseBulk struct {
		Status    sharedCommon.Status   `json:"status"`
		BulkItems []SaveVatRateBulkItem `json:"requests"`
	}

	SaveVatRateComponentResult struct {
		VatRateComponentID int `json:"vatRateComponentID"`
	}

	SaveVatRateComponentResultResponse struct {
		Status                     sharedCommon.Status          `json:"status"`
		SaveVatRateComponentResult []SaveVatRateComponentResult `json:"records"`
	}

	SaveVatRateComponentBulkItem struct {
		Status  sharedCommon.StatusBulk      `json:"status"`
		Records []SaveVatRateComponentResult `json:"records"`
	}
	SaveVatRateComponentResponseBulk struct {
		Status    sharedCommon.Status            `json:"status"`
		BulkItems []SaveVatRateComponentBulkItem `json:"requests"`
	}
)
