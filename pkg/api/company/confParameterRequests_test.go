package company

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
	"testing"
)

//works
func TestErplyClient_GetConfParameters(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)
	ctx := context.Background()
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))
	t.Run("test conf parameters", func(t *testing.T) {
		resp, err := cli.GetConfParameters(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})

	t.Run("test GetDefaultLanguage", func(t *testing.T) {
		lang, err := cli.GetDefaultLanguage(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(lang)
	})
}
