package common

type (
	//Addresses from getAddresses
	Addresses []Address

	//Address from getAddresses
	Address struct {
		AddressID  int         `json:"addressID"`
		OwnerID    int         `json:"ownerID"`
		TypeID     interface{} `json:"typeID"`
		Address2   string      `json:"address2"`
		Address    string      `json:"address"`
		Street     string      `json:"street"`
		PostalCode string      `json:"postalCode"`
		City       string      `json:"city"`
		State      string      `json:"state"`
		Country    string      `json:"country"`
	}

	ObjAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeType  string `json:"attributeType"`
		AttributeValue string `json:"attributeValue"`
	}
)
