package api

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type UserRights struct {
	UserName string `json:"userName"`
}

type GetCountriesResponse struct {
	Status    sharedCommon.Status `json:"status"`
	Countries []Country           `json:"records"`
}
type GetEmployeesResponse struct {
	Status    sharedCommon.Status `json:"status"`
	Employees []Employee          `json:"records"`
}
type GetBusinessAreasResponse struct {
	Status        sharedCommon.Status `json:"status"`
	BusinessAreas []BusinessArea      `json:"records"`
}
type GetCurrenciesResponse struct {
	Status     sharedCommon.Status `json:"status"`
	Currencies []Currency          `json:"records"`
}

type GetUserOperationsLogResponse struct {
	Status        sharedCommon.Status `json:"status"`
	OperationLogs []OperationLog      `json:"records"`
}
type OperationLog struct {
	LogID     int    `json:"logID"`
	Username  string `json:"username"`
	Timestamp uint64 `json:"timestamp"`
	TableName string `json:"tableName"`
	ItemID    int    `json:"itemID"`
	Operation string `json:"operation"`
}

type GetUserRightsResponse struct {
	Status  sharedCommon.Status `json:"status"`
	Records []UserRights        `json:"records"`
}

type Country struct {
	CountryId             uint   `json:"countryID"`
	CountryName           string `json:"countryName"`
	CountryCode           string `json:"countryCode"`
	MemberOfEuropeanUnion byte   `json:"memberOfEuropeanUnion"`
	LastModified          uint64 `json:"lastModified"`
	Added                 uint64 `json:"added"`
}

type Event struct {
	EventID       string `json:"eventID"`
	ID            string `json:"id"`
	Description   string `json:"description"`
	TypeID        string `json:"typeID"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	CustomerID    string `json:"customerID"`
	ContactID     string `json:"contactID"`
	ProjectID     string `json:"projectID"`
	EmployeeID    string `json:"employeeID"`
	SubmitterID   string `json:"submitterID"`
	SupplierID    string `json:"supplierID"`
	SupplierName  string `json:"supplierName"`
	StatusID      string `json:"statusID"`
	ResourceID    string `json:"resourceID"`
	Notes         string `json:"notes"`
	LastModified  string `json:"lastModified"`
	ContactName   string `json:"contactName"`
	CustomerName  string `json:"customerName"`
	EmployeeName  string `json:"employeeName"`
	SubmitterName string `json:"submitterName"`
	ProjectName   string `json:"projectName"`
	ResourceName  string `json:"resourceName"`
	StatusName    string `json:"statusName"`
	TypeName      string `json:"typeName"`
	Completed     string `json:"completed"`
}

type GetEventsResponse struct {
	Status sharedCommon.Status `json:"status"`
	Events []Event             `json:"records"`
}

type SaveEventResponse struct {
	Status  sharedCommon.Status
	Records []struct {
		EventID int `json:"eventID"`
	} `json:"records"`
}
type Employee struct {
	EmployeeID             string                      `json:"employeeID"`
	FullName               string                      `json:"fullName"`
	EmployeeName           string                      `json:"employeeName"`
	FirstName              string                      `json:"firstName"`
	LastName               string                      `json:"lastName"`
	Phone                  string                      `json:"phone"`
	Mobile                 string                      `json:"mobile"`
	Email                  string                      `json:"email"`
	Fax                    string                      `json:"fax"`
	Code                   string                      `json:"code"`
	Gender                 string                      `json:"gender"`
	UserID                 string                      `json:"userID"`
	Username               string                      `json:"username"`
	UserGroupID            string                      `json:"userGroupID"`
	Warehouses             []EmployeeWarehouse         `json:"warehouses"`
	PointsOfSale           string                      `json:"pointsOfSale"`
	ProductIDs             []EmployeeProduct           `json:"productIDs"`
	Attributes             []sharedCommon.ObjAttribute `json:"attributes"`
	LastModified           uint64                      `json:"lastModified"`
	LastModifiedByUserName string                      `json:"lastModifiedByUserName"`

	// detail fileds
	Skype        string `json:"skype"`
	Birthday     string `json:"birthday"`
	JobTitleID   uint   `json:"jobTitleID"`
	JobTitleName string `json:"jobTitleName"`
	Notes        string `json:"notes"`
	Added        uint64 `json:"added"`
}

type EmployeeWarehouse struct {
	Id uint `json:"id"`
}

type EmployeeProduct struct {
	ProductID    uint   `json:"productID"`
	ProductCode  string `json:"productCode"`
	ProductName  string `json:"productName"`
	ProductGroup uint   `json:"productGroup"`
}

type BusinessArea struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Added        uint64 `json:"added"`
	LastModified uint64 `json:"lastModified"`
}

type Currency struct {
	CurrencyID   string `json:"currencyID"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Default      string `json:"default"`
	NameShort    string `json:"nameShort"`
	NameFraction string `json:"nameFraction"`
	Added        string `json:"added"`
	LastModified string `json:"lastModified"`
}
