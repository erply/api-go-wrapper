package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
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
	cli := NewClient(common.NewClient(sk, cc, "", nil))
	t.Run("test get sales doc", func(t *testing.T) {
		paymentID, err := cli.GetSalesDocuments(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(paymentID)
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
		resp, err := cli.SaveSalesDocument(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
