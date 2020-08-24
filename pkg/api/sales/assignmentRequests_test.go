package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/internal/common"
	"testing"
)

//works
func TestSaveAssignment(t *testing.T) {
	const (
		//fill your data here
		sk        = ""
		cc        = ""
		vatRateID = ""
	)
	var (
		ctx = context.Background()
	)
	cli := NewClient(common.NewClient(sk, cc, "", nil, nil))

	resp, err := cli.SaveAssignment(ctx, map[string]string{
		"customerComment1":  "Test",
	})

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
