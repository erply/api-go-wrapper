package api

import "github.com/erply/api-go-wrapper/internal/common"

const (
	GetCountriesMethod                = "getCountries"
	GetEmployeesMethod                = "getEmployees"
	GetBusinessAreasMethod            = "getBusinessAreas"
	GetCurrenciesMethod               = "getCurrencies"
	GetUserOperationsLog              = "getUserOperationsLog"
	GetUserRightsMethod               = "getUserRights"
	logProcessingOfCustomerDataMethod = "logProcessingOfCustomerData"
	createInstallationMethod          = "createInstallation"
	SaveEventMethod                   = "saveEvent"
	GetEvents                         = "getEvents"
)

// SetBaseDomain sets the base domain used by the library for Erply API calls from erply.com to the specified domain
func SetBaseDomain(domain string) {
	common.SetBaseDomain(domain)
}
