package api

import (
	"context"
	"testing"
)

//works
func TestErplyClient_GetWarehouses(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := cli.GetWarehouses(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
