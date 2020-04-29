package api

import (
	"context"
	"encoding/json"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
)

type GetProjectsResponse struct {
	Status   Status    `json:"status"`
	Projects []Project `json:"records"`
}

type GetProjectStatusesResponse struct {
	Status          Status          `json:"status"`
	ProjectStatuses []ProjectStatus `json:"records"`
}

type Project struct {
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

type ProjectStatus struct {
	ProjectStatusID uint   `json:"projectStatusID"`
	Name            string `json:"name"`
	Finished        byte   `json:"finished"`
	Added           uint64 `json:"added"`
	LastModified    uint64 `json:"lastModified"`
}

// GetProjects will list projects according to specified filters.
func (cli *erplyClient) GetProjects(ctx context.Context, filters map[string]string) ([]Project, error) {
	resp, err := cli.sendRequest(ctx, GetProjectsMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetProjectsResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.Projects, nil
}

// GetProjectStatus will list projects statuses according to specified filters.
func (cli *erplyClient) GetProjectStatus(ctx context.Context, filters map[string]string) ([]ProjectStatus, error) {
	resp, err := cli.sendRequest(ctx, GetProjectStatusesMethod, filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectStatusesResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, erplyerr("failed to unmarshal GetProjectStatusesResponse", err)
	}
	if !isJSONResponseOK(&res.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(res.Status.ErrorCode), res.Status.Request+": "+res.Status.ResponseStatus)
	}
	return res.ProjectStatuses, nil
}
