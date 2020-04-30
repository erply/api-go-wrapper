package sales

/*
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

func (cli *Client) CalculateShoppingCart(ctx context.Context, filters map[string]string) (*ShoppingCartTotals, error) {

	resp, err := cli.SendRequest(ctx, api.calculateShoppingCartMethod, filters)
	if err != nil {
		return nil, erro.NewFromError("CalculateShoppingCart: error sending request", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, erro.NewFromError(fmt.Sprintf("CalculateShoppingCart: bad response status code: %d", resp.StatusCode), nil)
	}

	var respData struct {
		Status  common.Status
		Records []*ShoppingCartTotals
	}

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, erro.NewFromError("CalculateShoppingCart: unmarshaling response failed", err)
	}
	if !common.IsJSONResponseOK(&respData.Status) {
		return nil, erro.NewErplyError(strconv.Itoa(respData.Status.ErrorCode), respData.Status.Request+": "+respData.Status.ResponseStatus)
	}
	if len(respData.Records) < 1 {
		return nil, erro.NewFromError("CalculateShoppingCart: no records in response", nil)
	}

	return respData.Records[0], nil
}
*/
