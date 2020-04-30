package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"testing"
)

//works
func TestPaymentManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""

		documentID   = ""
		paymentType  = ""
		currencyCode = ""
	)

	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil))
	t.Run("test save payment", func(t *testing.T) {
		params := map[string]string{
			"documentID":   documentID,
			"type":         paymentType,
			"currencyCode": currencyCode,
		}
		paymentID, err := cli.SavePayment(ctx, params)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(paymentID)
	})
	t.Run("test get payments", func(t *testing.T) {
		resp, err := cli.GetPayments(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
