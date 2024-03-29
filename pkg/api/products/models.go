package products

import (
	"encoding/json"
	sharedCommon "github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/pkg/errors"
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

	GetProductFilesResponse struct {
		Status       sharedCommon.Status `json:"status"`
		ProductFiles []ProductFile       `json:"records"`
	}

	getProductGroupsResponse struct {
		Status        sharedCommon.Status `json:"status"`
		ProductGroups []ProductGroup      `json:"records"`
	}

	ProductPriorityGroup struct {
		PriorityGroupID   int    `json:"priorityGroupID"`
		PriorityGroupName string `json:"priorityGroupName"`
		Added             int64  `json:"added"`
		LastModified      int64  `json:"lastModified"`
	}

	GetProductPriorityGroups struct {
		Status  sharedCommon.Status    `json:"status"`
		Records []ProductPriorityGroup `json:"records"`
	}

	ProductDimensions struct {
		Name             string `json:"name"`
		Value            string `json:"value"`
		Order            int    `json:"order"`
		DimensionID      int    `json:"dimensionID"`
		DimensionValueID int    `json:"dimensionValueID"`
	}

	VariationDescription struct {
		Name        string `json:"name"`
		Value       string `json:"value"`
		Order       int    `json:"order"`
		DimensionID int    `json:"dimensionID"`
		VariationID int    `json:"variationID"`
	}

	ProductVariaton struct {
		ProductID  string              `json:"productID"`
		Name       string              `json:"name"`
		Code       string              `json:"code"`
		Code2      string              `json:"code2"`
		Dimensions []ProductDimensions `json:"dimensions"`
	}

	ProductPackage struct {
		PackageID          int     `json:"packageID"`
		PackageType        string  `json:"packageType"`
		PackageTypeID      int     `json:"packageTypeID"`
		PackageAmount      float64 `json:"packageAmount"`
		PackageCode        string  `json:"packageCode"`
		PackageNetWeight   float64 `json:"packageNetWeight"`
		PackageGrossWeight float64 `json:"packageGrossWeight"`
		PackageLength      float64 `json:"packageLength"`
		PackageWidth       float64 `json:"packageWidth"`
		PackageHeight      float64 `json:"packageHeight"`
	}

	ProductFile struct {
		FileID            int    `json:"productFileID"`
		Name              string `json:"name"`
		TypeID            int    `json:"typeID"`
		TypeName          string `json:"typeName"`
		IsInformationFile int    `json:"isInformationFile"`
		FileURL           string `json:"fileURL"`
		External          int    `json:"external"`
		HostingProvider   string `json:"hostingProvider"`
	}

	ProductComponent struct {
		ComponentID int     `json:"componentID"`
		Amount      float64 `json:"amount"`
	}

	PriceCalculationStep struct {
		PriceListID   int     `json:"priceListID"`
		PriceListName string  `json:"priceListName"`
		Price         float64 `json:"price"`
		Discount      float64 `json:"discount"`
		Type          string  `json:"type"`
		Percentage    float64 `json:"percentage"`
	}

	// Payload ...
	Product struct {
		ProductID int `json:"productID"`
		Active    int `json:"active"`
		NameLanguages
		Code                         string                 `json:"code"`
		Code2                        string                 `json:"code2"`
		Code3                        *string                `json:"code3"`
		SupplierCode                 string                 `json:"supplierCode"`
		Code5                        *string                `json:"code5"`
		Code6                        *string                `json:"code6"`
		Code7                        *string                `json:"code7"`
		Code8                        *string                `json:"code8"`
		GroupID                      uint                   `json:"groupID"`
		Price                        float64                `json:"price"`
		Cost                         float64                `json:"cost"`
		FifoCost                     float64                `json:"FIFOCost"`
		DisplayedInWebshop           byte                   `json:"displayedInWebshop"`
		BrandID                      uint                   `json:"brandID"`
		Description                  string                 `json:"description"`
		DescriptionLong              string                 `json:"longdesc"`
		DescriptionEng               string                 `json:"descriptionENG"`
		DescriptionLongEng           string                 `json:"longdescENG"`
		DescriptionSpa               string                 `json:"descriptionSPA"`
		DescriptionLongSpa           string                 `json:"longdescSPA"`
		DescriptionEst               string                 `json:"descriptionEST"`
		DescriptionLongEst           string                 `json:"longdescEST"`
		DescriptionGer               string                 `json:"descriptionGER"`
		DescriptionLongGer           string                 `json:"longdescGER"`
		DescriptionSwe               string                 `json:"descriptionSWE"`
		DescriptionLongSwe           string                 `json:"longdescSWE"`
		DescriptionFin               string                 `json:"descriptionFIN"`
		DescriptionLongFin           string                 `json:"longdescFIN"`
		DescriptionRus               string                 `json:"descriptionRUS"`
		DescriptionLongRus           string                 `json:"longdescRUS"`
		DescriptionLat               string                 `json:"descriptionLAT"`
		DescriptionLongLat           string                 `json:"longdescLAT"`
		DescriptionLit               string                 `json:"descriptionLIT"`
		DescriptionLongLit           string                 `json:"longdescLIT"`
		DescriptionGre               string                 `json:"descriptionGRE"`
		DescriptionLongGre           string                 `json:"longdescGRE"`
		AddedByUsername              string                 `json:"addedByUsername"`
		LastModifiedByUsername       string                 `json:"lastModifiedByUsername"`
		Added                        uint64                 `json:"added"`
		LastModified                 uint64                 `json:"lastModified"`
		VatrateID                    uint64                 `json:"vatrateID"`
		Vatrate                      float64                `json:"vatrate"`
		PriceWithVat                 float64                `json:"priceWithVat"`
		BackbarCharges               float64                `json:"backbarCharges"`
		PurchasePrice                float64                `json:"purchasePrice"`
		PriceListPrice               float64                `json:"priceListPrice"`
		PriceListPriceWithVat        float64                `json:"priceListPriceWithVat"`
		GrossWeight                  string                 `json:"grossWeight"`
		NetWeight                    string                 `json:"netWeight"`
		UnitName                     *string                `json:"unitName"`
		BrandName                    string                 `json:"brandName"`
		GroupName                    string                 `json:"groupName"`
		CategoryId                   uint                   `json:"categoryID"`
		UnitID                       uint                   `json:"unitID"`
		CategoryName                 string                 `json:"categoryName"`
		Status                       string                 `json:"status"`
		SupplierID                   int                    `json:"supplierID"`
		HasQuickSelectButton         int                    `json:"hasQuickSelectButton"`
		IsGiftCard                   int                    `json:"isGiftCard"`
		IsRegularGiftCard            int                    `json:"isRegularGiftCard"`
		PriorityGroupID              string                 `json:"priorityGroupID"`
		CountryOfOriginID            string                 `json:"countryOfOriginID"`
		LocationInWarehouseID        string                 `json:"locationInWarehouseID"`
		NonDiscountable              int                    `json:"nonDiscountable"`
		NonRefundable                int                    `json:"nonRefundable"`
		SupplierName                 string                 `json:"supplierName"`
		ManufacturerName             string                 `json:"manufacturerName"`
		Images                       []ProductImage         `json:"images"`
		ProductVariations            []string               `json:"productVariations"` // Variations of matrix product
		ParentProductID              int                    `json:"parentProductID"`
		NonStockProduct              int                    `json:"nonStockProduct"`
		CashierMustEnterPrice        int                    `json:"cashierMustEnterPrice"`
		TaxFree                      int                    `json:"taxFree"`
		ContainerID                  int                    `json:"containerID"`
		Width                        string                 `json:"width"`
		Height                       string                 `json:"height"`
		Length                       string                 `json:"length"`
		LengthInMinutes              int                    `json:"lengthInMinutes"`
		SetupTimeInMinutes           int                    `json:"setupTimeInMinutes"`
		CleanupTimeInMinutes         int                    `json:"cleanupTimeInMinutes"`
		WalkInService                int                    `json:"walkInService"`
		RewardPointsNotAllowed       int                    `json:"rewardPointsNotAllowed"`
		Volume                       string                 `json:"volume"`
		ReorderMultiple              int                    `json:"reorderMultiple"`
		ExtraField1ID                int                    `json:"extraField1ID"`
		ExtraField2ID                int                    `json:"extraField2ID"`
		ExtraField3ID                int                    `json:"extraField3ID"`
		ExtraField4ID                int                    `json:"extraField4ID"`
		Type                         string                 `json:"type"`
		DeliveryTime                 string                 `json:"deliveryTime"`
		ContainerName                string                 `json:"containerName"`
		ContainerCode                string                 `json:"containerCode"`
		ContainerAmount              json.Number            `json:"containerAmount"`
		PackagingType                string                 `json:"packagingType"`
		LocationInWarehouse          string                 `json:"locationInWarehouse"`
		LocationInWarehouseName      string                 `json:"locationInWarehouseName"`
		LocationInWarehouseText      string                 `json:"locationInWarehouseText"`
		ExtraField1Title             string                 `json:"extraField1Title"`
		ExtraField1Code              string                 `json:"extraField1Code"`
		ExtraField1Name              string                 `json:"extraField1Name"`
		ExtraField2Name              string                 `json:"extraField2Name"`
		ExtraField3Title             string                 `json:"extraField3Title"`
		ExtraField2Title             string                 `json:"extraField2Title"`
		ExtraField2Code              string                 `json:"extraField2Code"`
		ExtraField3Code              string                 `json:"extraField3Code"`
		ExtraField3Name              string                 `json:"extraField3Name"`
		ExtraField4Title             string                 `json:"extraField4Title"`
		ExtraField4Code              string                 `json:"extraField4Code"`
		ExtraField4Name              string                 `json:"extraField4Name"`
		SalesPackageClearBrownGlass  string                 `json:"salesPackageClearBrownGlass"`
		SalesPackageGreenOtherGlass  string                 `json:"salesPackageGreenOtherGlass"`
		SalesPackagePlasticPpPe      string                 `json:"salesPackagePlasticPpPe"`
		SalesPackagePlasticPet       string                 `json:"salesPackagePlasticPet"`
		SalesPackageMetalFe          string                 `json:"salesPackageMetalFe"`
		SalesPackageMetalAl          string                 `json:"salesPackageMetalAl"`
		SalesPackageOtherMetal       string                 `json:"salesPackageOtherMetal"`
		SalesPackageCardboard        string                 `json:"salesPackageCardboard"`
		SalesPackageWood             string                 `json:"salesPackageWood"`
		GroupPackagePaper            string                 `json:"groupPackagePaper"`
		GroupPackagePlastic          string                 `json:"groupPackagePlastic"`
		GroupPackageMetal            string                 `json:"groupPackageMetal"`
		GroupPackageWood             string                 `json:"groupPackageWood"`
		TransportPackageWood         string                 `json:"transportPackageWood"`
		TransportPackagePlastic      string                 `json:"transportPackagePlastic"`
		TransportPackageCardboard    string                 `json:"transportPackageCardboar"`
		RegistryNumber               string                 `json:"registryNumber"`
		AlcoholPercentage            string                 `json:"alcoholPercentage"`
		Batches                      string                 `json:"batches"`
		ExciseDeclaration            string                 `json:"exciseDeclaration"`
		ExciseFermentedProductUnder6 string                 `json:"exciseFermentedProductUnder6"`
		ExciseWineOver6              string                 `json:"exciseWineOver6"`
		ExciseFermentedProductOver6  string                 `json:"exciseFermentedProductOver6"`
		ExciseIntermediateProduct    string                 `json:"exciseIntermediateProduct"`
		ExciseOtherAlcohol           string                 `json:"exciseOtherAlcohol"`
		ExcisePackaging              string                 `json:"excisePackaging"`
		Warehouses                   map[uint]StockInfo     `json:"warehouses"`
		Parameters                   []Parameter            `json:"parameters"`
		RelatedProducts              []string               `json:"relatedProducts"`
		VariationList                []ProductVariaton      `json:"variationList"`
		VariationDescriptions        []VariationDescription `json:"variationDescription"`
		ProductPackages              []ProductPackage       `json:"productPackages"`
		ReplacementProducts          []string               `json:"replacementProducts"`
		RelatedFiles                 []ProductFile          `json:"relatedFiles"`
		ProductComponents            []ProductComponent     `json:"productComponents"`
		PriceCalculationSteps        []PriceCalculationStep `json:"priceCalculationSteps"`
		sharedCommon.Attributes
		sharedCommon.LongAttributes
	}

	Image struct {
		Added            int64       `json:"added"`
		External         int         `json:"external"`
		FullURL          string      `json:"fullURL"`
		Hash             *string     `json:"hash"`
		HostingProvider  string      `json:"hostingProvider"`
		LargeURL         string      `json:"largeURL"`
		LastModified     interface{} `json:"lastModified"`
		Name             string      `json:"name"`
		ProductID        int         `json:"productID"`
		ProductPictureID int         `json:"productPictureID"`
		SmallURL         string      `json:"smallURL"`
		Tenant           *string     `json:"tenant"`
		ThumbURL         string      `json:"thumbURL"`
	}

	Option struct {
		ID              int     `json:"optionID"`
		Name            string  `json:"optionName"`
		AdditionalPrice float64 `json:"optionAdditionalPrice"`
	}

	Parameter struct {
		ID      string   `json:"parameterID"`
		Name    string   `json:"parameterName"`
		Type    string   `json:"parameterType"`
		GroupID string   `json:"parameterGroupID"`
		Value   string   `json:"parameterValue"`
		Options []Option `json:"parameterOptions"`
	}

	StockInfo struct {
		WarehouseID   uint        `json:"warehouseID"`
		Free          float64     `json:"free"`
		OrderPending  int         `json:"orderPending"`
		ReorderPoint  int         `json:"reorderPoint"`
		Reserved      json.Number `json:"reserved"`
		TotalInStock  json.Number `json:"totalInStock"`
		RestockLevel  float64     `json:"restockLevel"`
		FifoCost      float32     `json:"FIFOCost"`
		PurchasePrice float32     `json:"purchasePrice"`
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
		ProductCategoryID   int    `json:"productCategoryID"`
		ParentCategoryID    int    `json:"parentCategoryID"`
		ProductCategoryName string `json:"productCategoryName"`
		Added               uint64 `json:"added"`
		LastModified        uint64 `json:"lastModified"`
		sharedCommon.Attributes
	}

	ProductBrand struct {
		ID           uint   `json:"brandID"`
		Name         string `json:"name"`
		Added        uint64 `json:"added"`
		LastModified uint64 `json:"lastModified"`
	}

	VatRateRef struct {
		WarehouseID int `json:"warehouseID"`
		VatrateID   int `json:"vatrateID"`
	}

	ProductGroup struct {
		ID int `json:"productGroupID"`
		NameLanguages
		ShowInWebshop   string         `json:"showInWebshop"`
		NonDiscountable int            `json:"nonDiscountable"`
		PositionNo      int            `json:"positionNo"`
		ParentGroupID   string         `json:"parentGroupID"`
		Added           uint64         `json:"added"`
		LastModified    uint64         `json:"lastModified"`
		SubGroups       []ProductGroup `json:"subGroups"`
		sharedCommon.Attributes
		Images      []ProductGroupImage `json:"images"`
		VatRateRefs []VatRateRef        `json:"vatrates"`
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
	// GetProductUnitsResponse ...
	GetProductUnitsResponse struct {
		Status       sharedCommon.Status `json:"status"`
		ProductUnits []ProductUnit       `json:"records"`
	}

	// ProductUnit ...
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

	GetProductStockResponseBulkItem struct {
		Status          sharedCommon.StatusBulk `json:"status"`
		GetProductStock []GetProductStock       `json:"records"`
	}

	GetProductStockFileResponseBulk struct {
		Status    sharedCommon.Status                   `json:"status"`
		BulkItems []GetProductStockFileResponseBulkItem `json:"requests"`
	}

	GetProductStockResponseBulk struct {
		Status    sharedCommon.Status               `json:"status"`
		BulkItems []GetProductStockResponseBulkItem `json:"requests"`
	}

	GetProductPicturesResponseBulk struct {
		Status    sharedCommon.Status          `json:"status"`
		BulkItems []GetProductPicturesResponse `json:"requests"`
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

	SaveProductCategoryResult struct {
		ProductCategoryID int `json:"productCategoryID"`
	}

	SaveProductCategoryResponse struct {
		Status                     sharedCommon.Status         `json:"status"`
		SaveProductCategoryResults []SaveProductCategoryResult `json:"records"`
	}

	SaveProductCategoryResponseBulkItem struct {
		Status  sharedCommon.StatusBulk     `json:"status"`
		Records []SaveProductCategoryResult `json:"records"`
	}

	SaveProductCategoryResponseBulk struct {
		Status    sharedCommon.Status                   `json:"status"`
		BulkItems []SaveProductCategoryResponseBulkItem `json:"requests"`
	}

	SaveBrandResult struct {
		BrandID int `json:"brandID"`
	}

	SaveBrandResultResponse struct {
		Status           sharedCommon.Status `json:"status"`
		SaveBrandResults []SaveBrandResult   `json:"records"`
	}

	SaveBrandResultResponseBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []SaveBrandResult       `json:"records"`
	}

	SaveBrandResponseBulk struct {
		Status    sharedCommon.Status               `json:"status"`
		BulkItems []SaveBrandResultResponseBulkItem `json:"requests"`
	}

	SaveProductPriorityGroupResult struct {
		PriorityGroupID int `json:"priorityGroupID"`
	}

	SaveProductPriorityGroupResponse struct {
		Status                          sharedCommon.Status              `json:"status"`
		SaveProductPriorityGroupResults []SaveProductPriorityGroupResult `json:"records"`
	}

	SaveProductPriorityGroupBulkItem struct {
		Status  sharedCommon.StatusBulk          `json:"status"`
		Records []SaveProductPriorityGroupResult `json:"records"`
	}

	SaveProductPriorityGroupResponseBulk struct {
		Status    sharedCommon.Status                `json:"status"`
		BulkItems []SaveProductPriorityGroupBulkItem `json:"requests"`
	}

	SaveProductGroupResult struct {
		ProductGroupID int `json:"productGroupID"`
	}

	SaveProductGroupResponse struct {
		Status                  sharedCommon.Status      `json:"status"`
		SaveProductGroupResults []SaveProductGroupResult `json:"records"`
	}

	SaveProductGroupBulkItem struct {
		Status  sharedCommon.StatusBulk  `json:"status"`
		Records []SaveProductGroupResult `json:"records"`
	}

	SaveProductGroupResponseBulk struct {
		Status    sharedCommon.Status        `json:"status"`
		BulkItems []SaveProductGroupBulkItem `json:"requests"`
	}

	DeleteProductGroupResponse struct {
		Status sharedCommon.Status `json:"status"`
	}

	DeleteProductGroupResponseBulkItem struct {
		Status sharedCommon.StatusBulk `json:"status"`
	}

	DeleteProductGroupResponseBulk struct {
		Status    sharedCommon.Status                  `json:"status"`
		BulkItems []DeleteProductGroupResponseBulkItem `json:"requests"`
	}

	GetProductPriorityGroupBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []ProductPriorityGroup  `json:"records"`
	}

	GetProductPriorityGroupResponseBulk struct {
		Status    sharedCommon.Status               `json:"status"`
		BulkItems []GetProductPriorityGroupBulkItem `json:"requests"`
	}

	GetProductCategoryBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []ProductCategory       `json:"records"`
	}

	GetProductCategoryResponseBulk struct {
		Status    sharedCommon.Status          `json:"status"`
		BulkItems []GetProductCategoryBulkItem `json:"requests"`
	}

	GetProductGroupBulkItem struct {
		Status  sharedCommon.StatusBulk `json:"status"`
		Records []ProductGroup          `json:"records"`
	}

	GetProductGroupResponseBulk struct {
		Status    sharedCommon.Status       `json:"status"`
		BulkItems []GetProductGroupBulkItem `json:"requests"`
	}

	GetProductPicturesResponse struct {
		Records []Image             `json:"records"`
		Status  sharedCommon.Status `json:"status"`
	}
)

