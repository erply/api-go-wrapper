package auth

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
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
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))

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
		partnerCli := NewClient(common.NewClient(sk, cc, partnerKey, nil, nil))

		resp, err := partnerCli.GetJWTToken(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
