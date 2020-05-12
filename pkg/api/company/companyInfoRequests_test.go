package company

import (
	"context"
	"testing"

	"github.com/tarmo-randma/api-go-wrapper/internal/common"
)

//works
func TestErplyClient_GetCompanyInfo(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil))
	resp, err := cli.GetCompanyInfo(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
