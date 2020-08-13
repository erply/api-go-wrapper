package sales

import (
	"context"
)

type SaleDocumentsListingDataProvider struct {
	erplyAPI Manager
}

func NewSaleDocumentsListingDataProvider(erplyClient Manager) *SaleDocumentsListingDataProvider {
	return &SaleDocumentsListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (l *SaleDocumentsListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := l.erplyAPI.GetSalesDocumentsBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (l *SaleDocumentsListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := l.erplyAPI.GetSalesDocumentsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, doc := range bulkItem.SaleDocuments {
			callback(doc)
		}
	}

	return nil
}
