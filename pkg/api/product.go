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

//Product ...
type Product struct {
	ProductID          int            `json:"productID"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	Status             string         `json:"status"`
	Code               string         `json:"code"`
	Code2              string         `json:"code2"`
	Code3              *string        `json:"code3"`
	Price              float64        `json:"price"`
	UnitName           *string        `json:"unitName"`
	Images             []ProductImage `json:"images"`
	DisplayedInWebshop byte           `json:"displayedInWebshop"`
	CategoryId         uint           `json:"categoryID"`
	CategoryName       string         `json:"categoryName"`
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
