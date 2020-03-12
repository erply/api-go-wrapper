package api

type PaymentAttribute struct {
	AttributeName  string `json:"attributeName"`
	AttributeType  string `json:"attributeType"`
	AttributeValue string `json:"attributeValue"`
}
type PaymentStatus string
type PaymentType string
