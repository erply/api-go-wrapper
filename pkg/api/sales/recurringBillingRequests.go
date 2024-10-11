package sales

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/pkg/errors"
	"net/http"
)

// ProcessRecurringBilling is Erply API call processRecurringBilling. https://learn-api.erply.com/requests/processrecurringbilling
func (cli *Client) ProcessRecurringBilling(ctx context.Context, filters map[string]string) ([]RecurringBillingProcessedInvoices, error) {
	resp, err := cli.SendRequest(ctx, "processRecurringBilling", filters)
	if err != nil {
		return nil, sharedCommon.NewFromError("ProcessRecurringBilling: error sending request", err, 0)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, sharedCommon.NewFromError(fmt.Sprintf("ProcessRecurringBilling: bad response status code: %d", resp.StatusCode), nil, 0)
	}

	respData := RecurringBillingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, errors.Wrap(err, "failed to decode processRecurringBilling response")
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, sharedCommon.NewFromResponseStatus(&respData.Status)
	}
	if len(respData.Records) > 0 {
		return respData.Records[0].ProcessedInvoices, nil
	} else {
		return make([]RecurringBillingProcessedInvoices, 0), nil
	}
}
