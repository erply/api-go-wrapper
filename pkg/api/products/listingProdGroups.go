package products

import (
	"context"
)

type ProductGroupsListingDataProvider struct {
	erplyAPI Manager
}

func NewProductGroupsListingDataProvider(erplyClient Manager) *ProductGroupsListingDataProvider {
	return &ProductGroupsListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (pgldp *ProductGroupsListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := pgldp.erplyAPI.GetProductGroupsBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (pgldp *ProductGroupsListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := pgldp.erplyAPI.GetProductGroupsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, productGroup := range bulkItem.Records {
			callback(productGroup)
		}
	}

	return nil
}
