package sales

import (
	"context"
	"testing"

	"github.com/tarmo-randma/api-go-wrapper/internal/common"
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
