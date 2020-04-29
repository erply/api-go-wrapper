package api

import (
	"context"
	"testing"
)

//works
func TestErplyClient_GetConfParameters(t *testing.T) {
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
	resp, err := cli.GetConfParameters(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
