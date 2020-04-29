package api

import "testing"

//works
func TestErplyClient_GetVatRatesByID(t *testing.T) {
	const (
		//fill your data here
		sk        = ""
		cc        = ""
		vatRateID = ""
	)

	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := cli.GetVatRatesByID(vatRateID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
