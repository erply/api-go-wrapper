package api

import "testing"

//works
func TestTokenRequests(t *testing.T) {
	const (
		//fill your data here
		sk         = ""
		cc         = ""
		jwt        = ""
		partnerKey = ""
	)
	cli := NewClient(sk, cc, nil)
	t.Run("test VerifyIdentityToken", func(t *testing.T) {
		specCli := NewClient("", cc, nil)
		resp, err := specCli.VerifyIdentityToken(jwt)
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
		resp, err := cli.GetIdentityToken()
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
		resp, err := cli.GetJWTToken(partnerKey)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
