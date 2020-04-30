package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"testing"
)

//works
func TestShoppingCart(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil))
	t.Run("test shopping cart", func(t *testing.T) {

		paymentID, err := cli.CalculateShoppingCart(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(paymentID)
	})
}
