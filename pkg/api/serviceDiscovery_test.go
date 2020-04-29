package api

import (
	"context"
	"testing"
)

//works
func TestGetServiceEndpoints(t *testing.T) {
	const (
		sk = ""
		cc = ""
	)
	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}

	endpoints, err := cli.GetServiceEndpoints(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	if endpoints.Cafa.Url == "" {
		t.Error(err)
		return
	}
	t.Log(endpoints)
}
