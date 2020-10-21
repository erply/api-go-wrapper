package addresses

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	//GetAddressesResponse ..
	Response struct {
		Status    sharedCommon.Status    `json:"status"`
		Addresses sharedCommon.Addresses `json:"records"`
	}

	GetAddressesResponseBulkItem struct {
		Status    sharedCommon.StatusBulk `json:"status"`
		Addresses sharedCommon.Addresses  `json:"records"`
	}

	GetAddressesResponseBulk struct {
		Status    sharedCommon.Status            `json:"status"`
		BulkItems []GetAddressesResponseBulkItem `json:"requests"`
	}

	SaveAddressResp struct {
		AddressID int `json:"addressID"`
	}

	SaveAddressesResponseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SaveAddressResp       `json:"records"`
	}

	SaveAddressesResponseBulk struct {
		Status    sharedCommon.Status             `json:"status"`
		BulkItems []SaveAddressesResponseBulkItem `json:"requests"`
	}

	DeleteAddressResponse struct {
		Status               sharedCommon.Status   `json:"status"`
	}

	DeleteAddressBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
	}

	DeleteAddressResponseBulk struct {
		Status    sharedCommon.Status     `json:"status"`
		BulkItems []DeleteAddressBulkItem `json:"requests"`
	}
)
