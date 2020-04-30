package sales

import (
	"context"
	"testing"
)

//works
func TestSalesDocuments(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	ctx := context.Background()
	cli := NewClient(sk, cc, "", nil)
	t.Run("test get sales doc", func(t *testing.T) {
		paymentID, err := cli.GetSalesDocuments(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(paymentID)
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
