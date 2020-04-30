package sales

type (
	PaymentAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeType  string `json:"attributeType"`
		AttributeValue string `json:"attributeValue"`
	}
	PaymentStatus string
	PaymentType   string

	PaymentInfo struct {
		DocumentID   int    `json:"documentID"` // Invoice ID
		Type         string `json:"type"`       // CASH, TRANSFER, CARD, CREDIT, GIFTCARD, CHECK, TIP
		Date         string `json:"date"`
		Sum          string `json:"sum"`
		CurrencyCode string `json:"currencyCode"` // EUR, USD
		Info         string `json:"info"`         // Information about the payer or payment transaction
		Added        uint64 `json:"added"`
	}
)
