package products

import (
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	GetProductsResponse struct {
		Status   sharedCommon.Status `json:"status"`
		Products []Product           `json:"records"`
	}

	getProductCategoriesResponse struct {
		Status            sharedCommon.Status `json:"status"`
		ProductCategories []ProductCategory   `json:"records"`
	}

	getProductBrandsResponse struct {
		Status        sharedCommon.Status `json:"status"`
		ProductBrands []ProductBrand      `json:"records"`
	}

	getProductGroupsResponse struct {
		Status        sharedCommon.Status `json:"status"`
		ProductGroups []ProductGroup      `json:"records"`
	}

	ProductDimensions struct {
		Name             string `json:"name"`
		Value            string `json:"value"`
		Order            int    `json:"order"`
		DimensionID      int    `json:"dimensionID"`
		DimensionValueID int    `json:"dimensionValueID"`
	}

	ProductVariaton struct {
		ProductID  string              `json:"productID"`
		Name       string              `json:"name"`
		Code       string              `json:"code"`
		Code2      string              `json:"code2"`
		Dimensions []ProductDimensions `json:"dimensions"`
	}

	//Payload ...
	Product struct {
		ProductID          int                `json:"productID"`
		Active             int                `json:"active"`
		Name               string             `json:"name"`
		NameEng            string             `json:"nameENG"`
		NameSpa            string             `json:"nameSPA"`
		NameEst            string             `json:"nameEST"`
		NameGer            string             `json:"nameGER"`
		NameSwe            string             `json:"nameSWE"`
		NameFin            string             `json:"nameFIN"`
		NameRus            string             `json:"nameRUS"`
		NameLat            string             `json:"nameLAT"`
		NameLit            string             `json:"nameLIT"`
		NameGre            string             `json:"nameGRE"`
		Code               string             `json:"code"`
		Code2              string             `json:"code2"`
		Code3              *string            `json:"code3"`
		GroupID            uint               `json:"groupID"`
		Price              float64            `json:"price"`
		DisplayedInWebshop byte               `json:"displayedInWebshop"`
		BrandID            uint               `json:"brandID"`
		Description        string             `json:"description"`
		DescriptionLong    string             `json:"longdesc"`
		DescriptionEng     string             `json:"descriptionENG"`
		DescriptionLongEng string             `json:"longdescENG"`
		DescriptionSpa     string             `json:"descriptionSPA"`
		DescriptionLongSpa string             `json:"longdescSPA"`
		DescriptionEst     string             `json:"descriptionEST"`
		DescriptionLongEst string             `json:"longdescEST"`
		DescriptionGer     string             `json:"descriptionGER"`
		DescriptionLongGer string             `json:"longdescGER"`
		DescriptionSwe     string             `json:"descriptionSWE"`
		DescriptionLongSwe string             `json:"longdescSWE"`
		DescriptionFin     string             `json:"descriptionFIN"`
		DescriptionLongFin string             `json:"longdescFIN"`
		DescriptionRus     string             `json:"descriptionRUS"`
		DescriptionLongRus string             `json:"longdescRUS"`
		DescriptionLat     string             `json:"descriptionLAT"`
		DescriptionLongLat string             `json:"longdescLAT"`
		DescriptionLit     string             `json:"descriptionLIT"`
		DescriptionLongLit string             `json:"longdescLIT"`
		DescriptionGre     string             `json:"descriptionGRE"`
		DescriptionLongGre string             `json:"longdescGRE"`
		Added              uint64             `json:"added"`
		LastModified       uint64             `json:"lastModified"`
		Vatrate            float64            `json:"vatrate"`
		PriceWithVat       float32            `json:"priceWithVat"`
		GrossWeight        string             `json:"grossWeight"`
		NetWeight          string             `json:"netWeight"`
		UnitName           *string            `json:"unitName"`
		BrandName          string             `json:"brandName"`
		GroupName          string             `json:"groupName"`
		CategoryId         uint               `json:"categoryID"`
		CategoryName       string             `json:"categoryName"`
		Status             string             `json:"status"`
		SupplierID         int                `json:"supplierID"`
		Images             []ProductImage     `json:"images"`
		ProductVariations  []string           `json:"productVariations"` // Variations of matrix product
		ParentProductID    int                `json:"parentProductID"`
		NonStockProduct    int                `json:"nonStockProduct"`
		TaxFree            int                `json:"taxFree"`
		Type               string             `json:"type"`
		Warehouses         map[uint]StockInfo `json:"warehouses"`

		RelatedProducts []string          `json:"relatedProducts"`
		VariationList   []ProductVariaton `json:"variationList"`
	}

	StockInfo struct {
		WarehouseID   uint    `json:"warehouseID"`
		Free          float64 `json:"free"`
		OrderPending  int     `json:"orderPending"`
		ReorderPoint  int     `json:"reorderPoint"`
		RestockLevel  int     `json:"restockLevel"`
		FifoCost      float32 `json:"FIFOCost"`
		PurchasePrice float32 `json:"purchasePrice"`
	}

	ProductImage struct {
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

	ProductCategory struct {
		ProductCategoryID   uint   `json:"productCategoryID"`
		ParentCategoryID    uint   `json:"parentCategoryID"`
		ProductCategoryName string `json:"productCategoryName"`
		Added               uint64 `json:"added"`
		LastModified        uint64 `json:"lastModified"`
	}

	ProductBrand struct {
		ID           uint   `json:"brandID"`
		Name         string `json:"name"`
		Added        uint64 `json:"added"`
		LastModified uint64 `json:"lastModified"`
	}

	ProductGroup struct {
		ID              uint                `json:"productGroupID"`
		Name            string              `json:"name"`
		ShowInWebshop   string              `json:"showInWebshop"`
		NonDiscountable byte                `json:"nonDiscountable"`
		PositionNo      int                 `json:"positionNo"`
		ParentGroupID   string              `json:"parentGroupID"`
		Added           uint64              `json:"added"`
		LastModified    uint64              `json:"lastModified"`
		SubGroups       []ProductGroup      `json:"subGroups"`
		Attributes      []map[string]string `json:"attributes,omitempty"`
	}

	//GetProductUnitsResponse ...
	GetProductUnitsResponse struct {
		Status       sharedCommon.Status `json:"status"`
		ProductUnits []ProductUnit       `json:"records"`
	}

	//ProductUnit ...
	ProductUnit struct {
		UnitID string `json:"unitID"`
		Name   string `json:"name"`
	}

	GetProductsResponseBulkItem struct {
		Status   sharedCommon.StatusBulk `json:"status"`
		Products []Product               `json:"records"`
	}

	GetProductsResponseBulk struct {
		Status    sharedCommon.Status           `json:"status"`
		BulkItems []GetProductsResponseBulkItem `json:"requests"`
	}
)
