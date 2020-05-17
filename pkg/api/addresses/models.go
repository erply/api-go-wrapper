package addresses

import (
	"github.com/breathbath/api-go-wrapper/internal/common"
	common2 "github.com/breathbath/api-go-wrapper/pkg/api/common"
)

type (
	//GetAddressesResponse ..
	Response struct {
		Status    common2.Status   `json:"status"`
		Addresses common.Addresses `json:"records"`
	}
)
