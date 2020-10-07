package main

import (
	"context"
	"fmt"
	"github.com/erply/api-go-wrapper/internal/common"
	"github.com/erply/api-go-wrapper/pkg/api"
	"time"
)

func main() {
	apiClient, err := api.BuildClient()
	common.Die(err)

	SaveVatRate(apiClient)
	GetVatRate(apiClient)
	SaveVatRateBulk(apiClient)
	SaveVatRateComponent(apiClient)
	SaveVatRateComponentBulk(apiClient)
}

func SaveVatRate(cl *api.Client) {
	mngr := cl.SalesManager

	filter := map[string]string{
		"vatRateID": "141",
		"name": "19%",
		"rate":    "0.19",
		"code": "19Percent",
		"active": "1",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := mngr.SaveVatRate(ctx, filter)
	common.Die(err)

	fmt.Println("SaveVatRate", common.ConvertSourceToJsonStrIfPossible(res))
}

func GetVatRate(cl *api.Client) {
	mngr := cl.SalesManager

	filter := map[string]string{
		"active": "1",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := mngr.GetVatRates(ctx, filter)
	common.Die(err)

	fmt.Println("GetVatRate", common.ConvertSourceToJsonStrIfPossible(res))
}

func SaveVatRateBulk(cl *api.Client) {
	cli := cl.SalesManager

	bulkItems := []map[string]interface{}{
		{
			"vatRateID": "141",
			"name": "19%",
			"rate":    "0.19",
			"code": "19Percent",
			"active": "1",
		},
		{
			"name": "20%",
			"rate":    "0.20",
			"code": "20Percent",
			"active": "1",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.SaveVatRateBulk(ctx, bulkItems, map[string]string{})
	common.Die(err)

	fmt.Println("SaveVatRateBulk", common.ConvertSourceToJsonStrIfPossible(resp))
}

func SaveVatRateComponent(cl *api.Client) {
	mngr := cl.SalesManager

	filter := map[string]string{
		"name": "Comp19%",
		"rate":    "0.19",
		"type": "STATE",
		"vatRateID": "141",
		"vatRateComponentID": "247",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := mngr.SaveVatRateComponent(ctx, filter)
	common.Die(err)

	fmt.Println("SaveVatRateComponent", common.ConvertSourceToJsonStrIfPossible(res))
}

func SaveVatRateComponentBulk(cl *api.Client) {
	cli := cl.SalesManager

	bulkItems := []map[string]interface{}{
		{
			"name": "Comp19%",
			"rate":    "0.19",
			"type": "STATE",
			"vatRateID": "142",
			"vatRateComponentID": "247",
		},
		{
			"name": "Comp20%",
			"rate":    "0.20",
			"type": "COUNTY",
			"vatRateID": "142",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := cli.SaveVatRateComponentBulk(ctx, bulkItems, map[string]string{})
	common.Die(err)

	fmt.Println("SaveVatRateComponentBulk", common.ConvertSourceToJsonStrIfPossible(resp))
}
