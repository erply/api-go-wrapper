package sales

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"testing"
)

func TestProjectManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil))
	//test this
	t.Run("test GetProjects", func(t *testing.T) {
		resp, err := cli.GetProjects(context.Background(), map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//test this
	t.Run("test GetProjectStatus", func(t *testing.T) {
		resp, err := cli.GetProjectStatus(context.Background(), map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
