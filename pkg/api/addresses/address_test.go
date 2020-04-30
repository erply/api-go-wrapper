package addresses

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
	cli := NewClient(common.NewClient(sk, cc, "", nil))
	resp, err := cli.GetAddresses(ctx, map[string]string{
		"ownerID": ownerID,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
	t.Run("test save address", func(t *testing.T) {
		filters := map[string]string{
			"ownerID": "", //put your value here
			"typeID":  "", //put your value here
		}
		resp, err := cli.SaveAddress(ctx, filters)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
