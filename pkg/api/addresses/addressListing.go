package addresses

import (
	"context"
)

type AddressListingDataProvider struct {
	erplyAPI Manager
}

func NewAddressListingDataProvider(erplyClient Manager) *AddressListingDataProvider {
	return &AddressListingDataProvider{
		erplyAPI: erplyClient,
	}
}

func (l *AddressListingDataProvider) Count(ctx context.Context, filters map[string]interface{}) (int, error) {
	filters["recordsOnPage"] = 1
	filters["pageNo"] = 1

	resp, err := l.erplyAPI.GetAddressesBulk(ctx, []map[string]interface{}{filters}, map[string]string{})

	if err != nil {
		return 0, err
	}

	if len(resp.BulkItems) == 0 {
		return 0, nil
	}

	return resp.BulkItems[0].Status.RecordsTotal, nil
}

func (l *AddressListingDataProvider) Read(ctx context.Context, bulkFilters []map[string]interface{}, callback func(item interface{})) error {
	resp, err := l.erplyAPI.GetAddressesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return err
	}

	for _, bulkItem := range resp.BulkItems {
		for _, address := range bulkItem.Addresses {
			callback(address)
		}
	}

	return nil
}
