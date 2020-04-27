package api

import "testing"

func TestGetServiceEndpoints(t *testing.T) {
	const (
		cc = ""
	)
	cli := NewClient("", cc, nil)

	endpoints, err := cli.GetServiceEndpoints()
	if err != nil {
		t.Error(err)
		return
	}

	if endpoints.Cafa.Url == "" {
		t.Error(err)
		return
	}

}
