package api

//GetProductsResponse ...
type GetProductsResponse struct {
	Status   Status    `json:"status"`
	Products []Product `json:"records"`
}

type getProductCategoriesResponse struct {
	Status            Status            `json:"status"`
	ProductCategories []ProductCategory `json:"records"`
}

type getProductBrandsResponse struct {
	Status        Status         `json:"status"`
	ProductBrands []ProductBrand `json:"records"`
}

type getProductGroupsResponse struct {
	Status        Status         `json:"status"`
	ProductGroups []ProductGroup `json:"records"`
}

//Product ...
type Product struct {
	ProductID          int                `json:"productID"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	DescriptionLong    string             `json:"longdesc"`
	Status             string             `json:"status"`
	Code               string             `json:"code"`
	Code2              string             `json:"code2"`
	Code3              *string            `json:"code3"`
	Price              float64            `json:"price"`
	UnitName           *string            `json:"unitName"`
	Images             []ProductImage     `json:"images"`
	DisplayedInWebshop byte               `json:"displayedInWebshop"`
	CategoryId         uint               `json:"categoryID"`
	CategoryName       string             `json:"categoryName"`
	BrandID            uint               `json:"brandID"`
	BrandName          string             `json:"brandName"`
	GroupID            uint               `json:"groupID"`
	GroupName          string             `json:"groupName"`
	Warehouses         map[uint]StockInfo `json:"warehouses"`
}

type StockInfo struct {
	WarehouseID   uint    `json:"warehouseID"`
	Reserved      int     `json:"reserved"`
	Free          int     `json:"free"`
	OrderPending  int     `json:"orderPending"`
	ReorderPoint  int     `json:"reorderPoint"`
	RestockLevel  int     `json:"restockLevel"`
	FifoCost      float32 `json:"FIFOCost"`
	PurchasePrice float32 `json:"purchasePrice"`
}

type ProductImage struct {
	PictureID       string  `json:"pictureID"`
	Name            string  `json:"name"`
	ThumbURL        string  `json:"thumbURL"`
	SmallURL        string  `json:"smallURL"`
	LargeURL        string  `json:"largeURL"`
	FullURL         string  `json:"fullURL"`
	External        byte    `json:"external"`
	HostingProvider string  `json:"hostingProvider"`
	Hash            *string `json:"hash"`
	Tenant          *string `json:"tenant"`
}

type ProductCategory struct {
	ProductCategoryID   uint   `json:"productCategoryID"`
	ParentCategoryID    uint   `json:"parentCategoryID"`
	ProductCategoryName string `json:"productCategoryName"`
	Added               uint64 `json:"added"`
	LastModified        uint64 `json:"lastModified"`
}

type ProductBrand struct {
	ID           uint   `json:"brandID"`
	Name         string `json:"name"`
	Added        uint64 `json:"added"`
	LastModified uint64 `json:"lastModified"`
}

type ProductGroup struct {
	ID              uint   `json:"productGroupID"`
	Name            string `json:"name"`
	ShowInWebshop   string `json:"showInWebshop"`
	NonDiscountable byte   `json:"nonDiscountable"`
	PositionNo      int    `json:"positionNo"`
	ParentGroupID   string `json:"parentGroupID"`
	Added           uint64 `json:"added"`
	LastModified    uint64 `json:"lastModified"`
}

//GetProductUnitsResponse ...
type GetProductUnitsResponse struct {
	Status       Status        `json:"status"`
	ProductUnits []ProductUnit `json:"records"`
}

//ProductUnit ...
type ProductUnit struct {
	UnitID string `json:"unitID"`
	Name   string `json:"name"`
}
