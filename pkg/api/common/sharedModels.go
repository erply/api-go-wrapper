package common

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type (
	//Addresses from getAddresses
	Addresses []Address

	//Address from getAddresses
	Address struct {
		AddressID        int
		OwnerID          int
		TypeID           int
		TypeActivelyUsed int
		Added            int64
		Address2         string
		TypeName         string
		Address          string
		Street           string
		PostalCode       string
		City             string
		State            string
		Country          string
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

func (u *Address) UnmarshalJSON(data []byte) error {

	raw := struct {
		AddressID        int         `json:"addressID"`
		OwnerID          int         `json:"ownerID"`
		TypeID           json.Number `json:"typeID"`
		TypeActivelyUsed int         `json:"typeActivelyUsed"`
		Added            int64       `json:"added"`
		Address2         string      `json:"address2"`
		TypeName         string      `json:"typeName"`
		Address          string      `json:"address"`
		Street           string      `json:"street"`
		PostalCode       string      `json:"postalCode"`
		City             string      `json:"city"`
		State            string      `json:"state"`
		Country          string      `json:"country"`
		LastModified
		Attributes
	}{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	u.AddressID = raw.AddressID
	u.OwnerID = raw.OwnerID
	typeID, err := raw.TypeID.Int64()
	if err != nil {
		return errors.Wrapf(err, "unable to unmarshal address. typeId did not contain an int: %s", raw.TypeID.String())
	}
	u.TypeID = int(typeID)

	u.TypeActivelyUsed = raw.TypeActivelyUsed
	u.Added = raw.Added
	u.Address2 = raw.Address2
	u.TypeName = raw.TypeName
	u.Address = raw.Address
	u.Street = raw.Street
	u.PostalCode = raw.PostalCode
	u.City = raw.City
	u.State = raw.State
	u.Country = raw.Country
	u.LastModified = raw.LastModified
	u.Attributes = raw.Attributes
	return nil
}
