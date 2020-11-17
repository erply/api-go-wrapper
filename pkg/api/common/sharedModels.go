package common

type (
	//Addresses from getAddresses
	Addresses []Address

	//Address from getAddresses
	Address struct {
		AddressID        int    `json:"addressID"`
		OwnerID          int    `json:"ownerID"`
		TypeID           int    `json:"typeID"`
		TypeActivelyUsed int    `json:"typeActivelyUsed"`
		Added            int64  `json:"added"`
		Address2         string `json:"address2"`
		TypeName         string `json:"typeName"`
		Address          string `json:"address"`
		Street           string `json:"street"`
		PostalCode       string `json:"postalCode"`
		City             string `json:"city"`
		State            string `json:"state"`
		Country          string `json:"country"`
		LastModified
		Attributes
	}

	Attributes struct {
		Attributes []ObjAttribute `json:"attributes"`
	}

	ObjAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeType  string `json:"attributeType"`
		AttributeValue string `json:"attributeValue"`
	}

	LongAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeValue string `json:"attributeValue"`
	}

	LongAttributes struct {
		LongAttributes []LongAttribute `json:"longAttributes"`
	}

	LastModified struct {
		LastModified           int64  `json:"lastModified"`
		LastModifierEmployeeID int64  `json:"lastModifierEmployeeID"`
		LastModifierUsername   string `json:"lastModifierUsername"`
	}
)
