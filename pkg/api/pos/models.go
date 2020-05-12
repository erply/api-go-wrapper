package pos

import "github.com/tarmo-randma/api-go-wrapper/internal/common"

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
