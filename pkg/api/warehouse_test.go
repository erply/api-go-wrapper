package api

import "testing"

//works
func TestErplyClient_GetWarehouses(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(sk, cc, nil)
	resp, err := cli.GetWarehouses()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
