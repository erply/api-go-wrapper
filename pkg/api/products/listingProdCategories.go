package products

import (
	"context"
)

type ProductCategoriesListingDataProvider struct {
	erplyAPI Manager
}

func NewProductCategoriesListingDataProvider(erplyClient Manager) *ProductCategoriesListingDataProvider {
	return &ProductCategoriesListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (pcldp *ProductCategoriesListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := pcldp.erplyAPI.GetProductCategoriesBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (pcldp *ProductCategoriesListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := pcldp.erplyAPI.GetProductCategoriesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, prodCat := range bulkItem.Records {
			callback(prodCat)
		}
	}

	return nil
}
