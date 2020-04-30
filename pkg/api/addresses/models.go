package addresses

import "github.com/erply/api-go-wrapper/pkg/common"

type (
	//GetAddressesResponse ..
	Response struct {
		Status    common.Status    `json:"status"`
		Addresses common.Addresses `json:"records"`
	}
)
