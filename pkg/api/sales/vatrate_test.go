package sales

import (
	"context"
	"testing"
)

//works
func TestVatRateManager(t *testing.T) {
	const (
		//fill your data here
		sk        = ""
		cc        = ""
		vatRateID = ""
	)
	var (
		ctx = context.Background()
	)

	cli := NewClient(sk, cc, "", nil)

	resp, err := cli.GetVatRates(ctx, map[string]string{
		"searchAttributeName":  "id",
		"searchAttributeValue": vatRateID,
		"active":               "1",
	})

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
