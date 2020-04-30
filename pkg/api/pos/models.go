package pos

import "github.com/erply/api-go-wrapper/pkg/common"

type (
	PointOfSale struct {
		PointOfSaleID uint   `json:"pointOfSaleID"`
		Name          string `json:"name"`
		WarehouseID   int    `json:"warehouseID"`
		WarehouseName string `json:"warehouseName"`
		Added         uint64 `json:"added"`
		LastModified  uint64 `json:"lastModified"`
	}

	GetPointsOfSaleResponse struct {
		Status       common.Status `json:"status"`
		PointsOfSale []PointOfSale `json:"records"`
	}
)
