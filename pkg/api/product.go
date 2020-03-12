package api

//GetProductsResponse ...
type GetProductsResponse struct {
	Status   Status    `json:"status"`
	Products []Product `json:"records"`
}

//Product ...
type Product struct {
	ProductID int     `json:"productID"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Code2     string  `json:"code2"`
	Code3     *string `json:"code3"`
	Price     float64 `json:"price"`
	UnitName  *string `json:"unitName"`
}

//GetProductUnitsResponse ...
type GetProductUnitsResponse struct {
	Status       Status        `json:"status"`
	ProductUnits []ProductUnit `json:"records"`
}

//ProductUnit ...
type ProductUnit struct {
	UnitID string `json:"unitID"`
	Name   string `json:"name"`
}
