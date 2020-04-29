package api

import (
	"context"
	"testing"
)

//here the "works" indicator will be under each test separately
func TestApiRequests(t *testing.T) {
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
	//works
	t.Run("test getUsername", func(t *testing.T) {
		resp, err := cli.GetUserRights(context.Background(), map[string]string{
			"getCurrentUser": "1",
		})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
