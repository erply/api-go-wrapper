package products

import (
	"encoding/json"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
)

const (
	ResponseTypeCSV = "CSV"
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
		ProductID int `json:"productID"`
		Active    int `json:"active"`
		NameLanguages
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
		ContainerID        int                `json:"containerID"`
		Type               string             `json:"type"`
		Warehouses         map[uint]StockInfo `json:"warehouses"`
		Parameters         []Parameter        `json:"parameters"`
		RelatedProducts    []string           `json:"relatedProducts"`
		VariationList      []ProductVariaton  `json:"variationList"`
	}

	Parameter struct {
		ID      string `json:"parameterID"`
		Name    string `json:"parameterName"`
		Type    string `json:"parameterType"`
		GroupID string `json:"parameterGroupID"`
		Value   string `json:"parameterValue"`
	}
	StockInfo struct {
		WarehouseID   uint    `json:"warehouseID"`
		Free          float64 `json:"free"`
		OrderPending  int     `json:"orderPending"`
		ReorderPoint  int     `json:"reorderPoint"`
		RestockLevel  float64 `json:"restockLevel"`
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
		ID uint `json:"productGroupID"`
		NameLanguages
		ShowInWebshop   string              `json:"showInWebshop"`
		NonDiscountable byte                `json:"nonDiscountable"`
		PositionNo      int                 `json:"positionNo"`
		ParentGroupID   string              `json:"parentGroupID"`
		Added           uint64              `json:"added"`
		LastModified    uint64              `json:"lastModified"`
		SubGroups       []ProductGroup      `json:"subGroups"`
		Attributes      []map[string]string `json:"attributes,omitempty"`
		Images          []ProductGroupImage `json:"images"`
	}

	ProductGroupImage struct {
		PictureID string `json:"pictureID"`
		ThumbURL  string `json:"thumbURL"`
		SmallURL  string `json:"smallURL"`
		LargeURL  string `json:"largeURL"`
	}

	NameLanguages struct {
		Name    string `json:"name"`
		NameEng string `json:"nameENG"`
		NameSpa string `json:"nameSPA"`
		NameEst string `json:"nameEST"`
		NameGer string `json:"nameGER"`
		NameSwe string `json:"nameSWE"`
		NameFin string `json:"nameFIN"`
		NameRus string `json:"nameRUS"`
		NameLat string `json:"nameLAT"`
		NameLit string `json:"nameLIT"`
		NameGre string `json:"nameGRE"`
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

	GetProductStockResponse struct {
		Status          sharedCommon.Status `json:"status"`
		GetProductStock []GetProductStock   `json:"records"`
	}

	GetProductStock struct {
		ProductID              int         `json:"productID"`
		AmountInStock          json.Number `json:"amountInStock"`
		AmountReserved         float64     `json:"amountReserved"`
		SuggestedPurchasePrice float64     `json:"suggestedPurchasePrice"`
		AveragePurchasePrice   float64     `json:"averagePurchasePrice"`
		AverageCost            float64     `json:"averageCost"`
		FirstPurchaseDate      string      `json:"firstPurchaseDate"`
		LastPurchaseDate       string      `json:"lastPurchaseDate"`
		LastSoldDate           string      `json:"lastSoldDate"`
		ReorderPoint           int         `json:"reorderPoint"`
		RestockLevel           float64     `json:"restockLevel"`
	}

	GetProductStockFileResponse struct {
		Status              sharedCommon.Status   `json:"status"`
		GetProductStockFile []GetProductStockFile `json:"records"`
	}

	GetProductStockFile struct {
		ReportLink string `json:"reportLink"`
	}

	GetProductStockFileResponseBulkItem struct {
		Status               sharedCommon.StatusBulk `json:"status"`
		GetProductStockFiles []GetProductStockFile   `json:"records"`
	}

	GetProductStockFileResponseBulk struct {
		Status    sharedCommon.Status                   `json:"status"`
		BulkItems []GetProductStockFileResponseBulkItem `json:"requests"`
	}

	SaveProductResult struct {
		ProductID int `json:"productID"`
	}

	SaveProductResponse struct {
		Status             sharedCommon.Status `json:"status"`
		SaveProductResults []SaveProductResult `json:"records"`
	}

	SaveProductResponseBulkItem struct {
		Status   sharedCommon.StatusBulk `json:"status"`
		Products []SaveProductResult     `json:"records"`
	}

	SaveProductResponseBulk struct {
		Status    sharedCommon.Status           `json:"status"`
		BulkItems []SaveProductResponseBulkItem `json:"requests"`
	}

	DeleteProductResponse struct {
		Status sharedCommon.Status `json:"status"`
	}

	DeleteProductResponseBulkItem struct {
		Status sharedCommon.StatusBulk `json:"status"`
	}

	DeleteProductResponseBulk struct {
		Status    sharedCommon.Status             `json:"status"`
		BulkItems []DeleteProductResponseBulkItem `json:"requests"`
	}

	SaveAssortmentResult struct {
		AssortmentID int `json:"assortmentID"`
	}

	SaveAssortmentResponse struct {
		Status                sharedCommon.Status    `json:"status"`
		SaveAssortmentResults []SaveAssortmentResult `json:"records"`
	}

	SaveAssortmentResponseBulkItem struct {
		Status                sharedCommon.StatusBulk `json:"status"`
		SaveAssortmentResults []SaveAssortmentResult  `json:"records"`
	}

	SaveAssortmentResponseBulk struct {
		Status    sharedCommon.Status              `json:"status"`
		BulkItems []SaveAssortmentResponseBulkItem `json:"requests"`
	}

	AddAssortmentProductsResult struct {
		ProductsAlreadyInAssortment string `json:"productsAlreadyInAssortment"`
		NonExistingIDs              string `json:"nonExistingIDs"`
	}

	AddAssortmentProductsResponse struct {
		Status                       sharedCommon.Status           `json:"status"`
		AddAssortmentProductsResults []AddAssortmentProductsResult `json:"records"`
	}

	AddAssortmentProductsResponseBulkItem struct {
		Status                       sharedCommon.StatusBulk       `json:"status"`
		AddAssortmentProductsResults []AddAssortmentProductsResult `json:"records"`
	}

	AddAssortmentProductsResponseBulk struct {
		Status    sharedCommon.Status                     `json:"status"`
		BulkItems []AddAssortmentProductsResponseBulkItem `json:"requests"`
	}

	EditAssortmentProductsResult struct {
		ProductsNotInAssortment string `json:"productsNotInAssortment"`
	}

	EditAssortmentProductsResponse struct {
		Status                        sharedCommon.Status            `json:"status"`
		EditAssortmentProductsResults []EditAssortmentProductsResult `json:"records"`
	}

	EditAssortmentProductsResponseBulkItem struct {
		Status                        sharedCommon.StatusBulk        `json:"status"`
		EditAssortmentProductsResults []EditAssortmentProductsResult `json:"records"`
	}

	EditAssortmentProductsResponseBulk struct {
		Status    sharedCommon.Status                      `json:"status"`
		BulkItems []EditAssortmentProductsResponseBulkItem `json:"requests"`
	}

	RemoveAssortmentProductResult struct {
		DeletedIDs              string `json:"deletedIDs"`
		ProductsNotInAssortment string `json:"productsNotInAssortment"`
	}

	RemoveAssortmentProductResponse struct {
		Status                         sharedCommon.Status             `json:"status"`
		RemoveAssortmentProductResults []RemoveAssortmentProductResult `json:"records"`
	}

	RemoveAssortmentProductResponseBulkItem struct {
		Status                         sharedCommon.StatusBulk         `json:"status"`
		RemoveAssortmentProductResults []RemoveAssortmentProductResult `json:"records"`
	}

	RemoveAssortmentProductResponseBulk struct {
		Status    sharedCommon.Status                       `json:"status"`
		BulkItems []RemoveAssortmentProductResponseBulkItem `json:"requests"`
	}
)
