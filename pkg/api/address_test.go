package api

import (
	"context"
	"testing"
)

//works
func TestAddressManager(t *testing.T) {
	const (
		//fill your data here
		sk      = ""
		cc      = ""
		ownerID = ""
	)
	ctx := context.Background()
	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := cli.GetAddresses(ctx, map[string]string{
		"ownerID": ownerID,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
	t.Run("test save address", func(t *testing.T) {
		req := &AddressRequest{
			OwnerID: 0, //put your value here
			TypeID:  0, //put your value here
		}
		resp, err := cli.SaveAddress(ctx, req)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
