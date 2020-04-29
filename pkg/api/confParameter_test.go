package api

import "testing"

//works
func TestErplyClient_GetConfParameters(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(sk, cc, nil)
	resp, err := cli.GetConfParameters()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
