package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"time"
)

func main() {
	apiClient, err := api.BuildClient()
	common.Die(err)

	custmrs, err := GetCustomersBulk(apiClient)
	common.Die(err)

	fmt.Printf("GetCustomersBulk:\n%+v\n", custmrs)

	customers2, err := GetCustomersInParallel(apiClient)
	common.Die(err)

	fmt.Printf("GetCustomersInParallel:\n%+v\n", customers2)
}

func GetCustomersBulk(cl *api.Client) (custmrs customers.Customers, err error) {
	customerCli := cl.CustomerManager

	bulkFilters := []map[string]interface{}{
		{
			"recordsOnPage": "100",
			"pageNo":        "1",
		},
		{
			"recordsOnPage": "100",
			"pageNo":        "2",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := customerCli.GetCustomersBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, customr := range bulkItem.Customers {
			custmrs = append(custmrs, customr)
		}
	}

	return
}

func GetCustomersInParallel(cl *api.Client) (customers.Customers, error) {
	customersListingDataProvider := customers.NewCustomerListingDataProvider(cl.CustomerManager)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 5,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        300,
			MaxFetchersCount:          10,
		},
		customersListingDataProvider,
		func(sleepTime time.Duration) {
			time.Sleep(sleepTime)
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	customersChan := lister.Get(ctx, map[string]interface{}{
		"changedSince": time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC).Unix(),
	})

	customrs := make(customers.Customers, 0)
	for customer := range customersChan {
		if customer.Err != nil {
			return customrs, customer.Err
		}
		customrs = append(customrs, customer.Payload.(customers.Customer))
	}

	return customrs, nil
}
