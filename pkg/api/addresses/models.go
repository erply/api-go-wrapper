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
)
