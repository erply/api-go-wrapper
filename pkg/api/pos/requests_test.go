package pos

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"testing"
)

//works
func TestPOSRequests(t *testing.T) {
	const (
		//fill your data here
		sk          = ""
		cc          = ""
		warehouseID = ""
	)
	cli := NewClient(common.NewClient(sk, cc, "", nil))
	t.Run("test by ID", func(t *testing.T) {

		t.Run("test general request", func(t *testing.T) {
			ctx := context.Background()
			resp2, err := cli.GetPointsOfSale(ctx, map[string]string{
				"warehouseID": warehouseID,
			})
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(resp2)
		})
	})
}
