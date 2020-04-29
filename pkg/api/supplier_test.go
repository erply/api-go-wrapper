package api

import (
	"context"
	"testing"
)

func TestSupplierManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)
	//and here
	var (
		testingCustomer = &CustomerRequest{
			RegistryCode: "",
			CompanyName:  "",
			Username:     "",
			Password:     "",
		}
		ctx = context.Background()
	)

	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("test get suppliers", func(t *testing.T) {
		suppliers, err := cli.GetSuppliers(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(suppliers)
	})
	t.Run("test post supplier", func(t *testing.T) {
		resp, err := cli.PostSupplier(ctx, testingCustomer, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
