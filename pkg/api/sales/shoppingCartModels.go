package sales

type ShoppingCartTotals struct {
	Rows     []ShoppingCartProduct `json:"rows"`
	NetTotal float64               `json:"netTotal"`
	VATTotal float64               `json:"vatTotal"`
	Total    float64               `json:"total"`
}

type ShoppingCartProduct struct {
	RowNumber            int     `json:"rowNumber"`
	ProductID            string  `json:"productID"`
	Amount               string  `json:"amount"`
	VatRateID            int     `json:"vatrateID"`
	VatRate              string  `json:"vatRate"`
	OriginalPrice        float64 `json:"originalPrice"`
	OriginalPriceWithVAT float64 `json:"originalPriceWithVAT"`
	PromotionDiscount    float64 `json:"promotionDiscount"`
	ManualDiscount       float64 `json:"manualDiscount"`
	Discount             float64 `json:"discount"`
	FinalPrice           float64 `json:"finalPrice"`
	FinalPriceWithVAT    float64 `json:"finalPriceWithVAT"`
	RowNetTotal          float64 `json:"rowNetTotal"`
	RowVAT               float64 `json:"rowVAT"`
	RowTotal             float64 `json:"rowTotal"`
}
