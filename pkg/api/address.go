package api

type (
	//GetAddressesResponse ..
	GetAddressesResponse struct {
		Status    Status    `json:"status"`
		Addresses Addresses `json:"records"`
	}

	//Addresses from getAddresses
	Addresses []Address

	//Address from getAddresses
	Address struct {
		AddressID  int    `json:"addressID"`
		Street     string `json:"street"`
		PostalCode string `json:"postalCode"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
	}
)
