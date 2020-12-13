package products

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

	resp, err := l.erplyAPI.GetProductsBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (l *ListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := l.erplyAPI.GetProductsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, prod := range bulkItem.Products {
			callback(prod)
		}
	}

	return nil
}


type PrioGroupListingDataProvider struct {
	erplyAPI Manager
}

func NewPrioGroupListingDataProvider(erplyClient Manager) *PrioGroupListingDataProvider {
	return &PrioGroupListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (pgldp *PrioGroupListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := pgldp.erplyAPI.GetProductPriorityGroupBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (pgldp *PrioGroupListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := pgldp.erplyAPI.GetProductPriorityGroupBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, prodPrioGroup := range bulkItem.Records {
			callback(prodPrioGroup)
		}
	}

	return nil
}
