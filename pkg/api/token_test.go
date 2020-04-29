package api

import (
	"context"
	"testing"
)

//works
func TestTokenRequests(t *testing.T) {
	const (
		//fill your data here
		sk         = ""
		cc         = ""
		jwt        = ""
		partnerKey = ""
	)
	var (
		ctx = context.Background()
	)
	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("test VerifyIdentityToken", func(t *testing.T) {
		resp, err := cli.VerifyIdentityToken(ctx, jwt)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp.SessionKey)
		if resp.SessionKey == "" {
			t.Error("got no session key")
			return
		}
	})
	t.Run("test GetIdentityToken", func(t *testing.T) {
		resp, err := cli.GetIdentityToken(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp.Jwt)
		if resp.Jwt == "" {
			t.Error("got no jwt key")
			return
		}
	})
	t.Run("test GetJWTToken", func(t *testing.T) {
		partnerCli, err := NewPartnerClient(sk, cc, partnerKey, nil)
		if err != nil {
			t.Error(err)
			return
		}
		resp, err := partnerCli.GetJWTToken(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
