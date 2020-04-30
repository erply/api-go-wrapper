package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

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

	ProjectManager interface {
		GetProjects(ctx context.Context, filters map[string]string) ([]Project, error)
		GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error)
	}
)

// GetProjects will list projects according to specified filters.
func (cli *Client) GetProjects(ctx context.Context, filters map[string]string) ([]Project, error) {
	resp, err := cli.SendRequest(ctx, "getProjects", filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetProjectsResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Projects, nil
}

// GetProjectStatus will list projects statuses according to specified filters.
func (cli *Client) GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error) {
	resp, err := cli.SendRequest(ctx, "getProjectStatuses", filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectStatusesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erro.NewFromError("failed to unmarshal GetProjectStatusesResponse", err)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProjectStatuses, nil
}
