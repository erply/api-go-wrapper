package addresses

import (
	"github.com/erply/api-go-wrapper/internal/common"
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	//GetAddressesResponse ..
	Response struct {
		Status    common2.Status   `json:"status"`
		Addresses common.Addresses `json:"records"`
	}
)
