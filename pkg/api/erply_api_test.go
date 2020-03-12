package api

import (
	"testing"
)

func TestApiRequests(t *testing.T) {
	const (
		//fill your data here
		sk  = ""
		cc  = ""
		jwt = ""
	)
	cli := NewClient(sk, cc, nil)
	t.Run("test VerifyIdentityToken", func(t *testing.T) {
		specCli := NewClient("", cc, nil)
		resp, err := specCli.VerifyIdentityToken(jwt)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp.SessionKey)
		if resp.SessionKey == "" {
			t.Error("got no session key")
			return
		}
	})
	t.Run("test GetIdentityToken", func(t *testing.T) {
		resp, err := cli.GetIdentityToken()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp.Jwt)
		if resp.Jwt == "" {
			t.Error("got no jwt key")
			return
		}
	})
	t.Run("test GetPOSes", func(t *testing.T) {
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
	})
}
