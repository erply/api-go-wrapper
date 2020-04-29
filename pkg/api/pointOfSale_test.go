package api

import (
	"context"
	"strconv"
	"testing"
)

//works
func TestPOSRequests(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)
	cli := NewClient(sk, cc, nil)
	t.Run("test by ID", func(t *testing.T) {
		resp, err := cli.GetPointsOfSaleByID("1")
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp.WarehouseID)
		if resp.WarehouseID == 0 {
			t.Error("got no warehouseID key")
			return
		}

		t.Run("test general request", func(t *testing.T) {
			ctx := context.Background()
			resp2, err := cli.GetPointsOfSale(ctx, map[string]string{
				"warehouseID": strconv.Itoa(resp.WarehouseID),
			})
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(resp2)
		})
	})
}
