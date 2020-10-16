package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
	"testing"
)

func TestGetSalesReport(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test sales report", func(t *testing.T) {

		salesReport, err := cli.GetSalesReport(ctx, map[string]string{
			"type": "SALES_BY_PRODUCT",
			"dateStart": "2010-01-01",
			"dateEnd": "2020-12-30",
		})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(salesReport)
	})
}
