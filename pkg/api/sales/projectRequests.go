package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/pkg/common"
	erro "github.com/erply/api-go-wrapper/pkg/errors"
	"strconv"
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
