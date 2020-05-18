package common

import (
	"net/url"
	"strings"
	"testing"
)

func TestClientCodeInsideClientsURL(t *testing.T) {
	authFunc := func(requestName string) url.Values {
		params := url.Values{}
		params.Add("request", requestName)
		params.Add(clientCode, "123")
		return params
	}
	t.Run("auth func case", func(t *testing.T) {
		c := NewClient("", "", "", nil, authFunc)
		if !strings.Contains(c.Url, "123") {
			t.Error(c.clientCode)
			return
		}
	})

	t.Run("client code provided case", func(t *testing.T) {
		c := NewClient("", "123", "", nil, nil)
		if !strings.Contains(c.Url, "123") {
			t.Error(c.clientCode)
			return
		}
	})
}
