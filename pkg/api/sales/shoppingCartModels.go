package sales

type ShoppingCartTotals struct {
	Rows     []ShoppingCartProduct `json:"rows"`
	NetTotal float64               `json:"netTotal"`
	VATTotal float64               `json:"vatTotal"`
	Total    float64               `json:"total"`
}

type ShoppingCartProduct struct {
	ProductID            string  `json:"productID"`
	Amount               string  `json:"amount"`
	OriginalPrice        float64 `json:"originalPrice"`
	OriginalPriceWithVAT float64 `json:"originalPriceWithVAT"`
	FinalPrice           float64 `json:"finalPrice"`
	FinalPriceWithVAT    float64 `json:"finalPriceWithVAT"`
	RowNetTotal          float64 `json:"rowNetTotal"`
	RowTotal             float64 `json:"rowTotal"`
	Discount             float64 `json:"discount"`
}
