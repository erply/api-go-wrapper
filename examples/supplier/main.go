package main

import (
	"context"
	"fmt"
	sharedCommon "github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"time"
)

func main() {
	apiClient, err := api.BuildClient()
	sharedCommon.Die(err)

	suppliers, err := GetSupplierBulk(apiClient)
	sharedCommon.Die(err)

	fmt.Println(suppliers)

	err = SaveSupplierBulk(apiClient)
	sharedCommon.Die(err)

	err = DeleteSupplierBulk(apiClient)
	sharedCommon.Die(err)
}

func GetSupplierBulk(cl *api.Client) (suppliers []customers.Supplier, err error) {
	supplierCli := cl.CustomerManager

	bulkFilters := []map[string]interface{}{
		{
			"recordsOnPage": 2,
			"pageNo":        1,
		},
		{
			"recordsOnPage": 2,
			"pageNo":        2,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := supplierCli.GetSuppliersBulk(ctx, bulkFilters, map[string]string{})
	if err != nil {
		return
	}

	for _, bulkItem := range bulkResp.BulkItems {
		for _, supplier := range bulkItem.Suppliers {
			suppliers = append(suppliers, supplier)
		}
	}

	return
}

func SaveSupplierBulk(cl *api.Client) (err error) {
	supplierCli := cl.CustomerManager

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	sup := []map[string]interface{}{
		{
			"phone":      "+1919-820-1136",
			"fullname":   "Max Mustermann",
			"supplierID": 12355,
		},
	}
	bulkResponse, err := supplierCli.SaveSupplierBulk(ctx, sup, map[string]string{})
	if err != nil {
		return
	}

	fmt.Printf("%+v", bulkResponse)

	return
}

func DeleteSupplierBulk(cl *api.Client) (err error) {
	supplierCli := cl.CustomerManager

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ids := []map[string]interface{}{
		{
			"supplierID": "100000049",
		},
		{
			"supplierID": "100000050",
		},
	}
	bulkResponse, err := supplierCli.DeleteSupplierBulk(ctx, ids, map[string]string{})
	if err != nil {
		return
	}

	fmt.Printf("%+v", bulkResponse)

	return
}
