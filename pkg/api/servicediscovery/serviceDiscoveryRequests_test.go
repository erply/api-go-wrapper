package servicediscovery

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"testing"
)

//works
func TestGetServiceEndpoints(t *testing.T) {
	const (
		sk = ""
		cc = ""
	)
	cli := NewClient(common.NewClient(sk, cc, "", nil))

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
