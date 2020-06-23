package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
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
