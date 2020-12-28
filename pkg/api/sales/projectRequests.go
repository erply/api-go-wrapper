package sales

import (
	"context"
	"encoding/json"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

// GetProjects will list projects according to specified filters.
func (cli *Client) GetProjects(ctx context.Context, filters map[string]string) ([]Project, error) {
	resp, err := cli.SendRequest(ctx, "getProjects", filters)
	if err != nil {
		return nil, err
	}
	var res GetProjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, sharedCommon.NewFromError("failed to unmarshal GetProjectsResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
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
		return nil, sharedCommon.NewFromError("failed to unmarshal GetProjectStatusesResponse", err, 0)
	}
	if !common.IsJSONResponseOK(&res.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&res.Status)
	}
	return res.ProjectStatuses, nil
}