func (u *Product) UnmarshalJSON(data []byte) error {

	raw := struct {
		ProductID int `json:"productID"`
		Active    int `json:"active"`
		NameLanguages
		Code                         string                 `json:"code"`
		Code2                        string                 `json:"code2"`
		Code3                        *string                `json:"code3"`
		SupplierCode                 string                 `json:"supplierCode"`
		Code5                        *string                `json:"code5"`
		Code6                        *string                `json:"code6"`
		Code7                        *string                `json:"code7"`
		Code8                        *string                `json:"code8"`
		GroupID                      uint                   `json:"groupID"`
		Price                        float64                `json:"price"`
		Cost                         float64                `json:"cost"`
		FifoCost                     float64                `json:"FIFOCost"`
		DisplayedInWebshop           byte                   `json:"displayedInWebshop"`
		BrandID                      uint                   `json:"brandID"`
		Description                  string                 `json:"description"`
		DescriptionLong              string                 `json:"longdesc"`
		DescriptionEng               string                 `json:"descriptionENG"`
		DescriptionLongEng           string                 `json:"longdescENG"`
		DescriptionSpa               string                 `json:"descriptionSPA"`
		DescriptionLongSpa           string                 `json:"longdescSPA"`
		DescriptionEst               string                 `json:"descriptionEST"`
		DescriptionLongEst           string                 `json:"longdescEST"`
		DescriptionGer               string                 `json:"descriptionGER"`
		DescriptionLongGer           string                 `json:"longdescGER"`
		DescriptionSwe               string                 `json:"descriptionSWE"`
		DescriptionLongSwe           string                 `json:"longdescSWE"`
		DescriptionFin               string                 `json:"descriptionFIN"`
		DescriptionLongFin           string                 `json:"longdescFIN"`
		DescriptionRus               string                 `json:"descriptionRUS"`
		DescriptionLongRus           string                 `json:"longdescRUS"`
		DescriptionLat               string                 `json:"descriptionLAT"`
		DescriptionLongLat           string                 `json:"longdescLAT"`
		DescriptionLit               string                 `json:"descriptionLIT"`
		DescriptionLongLit           string                 `json:"longdescLIT"`
		DescriptionGre               string                 `json:"descriptionGRE"`
		DescriptionLongGre           string                 `json:"longdescGRE"`
		AddedByUsername              string                 `json:"addedByUsername"`
		LastModifiedByUsername       string                 `json:"lastModifiedByUsername"`
		Added                        uint64                 `json:"added"`
		LastModified                 uint64                 `json:"lastModified"`
		VatrateID                    uint64                 `json:"vatrateID"`
		Vatrate                      float64                `json:"vatrate"`
		PriceWithVat                 float64                `json:"priceWithVat"`
		BackbarCharges               float64                `json:"backbarCharges"`
		PurchasePrice                float64                `json:"purchasePrice"`
		PriceListPrice               json.Number            `json:"priceListPrice"`
		PriceListPriceWithVat        json.Number            `json:"priceListPriceWithVat"`
		GrossWeight                  string                 `json:"grossWeight"`
		NetWeight                    string                 `json:"netWeight"`
		UnitName                     *string                `json:"unitName"`
		BrandName                    string                 `json:"brandName"`
		GroupName                    string                 `json:"groupName"`
		CategoryId                   uint                   `json:"categoryID"`
		UnitID                       uint                   `json:"unitID"`
		CategoryName                 string                 `json:"categoryName"`
		Status                       string                 `json:"status"`
		SupplierID                   int                    `json:"supplierID"`
		HasQuickSelectButton         int                    `json:"hasQuickSelectButton"`
		IsGiftCard                   int                    `json:"isGiftCard"`
		IsRegularGiftCard            int                    `json:"isRegularGiftCard"`
		PriorityGroupID              string                 `json:"priorityGroupID"`
		CountryOfOriginID            string                 `json:"countryOfOriginID"`
		LocationInWarehouseID        string                 `json:"locationInWarehouseID"`
		NonDiscountable              int                    `json:"nonDiscountable"`
		NonRefundable                int                    `json:"nonRefundable"`
		SupplierName                 string                 `json:"supplierName"`
		ManufacturerName             string                 `json:"manufacturerName"`
		Images                       []ProductImage         `json:"images"`
		ProductVariations            []string               `json:"productVariations"` // Variations of matrix product
		ParentProductID              int                    `json:"parentProductID"`
		NonStockProduct              int                    `json:"nonStockProduct"`
		CashierMustEnterPrice        int                    `json:"cashierMustEnterPrice"`
		TaxFree                      int                    `json:"taxFree"`
		ContainerID                  int                    `json:"containerID"`
		Width                        string                 `json:"width"`
		Height                       string                 `json:"height"`
		Length                       string                 `json:"length"`
		LengthInMinutes              int                    `json:"lengthInMinutes"`
		SetupTimeInMinutes           int                    `json:"setupTimeInMinutes"`
		CleanupTimeInMinutes         int                    `json:"cleanupTimeInMinutes"`
		WalkInService                int                    `json:"walkInService"`
		RewardPointsNotAllowed       int                    `json:"rewardPointsNotAllowed"`
		Volume                       string                 `json:"volume"`
		ReorderMultiple              int                    `json:"reorderMultiple"`
		ExtraField1ID                int                    `json:"extraField1ID"`
		ExtraField2ID                int                    `json:"extraField2ID"`
		ExtraField3ID                int                    `json:"extraField3ID"`
		ExtraField4ID                int                    `json:"extraField4ID"`
		Type                         string                 `json:"type"`
		DeliveryTime                 string                 `json:"deliveryTime"`
		ContainerName                string                 `json:"containerName"`
		ContainerCode                string                 `json:"containerCode"`
		ContainerAmount              json.Number            `json:"containerAmount"`
		PackagingType                string                 `json:"packagingType"`
		LocationInWarehouse          string                 `json:"locationInWarehouse"`
		LocationInWarehouseName      string                 `json:"locationInWarehouseName"`
		LocationInWarehouseText      string                 `json:"locationInWarehouseText"`
		ExtraField1Title             string                 `json:"extraField1Title"`
		ExtraField1Code              string                 `json:"extraField1Code"`
		ExtraField1Name              string                 `json:"extraField1Name"`
		ExtraField2Name              string                 `json:"extraField2Name"`
		ExtraField3Title             string                 `json:"extraField3Title"`
		ExtraField2Title             string                 `json:"extraField2Title"`
		ExtraField2Code              string                 `json:"extraField2Code"`
		ExtraField3Code              string                 `json:"extraField3Code"`
		ExtraField3Name              string                 `json:"extraField3Name"`
		ExtraField4Title             string                 `json:"extraField4Title"`
		ExtraField4Code              string                 `json:"extraField4Code"`
		ExtraField4Name              string                 `json:"extraField4Name"`
		SalesPackageClearBrownGlass  string                 `json:"salesPackageClearBrownGlass"`
		SalesPackageGreenOtherGlass  string                 `json:"salesPackageGreenOtherGlass"`
		SalesPackagePlasticPpPe      string                 `json:"salesPackagePlasticPpPe"`
		SalesPackagePlasticPet       string                 `json:"salesPackagePlasticPet"`
		SalesPackageMetalFe          string                 `json:"salesPackageMetalFe"`
		SalesPackageMetalAl          string                 `json:"salesPackageMetalAl"`
		SalesPackageOtherMetal       string                 `json:"salesPackageOtherMetal"`
		SalesPackageCardboard        string                 `json:"salesPackageCardboard"`
		SalesPackageWood             string                 `json:"salesPackageWood"`
		GroupPackagePaper            string                 `json:"groupPackagePaper"`
		GroupPackagePlastic          string                 `json:"groupPackagePlastic"`
		GroupPackageMetal            string                 `json:"groupPackageMetal"`
		GroupPackageWood             string                 `json:"groupPackageWood"`
		TransportPackageWood         string                 `json:"transportPackageWood"`
		TransportPackagePlastic      string                 `json:"transportPackagePlastic"`
		TransportPackageCardboard    string                 `json:"transportPackageCardboar"`
		RegistryNumber               string                 `json:"registryNumber"`
		AlcoholPercentage            string                 `json:"alcoholPercentage"`
		Batches                      string                 `json:"batches"`
		ExciseDeclaration            string                 `json:"exciseDeclaration"`
		ExciseFermentedProductUnder6 string                 `json:"exciseFermentedProductUnder6"`
		ExciseWineOver6              string                 `json:"exciseWineOver6"`
		ExciseFermentedProductOver6  string                 `json:"exciseFermentedProductOver6"`
		ExciseIntermediateProduct    string                 `json:"exciseIntermediateProduct"`
		ExciseOtherAlcohol           string                 `json:"exciseOtherAlcohol"`
		ExcisePackaging              string                 `json:"excisePackaging"`
		Warehouses                   map[uint]StockInfo     `json:"warehouses"`
		Parameters                   []Parameter            `json:"parameters"`
		RelatedProducts              []string               `json:"relatedProducts"`
		VariationList                []ProductVariaton      `json:"variationList"`
		VariationDescriptions        []VariationDescription `json:"variationDescription"`
		ProductPackages              []ProductPackage       `json:"productPackages"`
		ReplacementProducts          []string               `json:"replacementProducts"`
		RelatedFiles                 []ProductFile          `json:"relatedFiles"`
		ProductComponents            []ProductComponent     `json:"productComponents"`
		PriceCalculationSteps        []PriceCalculationStep `json:"priceCalculationSteps"`
		sharedCommon.Attributes
		sharedCommon.LongAttributes
	}{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	u.ProductID = raw.ProductID
	u.Active = raw.Active
	u.NameLanguages = raw.NameLanguages
	u.Code = raw.Code
	u.Code2 = raw.Code2
	u.Code3 = raw.Code3
	u.SupplierCode = raw.SupplierCode
	u.Code5 = raw.Code5
	u.Code6 = raw.Code6
	u.Code7 = raw.Code7
	u.Code8 = raw.Code8
	u.GroupID = raw.GroupID
	u.Price = raw.Price
	u.Cost = raw.Cost
	u.FifoCost = raw.FifoCost
	u.DisplayedInWebshop = raw.DisplayedInWebshop
	u.BrandID = raw.BrandID
	u.Description = raw.Description
	u.DescriptionLong = raw.DescriptionLong
	u.DescriptionEng = raw.DescriptionEng
	u.DescriptionLongEng = raw.DescriptionLongEng
	u.DescriptionSpa = raw.DescriptionSpa
	u.DescriptionLongSpa = raw.DescriptionLongSpa
	u.DescriptionEst = raw.DescriptionEst
	u.DescriptionLongEst = raw.DescriptionLongEst
	u.DescriptionGer = raw.DescriptionGer
	u.DescriptionLongGer = raw.DescriptionLongGer
	u.DescriptionSwe = raw.DescriptionSwe
	u.DescriptionLongSwe = raw.DescriptionLongSwe
	u.DescriptionFin = raw.DescriptionFin
	u.DescriptionLongFin = raw.DescriptionLongFin
	u.DescriptionRus = raw.DescriptionRus
	u.DescriptionLongRus = raw.DescriptionLongRus
	u.DescriptionLat = raw.DescriptionLat
	u.DescriptionLongLat = raw.DescriptionLongLat
	u.DescriptionLit = raw.DescriptionLit
	u.DescriptionLongLit = raw.DescriptionLongLit
	u.DescriptionGre = raw.DescriptionGre
	u.DescriptionLongGre = raw.DescriptionLongGre
	u.AddedByUsername = raw.AddedByUsername
	u.LastModifiedByUsername = raw.LastModifiedByUsername
	u.Added = raw.Added
	u.LastModified = raw.LastModified
	u.VatrateID = raw.VatrateID
	u.Vatrate = raw.Vatrate
	u.PriceWithVat = raw.PriceWithVat
	u.BackbarCharges = raw.BackbarCharges
	u.PurchasePrice = raw.PurchasePrice
	if raw.PriceListPrice.String() != "" {
		value, err := raw.PriceListPrice.Float64()
		if err != nil {
			return errors.Wrapf(err, "unable to unmarshal address. priceListPrice did not contain a float: %s", raw.PriceListPrice.String())
		}
		u.PriceListPrice = value
	}
	if raw.PriceListPriceWithVat.String() != "" {
		value, err := raw.PriceListPriceWithVat.Float64()
		if err != nil {
			return errors.Wrapf(err, "unable to unmarshal address. priceListPriceWithVat did not contain a float: %s", raw.PriceListPriceWithVat.String())
		}
		u.PriceListPriceWithVat = value
	}
	u.GrossWeight = raw.GrossWeight
	u.NetWeight = raw.NetWeight
	u.UnitName = raw.UnitName
	u.BrandName = raw.BrandName
	u.GroupName = raw.GroupName
	u.CategoryId = raw.CategoryId
	u.UnitID = raw.UnitID
	u.CategoryName = raw.CategoryName
	u.Status = raw.Status
	u.SupplierID = raw.SupplierID
	u.HasQuickSelectButton = raw.HasQuickSelectButton
	u.IsGiftCard = raw.IsGiftCard
	u.IsRegularGiftCard = raw.IsRegularGiftCard
	u.PriorityGroupID = raw.PriorityGroupID
	u.CountryOfOriginID = raw.CountryOfOriginID
	u.LocationInWarehouseID = raw.LocationInWarehouseID
	u.NonDiscountable = raw.NonDiscountable
	u.NonRefundable = raw.NonRefundable
	u.SupplierName = raw.SupplierName
	u.ManufacturerName = raw.ManufacturerName
	u.Images = raw.Images
	u.ProductVariations = raw.ProductVariations
	u.ParentProductID = raw.ParentProductID
	u.NonStockProduct = raw.NonStockProduct
	u.CashierMustEnterPrice = raw.CashierMustEnterPrice
	u.TaxFree = raw.TaxFree
	u.ContainerID = raw.ContainerID
	u.Width = raw.Width
	u.Height = raw.Height
	u.Length = raw.Length
	u.LengthInMinutes = raw.LengthInMinutes
	u.SetupTimeInMinutes = raw.SetupTimeInMinutes
	u.CleanupTimeInMinutes = raw.CleanupTimeInMinutes
	u.WalkInService = raw.WalkInService
	u.RewardPointsNotAllowed = raw.RewardPointsNotAllowed
	u.Volume = raw.Volume
	u.ReorderMultiple = raw.ReorderMultiple
	u.ExtraField1ID = raw.ExtraField1ID
	u.ExtraField2ID = raw.ExtraField2ID
	u.ExtraField3ID = raw.ExtraField3ID
	u.ExtraField4ID = raw.ExtraField4ID
	u.Type = raw.Type
	u.DeliveryTime = raw.DeliveryTime
	u.ContainerName = raw.ContainerName
	u.ContainerCode = raw.ContainerCode
	u.ContainerAmount = raw.ContainerAmount
	u.PackagingType = raw.PackagingType
	u.LocationInWarehouse = raw.LocationInWarehouse
	u.LocationInWarehouseName = raw.LocationInWarehouseName
	u.LocationInWarehouseText = raw.LocationInWarehouseText
	u.ExtraField1Title = raw.ExtraField1Title
	u.ExtraField1Code = raw.ExtraField1Code
	u.ExtraField1Name = raw.ExtraField1Name
	u.ExtraField2Name = raw.ExtraField2Name
	u.ExtraField3Title = raw.ExtraField3Title
	u.ExtraField2Title = raw.ExtraField2Title
	u.ExtraField2Code = raw.ExtraField2Code
	u.ExtraField3Code = raw.ExtraField3Code
	u.ExtraField3Name = raw.ExtraField3Name
	u.ExtraField4Title = raw.ExtraField4Title
	u.ExtraField4Code = raw.ExtraField4Code
	u.ExtraField4Name = raw.ExtraField4Name
	u.SalesPackageClearBrownGlass = raw.SalesPackageClearBrownGlass
	u.SalesPackageGreenOtherGlass = raw.SalesPackageGreenOtherGlass
	u.SalesPackagePlasticPpPe = raw.SalesPackagePlasticPpPe
	u.SalesPackagePlasticPet = raw.SalesPackagePlasticPet
	u.SalesPackageMetalFe = raw.SalesPackageMetalFe
	u.SalesPackageMetalAl = raw.SalesPackageMetalAl
	u.SalesPackageOtherMetal = raw.SalesPackageOtherMetal
	u.SalesPackageCardboard = raw.SalesPackageCardboard
	u.SalesPackageWood = raw.SalesPackageWood
	u.GroupPackagePaper = raw.GroupPackagePaper
	u.GroupPackagePlastic = raw.GroupPackagePlastic
	u.GroupPackageMetal = raw.GroupPackageMetal
	u.GroupPackageWood = raw.GroupPackageWood
	u.TransportPackageWood = raw.TransportPackageWood
	u.TransportPackagePlastic = raw.TransportPackagePlastic
	u.TransportPackageCardboard = raw.TransportPackageCardboard
	u.RegistryNumber = raw.RegistryNumber
	u.AlcoholPercentage = raw.AlcoholPercentage
	u.Batches = raw.Batches
	u.ExciseDeclaration = raw.ExciseDeclaration
	u.ExciseFermentedProductUnder6 = raw.ExciseFermentedProductUnder6
	u.ExciseWineOver6 = raw.ExciseWineOver6
	u.ExciseFermentedProductOver6 = raw.ExciseFermentedProductOver6
	u.ExciseIntermediateProduct = raw.ExciseIntermediateProduct
	u.ExciseOtherAlcohol = raw.ExciseOtherAlcohol
	u.ExcisePackaging = raw.ExcisePackaging
	u.Warehouses = raw.Warehouses
	u.Parameters = raw.Parameters
	u.RelatedProducts = raw.RelatedProducts
	u.VariationList = raw.VariationList
	u.VariationDescriptions = raw.VariationDescriptions
	u.ProductPackages = raw.ProductPackages
	u.ReplacementProducts = raw.ReplacementProducts
	u.RelatedFiles = raw.RelatedFiles
	u.ProductComponents = raw.ProductComponents
	u.PriceCalculationSteps = raw.PriceCalculationSteps
	u.Attributes = raw.Attributes
	u.LongAttributes = raw.LongAttributes
	return nil
}
