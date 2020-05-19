package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type structToConvert struct {
	SupplierId      uint           `json:"supplierID"`
	SupplierType    string         `json:"supplierType"`
	FullName        string         `json:"fullName"`
	CompanyName     string         `json:"companyName"`
	FirstName       string         `json:"firstName"`
	LstName         string         `json:"lastName"`
	GroupId         uint           `json:"groupID"`
	GroupName       string         `json:"groupName"`
	Phone           string         `json:"phone"`
	Mobile          string         `json:"mobile"`
	Email           string         `json:"email"`
	Fax             string         `json:"fax"`
	Code            string         `json:"code"`
	IntegrationCode string         `json:"integrationCode"`
	VatrateID       uint           `json:"vatrateID"`
	CurrencyCode    string         `json:"currencyCode"`
	DeliveryTermsID uint           `json:"deliveryTermsID"`
	CountryId       uint           `json:"countryID"`
	CountryName     string         `json:"countryName"`
	CountryCode     string         `json:"countryCode"`
	Address         string         `json:"address"`
	Gln             string         `json:"GLN"`
	Attributes      []ObjAttribute `json:"attributes"`

	// Detail fields
	VatNumber           string `json:"vatNumber"`
	Skype               string `json:"skype"`
	Website             string `json:"website"`
	BankName            string `json:"bankName"`
	BankAccountNumber   string `json:"bankAccountNumber"`
	BankIBAN            string `json:"bankIBAN"`
	BankSWIFT           string `json:"bankSWIFT"`
	Birthday            string `json:"birthday"`
	CompanyID           uint   `json:"companyID"`
	ParentCompanyName   string `json:"parentCompanyName"`
	SupplierManagerID   uint   `json:"supplierManagerID"`
	SupplierManagerName string `json:"supplierManagerName"`
	PaymentDays         uint   `json:"paymentDays"`
	Notes               string `json:"notes"`
	LastModified        string `json:"lastModified"`
	Added               uint64 `json:"added"`
}

func TestConvertingStructToMap(t *testing.T) {
	s := structToConvert{
		SupplierId:          1,
		SupplierType:        "some type",
		FullName:            "Some full name",
		CompanyName:         "Some company",
		FirstName:           "some first",
		LstName:             "some last",
		GroupId:             3,
		GroupName:           "some group",
		Phone:               "3334444444",
		Mobile:              "341431434",
		Email:               "no@mail.me",
		Fax:                 "341234343241",
		Code:                "32413",
		IntegrationCode:     "341324",
		VatrateID:           5,
		CurrencyCode:        "eur",
		DeliveryTermsID:     7,
		CountryId:           6,
		CountryName:         "Deutschland",
		CountryCode:         "DE",
		Address:             "Elm Str 11",
		Gln:                 "gln222",
		Attributes:          []ObjAttribute{},
		VatNumber:           "3431241",
		Skype:               "nono",
		Website:             "ya.ru",
		BankName:            "some swiss bank",
		BankAccountNumber:   "3413412434",
		BankIBAN:            "341t45243535",
		BankSWIFT:           "some swift",
		Birthday:            "11.11.2011",
		CompanyID:           9,
		ParentCompanyName:   "some parent",
		SupplierManagerID:   10,
		SupplierManagerName: "Some manager",
		PaymentDays:         11,
		Notes:               "some notes",
		LastModified:        "2001-01-01 00:00:00",
		Added:               1,
	}

	actualMap, err := ConvertStructToMap(s)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	expectedMap := map[string]string{
		"GLN": "gln222",
		"added": "1",
		"address": "Elm Str 11",
		"attributes": "[]",
		"bankAccountNumber": "3413412434",
		"bankIBAN": "341t45243535",
		"bankName": "some swiss bank",
		"bankSWIFT": "some swift",
		"birthday": "11.11.2011", 
		"code": "32413", 
		"companyID": "9", 
		"companyName": "Some company", 
		"countryCode": "DE", 
		"countryID": "6", 
		"countryName": "Deutschland", 
		"currencyCode": "eur", 
		"deliveryTermsID": "7", 
		"email": "no@mail.me", 
		"fax": "341234343241", 
		"firstName": "some first", 
		"fullName": "Some full name", 
		"groupID": "3", 
		"groupName": "some group", 
		"integrationCode": "341324", 
		"lastModified": "2001-01-01 00:00:00", 
		"lastName": "some last", 
		"mobile": "341431434", 
		"notes": "some notes", 
		"parentCompanyName": "some parent", 
		"paymentDays": "11", 
		"phone": "3334444444", 
		"skype": "nono", 
		"supplierID": "1",
		"supplierManagerID": "10", 
		"supplierManagerName": "Some manager", 
		"supplierType": "some type", 
		"vatNumber": "3431241", 
		"vatrateID": "5", 
		"website": "ya.ru",
	}
	assert.Equal(t, expectedMap, actualMap)
}
