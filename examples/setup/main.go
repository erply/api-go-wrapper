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

	GetEmployeesBulk(apiClient)
}

func GetEmployeesBulk(cl *api.Client) {
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

	bulkResp, err := cl.GetEmployeesBulk(ctx, bulkFilters, map[string]string{})
	common.Die(err)

	fmt.Println("GetEmployeesBulk:")
	fmt.Println(common.ConvertSourceToJsonStrIfPossible(bulkResp))
}
