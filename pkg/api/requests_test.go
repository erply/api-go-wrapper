package api

import (
	"testing"
)

//here the "works" indicator will be under each test separately
func TestApiRequests(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(sk, cc, nil)
	//works
	t.Run("test getUsername", func(t *testing.T) {
		resp, err := cli.GetUserName()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
