package api

type (
	//CompanyInfos ..
	CompanyInfos []CompanyInfo
	//CompanyInfo ..
	CompanyInfo struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		Code               string `json:"code"`
		VAT                string `json:"VAT"`
		Phone              string `json:"phone"`
		Mobile             string `json:"mobile"`
		Fax                string `json:"fax"`
		Email              string `json:"email"`
		Web                string `json:"web"`
		BankAccountNumber  string `json:"bankAccountNumber"`
		BankName           string `json:"bankName"`
		BankSWIFT          string `json:"bankSWIFT"`
		BankIBAN           string `json:"bankIBAN"`
		BankAccountNumber2 string `json:"bankAccountNumber2"`
		BankName2          string `json:"bankName2"`
		BankSWIFT2         string `json:"bankSWIFT2"`
		BankIBAN2          string `json:"bankIBAN2"`
		Address            string `json:"address"`
		Country            string `json:"country"`

		//field for ConfParameters
		ConfParameters ConfParameter
	}
)
