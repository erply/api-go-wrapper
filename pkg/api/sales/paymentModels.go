package sales

import sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"

type (
	PaymentAttribute struct {
		AttributeName  string `json:"attributeName"`
		AttributeType  string `json:"attributeType"`
		AttributeValue string `json:"attributeValue"`
	}
	PaymentStatus string
	PaymentType   string

	PaymentInfo struct {
		DocumentID             int    `json:"documentID"` // Invoice ID
		PaymentID              int    `json:"paymentID"`
		CustomerID             int    `json:"customerID"`
		TypeID                 string `json:"typeID"`
		BankTransactionID      int    `json:"bankTransactionID"`
		Type                   string `json:"type"` // CASH, TRANSFER, CARD, CREDIT, GIFTCARD, CHECK, TIP
		Date                   string `json:"date"`
		Sum                    string `json:"sum"`
		CardHolder             string `json:"cardHolder"`
		CardType               string `json:"cardType"`
		CardNumber             string `json:"cardNumber"`
		AuthorizationCode      string `json:"authorizationCode"`
		ReferenceNumber        string `json:"referenceNumber"`
		CurrencyRate           string `json:"currencyRate"`
		CashPaid               string `json:"cashPaid"`
		CashChange             string `json:"cashChange"`
		CurrencyCode           string `json:"currencyCode"` // EUR, USD
		Info                   string `json:"info"`         // Information about the payer or payment transaction
		Added                  uint64 `json:"added"`
		IsPrepayment           uint64 `json:"isPrepayment"`
		StoreCredit            uint64 `json:"storeCredit"`
		BankAccount            string `json:"bankAccount"`
		BankDocumentNumber     string `json:"bankDocumentNumber"`
		BankDate               string `json:"bankDate"`
		BankPayerAccount       string `json:"bankPayerAccount"`
		BankPayerName          string `json:"bankPayerName"`
		BankPayerCode          string `json:"bankPayerCode"`
		BankSum                string `json:"bankSum"`
		BankReferenceNumber    string `json:"bankReferenceNumber"`
		BankDescription        string `json:"bankDescription"`
		BankCurrency           string `json:"bankCurrency"`
		ArchivalNumber         string `json:"archivalNumber"`
		PaymentServiceProvider string `json:"paymentServiceProvider"`
		Aid                    string `json:"aid"`
		ApplicationLabel       string `json:"applicationLabel"`
		PinStatement           string `json:"pinStatement"`
		CryptogramType         string `json:"cryptogramType"`
		Cryptogram             string `json:"cryptogram"`
		ExpirationDate         string `json:"expirationDate"`
		EntryMethod            string `json:"entryMethod"`
		TransactionNumber      string `json:"transactionNumber"`
		TransactionId          string `json:"transactionId"`
		TransactionType        string `json:"transactionType"`
		TransactionTime        int64  `json:"transactionTime"`
		KlarnaPaymentID        string `json:"klarnaPaymentID"`
		CertificateBalance     string `json:"certificateBalance"`
		StatusCode             string `json:"statusCode"`
		StatusMessage          string `json:"statusMessage"`
		GiftCardVatRateID      int    `json:"giftCardVatRateID"`
		LastModified           uint64 `json:"lastModified"`
	}

	GetPaymentsBulkItem struct {
		Status       sharedCommon.StatusBulk `json:"status"`
		PaymentInfos []PaymentInfo           `json:"records"`
	}

	GetPaymentsResponseBulk struct {
		Status    sharedCommon.Status   `json:"status"`
		BulkItems []GetPaymentsBulkItem `json:"requests"`
	}
)
