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

func (sdldp *SaleDocumentsListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := sdldp.erplyAPI.GetSalesDocumentsBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (sdldp *SaleDocumentsListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := sdldp.erplyAPI.GetSalesDocumentsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for i := range bulkItem.SaleDocuments {
			callback(bulkItem.SaleDocuments[i])
		}
	}

	return nil
}

type VatRatesListingDataProvider struct {
	erplyAPI Manager
}

func NewVatRatesListingDataProvider(erplyClient Manager) *VatRatesListingDataProvider {
	return &VatRatesListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (vrldp *VatRatesListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := vrldp.erplyAPI.GetVatRatesBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (vrldp *VatRatesListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := vrldp.erplyAPI.GetVatRatesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, doc := range bulkItem.VatRates {
			callback(doc)
		}
	}

	return nil
}

type PaymentsListingDataProvider struct {
	erplyAPI Manager
}

func NewPaymentsListingDataProvider(erplyClient Manager) *PaymentsListingDataProvider {
	return &PaymentsListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (sdldp *PaymentsListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := sdldp.erplyAPI.GetPaymentsBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (sdldp *PaymentsListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := sdldp.erplyAPI.GetPaymentsBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for i := range bulkItem.PaymentInfos {
			callback(bulkItem.PaymentInfos[i])
		}
	}

	return nil
}