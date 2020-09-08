package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/erply/api-go-wrapper/pkg/api/warehouse"
	"net/http"
	"time"
)

func main() {
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	clientCode := flag.String("cc", "", "client code")
	flag.Parse()

	sessionKey, err := auth.VerifyUser(*username, *password, *clientCode, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	apiClient, err := api.NewClient(sessionKey, *clientCode, nil)
	if err != nil {
		panic(err)
	}

	warehouses, err := GetWarehousesBulk(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetWarehousesBulk: %+v\n", warehouses)

	warehouseFromParallel, err := GetWarehousesInParallel(apiClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("GetWarehousesInParallel: %+v\n", warehouseFromParallel)
}

func GetWarehousesBulk(cl *api.Client) (warehouses warehouse.Warehouses, err error) {
	warehouseManager := cl.WarehouseManager

	bulkFilters := []map[string]interface{}{
		{
			"warehouseID": 1,
		},
		{
			"warehouseID": 2,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := warehouseManager.GetWarehousesBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, warehouseItem := range bulkItem.Warehouses {
			warehouses = append(warehouses, warehouseItem)
		}
	}

	return
}

func GetWarehousesInParallel(cl *api.Client) (warehouse.Warehouses, error) {
	listingDataProvider := warehouse.NewListingDataProvider(cl.WarehouseManager)

	lister := sharedCommon.NewLister(
		sharedCommon.ListingSettings{
			MaxRequestsCountPerSecond: 5,
			StreamBufferLength:        10,
			MaxItemsPerRequest:        30,
			MaxFetchersCount:          2,
		},
		listingDataProvider,
		func(sleepTime time.Duration) {
			time.Sleep(sleepTime)
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	warehousesChan := lister.Get(ctx, map[string]interface{}{
		"code": "108",
	})

	warehouses := make(warehouse.Warehouses, 0)
	for wrs := range warehousesChan {
		if wrs.Err != nil {
			return warehouses, wrs.Err
		}
		warehouses = append(warehouses, wrs.Payload.(warehouse.Warehouse))
	}

	return warehouses, nil
}
