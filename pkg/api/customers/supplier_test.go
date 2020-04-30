package customers

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/api"
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

	cli := NewClient(sk, cc, "", nil)
	t.Run("test get suppliers", func(t *testing.T) {
		suppliers, err := cli.GetSuppliers(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(suppliers)
	})
	t.Run("test post supplier", func(t *testing.T) {
		params := map[string]string{
			"companyName": testingCustomer.CompanyName,
			"code":        testingCustomer.RegistryCode,
		}
		resp, err := cli.PostSupplier(ctx, params)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
}
