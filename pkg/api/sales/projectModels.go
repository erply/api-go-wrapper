package sales

import "github.com/erply/api-go-wrapper/pkg/common"

type (
	GetProjectsResponse struct {
		Status   common.Status `json:"status"`
		Projects []Project     `json:"records"`
	}

	GetProjectStatusesResponse struct {
		Status          common.Status   `json:"status"`
		ProjectStatuses []ProjectStatus `json:"records"`
	}

	Project struct {
		ProjectID    uint   `json:"projectID"`
		Name         string `json:"name"`
		CustomerID   uint   `json:"customerID"`
		CustomerName string `json:"customerName"`
		EmployeeID   uint   `json:"employeeID"`
		EmployeeName string `json:"employeeName"`
		TypeID       uint   `json:"typeID"`
		TypeName     string `json:"typeName"`
		StatusID     uint   `json:"statusID"`
		StatusName   string `json:"statusName"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		Notes        string `json:"notes"`
		LastModified uint64 `json:"lastModified"`
	}

	ProjectStatus struct {
		ProjectStatusID uint   `json:"projectStatusID"`
		Name            string `json:"name"`
		Finished        byte   `json:"finished"`
		Added           uint64 `json:"added"`
		LastModified    uint64 `json:"lastModified"`
	}
)
