package api

type Customer struct {
	ID                int            `json:"id"`
	CustomerID        int            `json:"customerID"`
	TypeID            string         `json:"type_id"`
	FullName          string         `json:"fullName"`
	CompanyName       string         `json:"companyName"`
	FirstName         string         `json:"firstName"`
	LastName          string         `json:"lastName"`
	GroupID           int            `json:"groupID"`
	EDI               string         `json:"EDI"`
	Phone             string         `json:"phone"`
	EInvoiceEmail     string         `json:"eInvoiceEmail"`
	Email             string         `json:"email"`
	Fax               string         `json:"fax"`
	Code              string         `json:"code"`
	ReferenceNumber   string         `json:"referenceNumber"`
	VatNumber         string         `json:"vatNumber"`
	BankName          string         `json:"bankName"`
	BankAccountNumber string         `json:"bankAccountNumber"`
	BankIBAN          string         `json:"bankIBAN"`
	BankSWIFT         string         `json:"bankSWIFT"`
	PaymentDays       int            `json:"paymentDays"`
	Notes             string         `json:"notes"`
	LastModified      int            `json:"lastModified"`
	CustomerType      string         `json:"customerType"`
	Address           string         `json:"address"`
	CustomerAddresses Addresses      `json:"addresses"`
	Street            string         `json:"street"`
	Address2          string         `json:"address2"`
	City              string         `json:"city"`
	PostalCode        string         `json:"postalCode"`
	Country           string         `json:"country"`
	State             string         `json:"state"`
	ContactPersons    ContactPersons `json:"contactPersons"`
}
type Customers []Customer

//Attribute field
type Attribute struct {
	Name  string `json:"attributeNam"`
	Type  string `json:"attributeType"`
	Value string `json:"attributeValue"`
}

type CustomerConstructor struct {
	CompanyName       string
	Address           string
	PostalCode        string
	Country           string
	FullName          string
	RegistryCode      string
	VatNumber         string
	Email             string
	Phone             string
	BankName          string
	BankAccountNumber string
}
