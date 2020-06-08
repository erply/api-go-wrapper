package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
	"testing"
)

//works
func TestSalesDocuments(t *testing.T) {
	const (
		//fill your data here
		sk              = ""
		cc              = ""
		invoiceNoToSave = ""
		supplierID      = ""
		vatrateID       = ""
		amount          = ""
		price           = ""
	)

	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test get sales doc", func(t *testing.T) {
		saleDocs, err := cli.GetSalesDocuments(ctx, map[string]string{
			"id": "",
		})
		if err != nil {
			t.Error(err)
			return
		}

		for _, r := range saleDocs[0].InvoiceRows {
			t.Logf("row's code2: %s", r.Code2)
			t.Logf(r.StableRowID)
		}
	})

	t.Run("test save purchase", func(t *testing.T) {
		resp, err := cli.SavePurchaseDocument(ctx, map[string]string{
			"currencyCode": "EUR",
			"no":           invoiceNoToSave,
			"supplierID":   supplierID,
			"vatrateID":    vatrateID,
			"amount":       amount,
			"price":        price,
		})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	t.Run("test save sales doc", func(t *testing.T) {
		reports, err := cli.SaveSalesDocument(ctx, map[string]string{
			"id":         "57",
			"productID1": "4",
			"amount1":    "2",
			"price1":     "20",
		})
		if err != nil {
			t.Error(err)
			return
		}
		for _, r := range reports {
			t.Log(r.InvoiceID)
		}
	})
}
