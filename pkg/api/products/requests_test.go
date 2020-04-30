package products

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/common"
	"testing"
)

//works
func TestProductManager(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)
	//and here
	var (
		ctx = context.Background()
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil))

	t.Run("test GetProducts", func(t *testing.T) {
		products, err := cli.GetProducts(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(products)
	})
	t.Run("test get product units", func(t *testing.T) {
		units, err := cli.GetProductUnits(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(units)
	})

	t.Run("test get product categories", func(t *testing.T) {
		cats, err := cli.GetProductCategories(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(cats)
	})
	t.Run("test get product brands", func(t *testing.T) {
		brands, err := cli.GetProductCategories(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(brands)
	})
	t.Run("test get product groups", func(t *testing.T) {
		groups, err := cli.GetProductGroups(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(groups)
	})
}
