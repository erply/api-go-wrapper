package addresses

import (
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	//GetAddressesResponse ..
	Response struct {
		Status    common2.Status    `json:"status"`
		Addresses common2.Addresses `json:"records"`
	}
)
