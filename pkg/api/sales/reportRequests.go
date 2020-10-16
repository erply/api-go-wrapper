package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	erro "github.com/erply/api-go-wrapper/internal/errors"
	"net/http"
)

func (cli *Client) GetSalesReport(ctx context.Context, filters map[string]string)(*GetSalesReport, error) {
	var salesReportResp GetSalesReport
	resp, err := cli.SendRequest(ctx, "getSalesReport", filters)
	if err != nil {
		return nil, erro.NewFromError("getSalesReport: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, erro.NewFromError(fmt.Sprintf("getSalesReport: bad response status code: %d", resp.StatusCode), nil)
	}

	if err := json.NewDecoder(resp.Body).Decode(&salesReportResp); err != nil {
		return nil, erro.NewFromError("getSalesReport: unmarshaling response failed", err)
	}
	if !common.IsJSONResponseOK(&salesReportResp.Status) {
		return &salesReportResp, erro.NewErplyError(salesReportResp.Status.ErrorCode.String(), salesReportResp.Status.Request+": "+salesReportResp.Status.ResponseStatus)
	}
	if len(salesReportResp.Records) < 1 {
		return &salesReportResp, erro.NewFromError("getSalesReport: no records in response", nil)
	}

	return &salesReportResp, nil
}
