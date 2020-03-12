package api

type NetTotalsByRates []NetTotalsByRate
type NetTotalsByRate struct {
	//Num1 float64 `json:"1"`
}
type VatTotalsByRates []VatTotalsByRate
type VatTotalsByRate struct {
	//Num1 float64 `json:"1"`
}
type VatRate struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rate   string `json:"rate"`
	Code   string `json:"code"`
	Active string `json:"active"`
	//Added        string `json:"added"`
	LastModified string `json:"lastModified"`
	//IsReverseVat int    `json:"isReverseVat"`
	//ReverseRate int `json:"reverseRate"`
}

type VatRates []VatRate
type VatTotalsByTaxRates []VatTotalsByTaxRate

type VatTotalsByTaxRate struct {
	VatrateID int     `json:"vatrateID"`
	Total     float64 `json:"total"`
}
type NetTotalsByTaxRates []NetTotalsByTaxRate
type NetTotalsByTaxRate struct {
	VatrateID int     `json:"vatrateID"`
	Total     float64 `json:"total"`
}
