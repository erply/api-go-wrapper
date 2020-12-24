package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"net/http"
)

func (cli *Client) GetSalesReport(ctx context.Context, filters map[string]string) (*GetSalesReport, error) {
	var salesReportResp GetSalesReport
	resp, err := cli.SendRequest(ctx, "getSalesReport", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("getSalesReport: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("getSalesReport: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	if err := json.NewDecoder(resp.Body).Decode(&salesReportResp); err != nil {
		return nil, sharedCommon.NewFromError("getSalesReport: unmarshaling response failed", err, 0)
	}
	if !common.IsJSONResponseOK(&salesReportResp.Status) {
		return &salesReportResp, sharedCommon.NewErplyError(
			salesReportResp.Status.ErrorCode.String(),
			salesReportResp.Status.Request+": "+salesReportResp.Status.ResponseStatus,
			salesReportResp.Status.ErrorCode,
		)
	}
	if len(salesReportResp.Records) < 1 {
		return &salesReportResp, sharedCommon.NewFromError("getSalesReport: no records in response", nil, salesReportResp.Status.ErrorCode)
	}

	return &salesReportResp, nil
}
