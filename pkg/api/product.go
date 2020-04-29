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

type ProductDimensions struct {
	Name             string `json:"name"`
	Value            string `json:"value"`
	Order            int    `json:"order"`
	DimensionID      int    `json:"dimensionID"`
	DimensionValueID int    `json:"dimensionValueID"`
}

type ProductVariaton struct {
	ProductID  string              `json:"productID"`
	Name       string              `json:"name"`
	Code       string              `json:"code"`
	Code2      string              `json:"code2"`
	Dimensions []ProductDimensions `json:"dimensions"`
}

//Product ...
type Product struct {
	ProductID          int                `json:"productID"`
	ParentProductID    int                `json:"parentProductID"`
	Type               string             `json:"type"`
	Name               string             `json:"name"`
	NameEng            string             `json:"nameENG"`
	Description        string             `json:"description"`
	DescriptionEng     string             `json:"descriptionENG"`
	DescriptionLong    string             `json:"longdesc"`
	DescriptionLongEng string             `json:"longdescENG"`
	Status             string             `json:"status"`
	Code               string             `json:"code"`
	Code2              string             `json:"code2"`
	Code3              *string            `json:"code3"`
	Price              float64            `json:"price"`
	PriceWithVat       float32            `json:"priceWithVat"`
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
	RelatedProducts    []string           `json:"relatedProducts"`
	Vatrate            float64            `json:"vatrate"`
	ProductVariations  []string           `json:"productVariations"` // Variations of matrix product
	VariationList      []ProductVariaton  `json:"variationList"`
}

type StockInfo struct {
	WarehouseID   uint    `json:"warehouseID"`
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
	ID              uint           `json:"productGroupID"`
	Name            string         `json:"name"`
	ShowInWebshop   string         `json:"showInWebshop"`
	NonDiscountable byte           `json:"nonDiscountable"`
	PositionNo      int            `json:"positionNo"`
	ParentGroupID   string         `json:"parentGroupID"`
	Added           uint64         `json:"added"`
	LastModified    uint64         `json:"lastModified"`
	SubGroups       []ProductGroup `json:"subGroups"`
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
