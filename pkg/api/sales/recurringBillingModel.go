package sales

import common2 "github.com/erply/api-go-wrapper/pkg/api/common"

type (
	RecurringBillingResponse struct {
		Status  common2.Status
		Records []RecurringBillingRecord `json:"records"`
	}
	RecurringBillingRecord struct {
		ProcessedInvoices []RecurringBillingProcessedInvoices `json:"processedInvoices"`
	}
	RecurringBillingProcessedInvoices struct {
		ID      int  `json:"id"`
		Created bool `json:"created"`
		Updated bool `json:"updated"`
	}
)
