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

	GetProductPriorityGroups struct {
		Status  sharedCommon.Status `json:"status"`
		Records []struct {
			PriorityGroupID   int    `json:"priorityGroupID"`
			PriorityGroupName string `json:"priorityGroupName"`
			Added             int    `json:"added"`
			LastModified      int    `json:"lastModified"`
		} `json:"records"`
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
		FileID            int    `json:"fileID"`
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

	//Payload ...
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
)
