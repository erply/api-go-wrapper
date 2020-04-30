package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
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
	cli := NewClient(common.NewClient(sk, cc, "", nil))

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
