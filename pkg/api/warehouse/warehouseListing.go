package warehouse

import (
	"context"
)

type ListingDataProvider struct {
	erplyAPI Manager
}

func NewListingDataProvider(erplyClient Manager) *ListingDataProvider {
	return &ListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (l *ListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := l.erplyAPI.GetWarehousesBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (l *ListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := l.erplyAPI.GetWarehousesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, warehouse := range bulkItem.Warehouses {
			callback(warehouse)
		}
	}

	return nil
}
