package sales

import sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"

type (
	Coupon struct {
		Added                      int    `json:"added"`
		CampaignID                 string `json:"campaignID"`
		Code                       string `json:"code"`
		CouponID                   int    `json:"couponID"`
		Description                string `json:"description"`
		IssuedFromDate             string `json:"issuedFromDate"`
		IssuedUntilDate            string `json:"issuedUntilDate"`
		LastModified               int    `json:"lastModified"`
		Measure                    string `json:"measure"`
		Name                       string `json:"name"`
		PrintedAutomaticallyInPOS  int    `json:"printedAutomaticallyInPOS"`
		PrintingCostInRewardPoints int    `json:"printingCostInRewardPoints"`
		PromptCashier              int    `json:"promptCashier"`
		Threshold                  string `json:"threshold"`
		ThresholdType              string `json:"thresholdType"`
		Treshold                   int    `json:"treshold"`
		TresholdType               string `json:"tresholdType"`
		WarehouseID                string `json:"warehouseID"`
	}

	GetCouponsResponse struct {
		Status  sharedCommon.Status `json:"status"`
		Coupons []Coupon            `json:"records"`
	}
)
