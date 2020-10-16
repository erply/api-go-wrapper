package sales

import common2 "github.com/erply/api-go-wrapper/pkg/api/common"

type GetSalesReport struct {
	Status  common2.Status `json:"status"`
	Records []struct {
		ReportLink string `json:"reportLink"`
	}
}
