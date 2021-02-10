package pos

import (
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

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
		Status       common2.Status `json:"status"`
		PointsOfSale []PointOfSale  `json:"records"`
	}

	Clocking struct {
		InUnixTime        int64 `json:"InUnixTime"`
		OutUnixTime       int64 `json:"OutUnixTime"`
		EmployeeID        int64 `json:"employeeID"`
		TimeClockRecordID int64 `json:"timeclockRecordID"`
		WarehouseID       int64 `json:"warehouseID"`
	}

	GetClockInsResponse struct {
		Status   common2.Status `json:"status"`
		ClockIns []Clocking     `json:"records"`
	}
)
