package common

import (
	"fmt"
	"github.com/pkg/errors"
)

type ApiError int

const (
	ServerMaintenance                                        ApiError = 1000
	AccountNotFound                                          ApiError = 1001
	HourlyRequestQuota                                       ApiError = 1002
	AccountDbConnError                                       ApiError = 1003
	UnknownApi                                               ApiError = 1005
	ApiNotAvailable                                          ApiError = 1006
	UnknownOutputFormat                                      ApiError = 1007
	DbError                                                  ApiError = 1008
	MissingAuth                                              ApiError = 1009
	RequiredParamMissing                                     ApiError = 1010
	InvalidClassifierID                                      ApiError = 1011
	ParamIsNotUnique                                         ApiError = 1012
	InconsistentParam                                        ApiError = 1013
	InvalidFormat                                            ApiError = 1014
	MalformedRequest                                         ApiError = 1015
	InvalidValue                                             ApiError = 1016
	IsAlreadyConfirmed                                       ApiError = 1017
	MultipleMatchesFound                                     ApiError = 1018
	NoRecordsFound                                           ApiError = 1019
	TooManyBulkSubRequests                                   ApiError = 1020
	SameInstanceIsRunning                                    ApiError = 1021
	RelatedDeletionError                                     ApiError = 1022
	WrongRowsSequence                                        ApiError = 1023
	MasterListLimitation                                     ApiError = 1024
	BinNotEmpty                                              ApiError = 1025
	IdenticalRecordExists                                    ApiError = 1026
	NotAllowedToChangeValue                                  ApiError = 1027
	NotUsableField                                           ApiError = 1028
	IncorrectList                                            ApiError = 1029
	ArrayValueRequired                                       ApiError = 1030
	MissingCouponError                                       ApiError = 1040
	CouponAlredyUsedError                                    ApiError = 1041
	NotEnoughRewardPoints                                    ApiError = 1042
	AppointmentBusyError                                     ApiError = 1043
	NotPossibleTimeSlots                                     ApiError = 1044
	CouponExpired                                            ApiError = 1045
	ConflictingRequirements                                  ApiError = 1046
	NoPromotionRequirements                                  ApiError = 1047
	ConflictingAwards                                        ApiError = 1048
	AwardsNotSpecified                                       ApiError = 1049
	AuthMissing                                              ApiError = 1050
	LoginFailed                                              ApiError = 1051
	UserBlocked                                              ApiError = 1052
	MissingSavedPassword                                     ApiError = 1053
	APISessionExpired                                        ApiError = 1054
	InvalidSession                                           ApiError = 1055
	SessionTooOld                                            ApiError = 1056
	DemoAccountExpired                                       ApiError = 1057
	PinLoginNotSupported                                     ApiError = 1058
	NoUserGroupDetected                                      ApiError = 1059
	NoViewingRights                                          ApiError = 1060
	NoAddingRights                                           ApiError = 1061
	NoEditingRights                                          ApiError = 1062
	NoDeletingRights                                         ApiError = 1063
	NoLocationAccess                                         ApiError = 1064
	NoAPIAccess                                              ApiError = 1065
	NoGroupManagementRights                                  ApiError = 1066
	WrongAccountFranchise                                    ApiError = 1067
	AccountNotConfirmed                                      ApiError = 1068
	BuyUpfrontOnly                                           ApiError = 1071
	NoPointsEarned                                           ApiError = 1072
	InconsistentDocumentsForInvoice                          ApiError = 1073
	WrongSourceDocument                                      ApiError = 1074
	MultiComponentsTax                                       ApiError = 1075
	WrongSalesPromotionType                                  ApiError = 1076
	WrongPriceListAssociation                                ApiError = 1077
	AmountFieldError                                         ApiError = 1078
	ProductIDChangeFailure                                   ApiError = 1079
	PrintingServiceFailure                                   ApiError = 1080
	EmailSendingFailure                                      ApiError = 1081
	EmailSettingFailure                                      ApiError = 1082
	WrongMasterListSetup                                     ApiError = 1083
	WrongMasterListUniqueField                               ApiError = 1084
	NoFileAttached                                           ApiError = 1090
	WrongFileEncoding                                        ApiError = 1091
	TooBigFile                                               ApiError = 1092
	PasswordLengthFailure                                    ApiError = 1100
	WrongLettersInPassword                                   ApiError = 1101
	PasswordComplexityError                                  ApiError = 1102
	PasswordChangeLimitFailure                               ApiError = 1103
	MultipleConflictingSettingsInSalesPromotion              ApiError = 1110
	SalesPromotionPurchasedAmountConflict                    ApiError = 1111
	SalesPromotionMultipleConflictingPurchaseOptions         ApiError = 1112
	SalesPromotionAwardedProductWithSumOffConflict           ApiError = 1113
	SalesPromotionAwardedProductConflict                     ApiError = 1114
	SalesPromotionPercentageOffConflict                      ApiError = 1115
	SalesPromotionSumOffConflict                             ApiError = 1116
	SalesPromotionPriceWithPurchasedAmountConflict           ApiError = 1117
	SalesPromotionMaxPointsDiscountWithRewardPointsConflict  ApiError = 1118
	SalesPromotionLowestPriceWithSumOffConflict              ApiError = 1119
	CustomerRegistryServiceUsed                              ApiError = 1120
	CannotChangeTypeOfDocument                               ApiError = 1121
	SalesPromotionSpecialPriceWithPurchasedAmountConflict    ApiError = 1122
	SalesPromotionPercentageOffWithPurchasedAmountConflict   ApiError = 1123
	SalesDocumentAccountNeedsUpdate                          ApiError = 1124
	AccountLimitationOnIntegration                           ApiError = 1126
	AccountInputFieldIsNotAllowed                            ApiError = 1127
	GreekAccountsOnly                                        ApiError = 1128
	OnlyOneValueForSalesPromotion                            ApiError = 1129
	WrongBillingAssociation                                  ApiError = 1130
	SalesPromotionWrongReasonCode                            ApiError = 1131
	SalesPromotionPurchasedProductWithSumOffConflict         ApiError = 1132
	SalesPromotionPurchasedProductSameAmountError            ApiError = 1133
	SalesPromotionAwardedProductAmountError                  ApiError = 1134
	NoAssortmentPossibleForLocation                          ApiError = 1136
	TooManyProducts                                          ApiError = 1137
	TooManyBillingIds                                        ApiError = 1138
	SalesPromotionSpecialUnitPurchasedAmountConflict         ApiError = 1139
	SalesPromotionMaxItemsBiggerThanPurchasedAmount          ApiError = 1140
	SalesPromotionPurchasedAmountMissingPurchasedProductData ApiError = 1141
	ExternalCouponRegistryService                            ApiError = 1142
	ExternalCouponRegistryServiceWrongInputField             ApiError = 1143
	SalesPromotionRedemptionLimitTooBig                      ApiError = 1144
	SalesPromotionRedemptionLimitWithMaxItemsConflict        ApiError = 1145
	NoEmployeeRecord                                         ApiError = 1146
	ComplianceAlreadyConfirmed                               ApiError = 1147
	NoAccessToCustomerData                                   ApiError = 1148
	NonEuCountry                                             ApiError = 1149
	IntegrationSpecificFieldMissingInputParameter            ApiError = 1150
	ProductAlreadyAppearsInSpecialPriceList                  ApiError = 1151
	PriceListOverlapsWithSpecialPriceList                    ApiError = 1152
	WrongPurposeForReasonCode                                ApiError = 1153
	OldInventoryModule                                       ApiError = 1154
	CreateAccountError                                       ApiError = 1155
	ValueLengthError                                         ApiError = 1156
	BulkSubRequestDocumentError                              ApiError = 1157
	BulkSubRequestDuplicateError                             ApiError = 1158
	EuAccountsField                                          ApiError = 1159
	GiftCardVatRateWithPaymentTypeConflict                   ApiError = 1160
	PosAppRequestType                                        ApiError = 1161
	TooLongListOfElements                                    ApiError = 1162
	CustomerCodeMissingInJWT                                 ApiError = 1170
	NoRightsForBackOffice                                    ApiError = 1171
	MissingUsernameInJWT                                     ApiError = 1172
	MissingUserName                                          ApiError = 1173
	PendingStatusForUser                                     ApiError = 1174
	UsernameAlreadyExists                                    ApiError = 1175
	NotPossibleToExtendSessionForJWT                         ApiError = 1176
	SignupNotAllowedForCountry                               ApiError = 1177
	ContactPersonBelongsToAnotherContact                     ApiError = 1178
	DateIsNotFutureError                                     ApiError = 1179
	WrongVersionNumberForInvoice                             ApiError = 1180
	StockTakingIsConfirmed                                   ApiError = 1181
	SalesPromotionFlagExcludePromotionWrongValue             ApiError = 1182
	CDNIntegrationRequired                                   ApiError = 1183
	SalesPromotionMaxNrOfMatchingItemsConflictingValue       ApiError = 1184
	SalesPromotionMaxNrOfMatchingItemsWrongValue             ApiError = 1185
	StocktakingAlreadyConnectedToDocumentType                ApiError = 1186
	ConfigurationCallIsMissing                               ApiError = 1187
	ConfigurationCallIsInvalid                               ApiError = 1188
	WarehouseNoLocalQuickButtonEnabled                       ApiError = 1189
	WrongJWTAccount                                          ApiError = 1190
	JWTDecodingFailure                                       ApiError = 1191
	JWTExpired                                               ApiError = 1194
	WrongLanguageCode                                        ApiError = 1195
	NotNewPassword                                           ApiError = 1196
)

var errorsMap map[int]string

func init() {
	errorsMap = map[int]string{
		int(ServerMaintenance):                                `API is under maintenance, please try again in a couple of minutes.`,
		int(AccountNotFound):                                  `Account not found. (It may also mean that input parameter "clientCode" is missing.)`,
		int(HourlyRequestQuota):                               `Hourly request quota (by default 2000 requests) has been exceeded for this account. Please resume next hour.`,
		int(AccountDbConnError):                               `Cannot connect to account database.`,
		int(UnknownApi):                                       `API call name (input parameter "request") not specified, or unknown API call.`,
		int(ApiNotAvailable):                                  `This API call is not available on this account. (Account needs upgrading, or an extra module needs to be installed.)`,
		int(UnknownOutputFormat):                              `Unknown output format requested; input parameter "responseType" must be set to either "JSON" or "XML".`,
		int(DbError):                                          `Either a) database is under regular maintenance (please try again in a couple of minutes), or b) your application is not connecting to the correct API server. Make sure that you are using correct API URL: https://YOURCUSTOMERCODE.erply.com/api/. If your API URL is correct, it might be that your ERPLY account has been recently transferred between hosting environments and your local DNS cache is out of date (domain name YOURCUSTOMERCODE.erply.com is not being resolved to correct web server). Try flushing DNS cache in your computer, server, or application engine.`,
		int(MissingAuth):                                      `This API call requires authentication parameters (a session key, authentication key, or service key), but none were found.`,
		int(RequiredParamMissing):                             `Required parameters are missing. (Attribute "errorField" indicates the missing input parameter.)`,
		int(InvalidClassifierID):                              `Invalid classifier ID, there is no such item. (Attribute "errorField" indicates the invalid input parameter.)`,
		int(ParamIsNotUnique):                                 `A parameter must have a unique value. (Attribute "errorField" indicates the invalid input parameter.)`,
		int(InconsistentParam):                                `Inconsistent parameter set (for example, both product and service IDs specified for an invoice row).`,
		int(InvalidFormat):                                    `Incorrect data type or format. (Attribute "errorField" indicates the invalid input parameter.)`,
		int(MalformedRequest):                                 `Malformed request (eg.parameters containing invalid characters).`,
		int(InvalidValue):                                     `Invalid value.(Attribute "errorField" indicates the field that contains an invalid value.)`,
		int(IsAlreadyConfirmed):                               `Document has been confirmed and its contents and warehouse ID cannot be edited any more.`,
		int(MultipleMatchesFound):                             `Multiple matches found, all have the same attribute value.No records will be updated.`,
		int(NoRecordsFound):                                   `No records found with this attribute value.`,
		int(TooManyBulkSubRequests):                           `Bulk API call contained more than 100 sub-requests (max 100 allowed).The whole request has been ignored.`,
		int(SameInstanceIsRunning):                            `Another instance of the same report is currently running.Please wait and try again in a minute.(For long-running reports, API processes incoming requests only one at a time.)`,
		int(RelatedDeletionError):                             `This item cannot be deleted because there are other records that reference it.`,
		int(WrongRowsSequence):                                `Request has product rows in wrong order, or some rows are missing.When editing a confirmed Inventory Registration, only prices can be updated (not quantities and product IDs), and the request must include all rows.`,
		int(MasterListLimitation):                             `"Master List" functionality has been activated - products cannot be added directly to the product catalog.`,
		int(BinNotEmpty):                                      `This bin cannot be archived because it has quantities in it.`,
		int(IdenticalRecordExists):                            `An identical record already exists.`,
		int(NotAllowedToChangeValue):                          `On an existing record, it is not allowed to change the value of this field.`,
		int(NotUsableField):                                   `This input field in this API call cannot be used.(Account needs upgrading, or an extra module needs to be installed.)`,
		int(IncorrectList):                                    `One or more values in a comma-separated list are incorrect.`,
		int(ArrayValueRequired):                               `Input parameter must not be an array.(Attribute "errorField" indicates the invalid input parameter.)`,
		int(MissingCouponError):                               `Invalid coupon identifier - such coupon has not been issued.`,
		int(CouponAlredyUsedError):                            `Invalid coupon identifier - this coupon has already been redeemed.`,
		int(NotEnoughRewardPoints):                            `Customer does not have enough reward points.`,
		int(AppointmentBusyError):                             `Employee already has an appointment on that time slot.Please choose a different start and end time for appointment.`,
		int(NotPossibleTimeSlots):                             `Default length for this service has not been defined in Erply backend - cannot suggest possible time slots.`,
		int(CouponExpired):                                    `Invalid coupon identifier - this coupon has expired.`,
		int(ConflictingRequirements):                          `Sales Promotion - The promotion contains multiple conflicting requirements or conditions, please specify only one.`,
		int(NoPromotionRequirements):                          `Sales Promotion - Promotion requirements or conditions not specified.`,
		int(ConflictingAwards):                                `Sales Promotion - The promotion contains multiple conflicting awards, please specify only one.`,
		int(AwardsNotSpecified):                               `Sales Promotion - Promotion awards not specified.`,
		int(AuthMissing):                                      `Username/password missing.`,
		int(LoginFailed):                                      `Login failed.`,
		int(UserBlocked):                                      `User has been temporarily blocked because of repeated unsuccessful login attempts.`,
		int(MissingSavedPassword):                             `No password has been set for this user, therefore the user cannot be logged in.`,
		int(APISessionExpired):                                `API session has expired. Please call API "verifyUser" again (with correct credentials) to receive a new session key.`,
		int(InvalidSession):                                   `Supplied session key is invalid; session not found.`,
		int(SessionTooOld):                                    `Supplied session key is too old. User switching is no longer possible with this session key, please perform a full re-authentication via API "verifyUser".`,
		int(DemoAccountExpired):                               `Your time-limited demo account has expired. Please create a new ERPLY demo account, or sign up for a paid account.`,
		int(PinLoginNotSupported):                             `PIN login is not supported. Provide a user name and password instead, or use the "switchUser" API call.`,
		int(NoUserGroupDetected):                              `Unable to detect your user group.`,
		int(NoViewingRights):                                  `No viewing rights (in this module/for this item).`,
		int(NoAddingRights):                                   `No adding rights (in this module).`,
		int(NoEditingRights):                                  `No editing rights (in this module/for this item).`,
		int(NoDeletingRights):                                 `No deleting rights (in this module/for this item).`,
		int(NoLocationAccess):                                 `User does not have access to this location (store, warehouse).`,
		int(NoAPIAccess):                                      `This user account does not have API access. (It may be limited to POS or Erply backend operations only.)`,
		int(NoGroupManagementRights):                          `This user does not have the right to manage specified user group. (Error may occur when attempting to create a new user, or modify an existing one.)`,
		int(WrongAccountFranchise):                            `This account does not belong to a franchise and this API call cannot be used.`,
		int(AccountNotConfirmed):                              `This user cannot yet log in to Erply. A confirmation email has been sent; user needs to click a link in that email to verify their address.`,
		int(BuyUpfrontOnly):                                   `This customer can buy for a full up-front payment only.`,
		int(NoPointsEarned):                                   `This customer does not earn new reward points.`,
		int(InconsistentDocumentsForInvoice):                  `It is not possible to create an invoice from these source documents.All source documents must have the same type and same client (or the same payer).`,
		int(WrongSourceDocument):                              `Source document cannot be an invoice, invoice-waybill, POS receipt or a credit invoice.`,
		int(MultiComponentsTax):                               `Tax already has more than one component of this type, you must use saveVatRateComponent to add or change tax components.(Attribute "errorField" indicates the invalid input parameter.)`,
		int(WrongSalesPromotionType):                          `Sales Promotion - Only promotions with type "manual" can be set to require manager's approval.`,
		int(WrongPriceListAssociation):                        `This price list is not associated with this store region.`,
		int(AmountFieldError):                                 `The "amount" field for price list items can be used only if the "Quantity Price Lists" module has been enabled on your account.`,
		int(ProductIDChangeFailure):                           `When editing a price list item, product ID can not be changed.`,
		int(PrintingServiceFailure):                           `Printing service is not running at the moment.(User can turn printing service on from their Erply account).`,
		int(EmailSendingFailure):                              `Email sending failed.`,
		int(EmailSettingFailure):                              `Email sending has been incorrectly set up, review settings in ERPLY.(Missing sender's address or empty message content).`,
		int(WrongMasterListSetup):                             `"Master List" functionality has not been fully set up yet, some requirements are missing.`,
		int(WrongMasterListUniqueField):                       `Configuration parameter "master_list_unique_field" has been incorrectly set up.`,
		int(NoFileAttached):                                   `No file attached.`,
		int(WrongFileEncoding):                                `Attached file is not encoded with Base64.`,
		int(TooBigFile):                                       `Attached file exceeds allowed size limit.`,
		int(PasswordLengthFailure):                            `New password must contain at least 8 characters.`,
		int(WrongLettersInPassword):                           `New password may only contain Latin letters and digits.(This rule is enforced by configuration parameter "password_only_alphanumeric_allowed").`,
		int(PasswordComplexityError):                          `New password must contain at least one small letter, one capital letter and one digit.`,
		int(PasswordChangeLimitFailure):                       `A configuration setting does not allow the user to change own password more often than once every N days.`,
		int(MultipleConflictingSettingsInSalesPromotion):      `Sales Promotion - Multiple conflicting settings.A promotion must apply to all stores, or specific regions only, or specific location only, or specific store group only.`,
		int(SalesPromotionPurchasedAmountConflict):            `Sales Promotion - Fields "purchasedProductGroupID", "purchasedProductCategoryID" or "purchasedProducts" are only allowed together with "purchasedAmount".`,
		int(SalesPromotionMultipleConflictingPurchaseOptions): `Sales Promotion - Multiple conflicting purchase options ("purchasedProductGroupID", "purchasedProductCategoryID" or "purchasedProducts") have been specified at the same time.`,
		int(SalesPromotionAwardedProductWithSumOffConflict):   `Sales Promotion - Fields "awardedProductGroupID", "awardedProductCategoryID", "awardedProducts", "awardedAmount" are only allowed together with "sumOFF" or "percentageOFF".`,
		int(SalesPromotionAwardedProductConflict):             `Sales Promotion - Multiple conflicting award options ("awardedProductGroupID", "awardedProductCategoryID", or "awardedProducts", ) have been specified at the same time.`,
		int(SalesPromotionPercentageOffConflict):              `Sales Promotion - Fields "percentageOffExcludedProducts" and "percentageOffIncludedProducts" are only allowed together with "percentageOffEntirePurchase".`,
		int(SalesPromotionSumOffConflict):                     `Sales Promotion - Fields "sumOffExcludedProducts" and "sumOffIncludedProducts" are only allowed together with "sumOffEntirePurchase".`,
		int(SalesPromotionPriceWithPurchasedAmountConflict):   `Sales Promotion - Fields "priceAtLeast" and "priceAtMost" are only allowed together with "purchasedAmount".`,
		int(SalesPromotionMaxPointsDiscountWithRewardPointsConflict):  `Sales Promotion - Field "maximumPointsDiscount" can only be used together with "rewardPoints" and "sumOffEntirePurchase".`,
		int(SalesPromotionLowestPriceWithSumOffConflict):              `Sales Promotion - Field "lowestPriceItemIsAwarded" can only be used together with "sumOFF" or "percentageOFF".`,
		int(CustomerRegistryServiceUsed):                              `This account uses customer registry microservice.The list of customers, their groups and addresses is stored outside of ERPLY.Queries and updates must be sent directly to the service, using the service's own API. See the output of verifyUser for a service endpoint and authentication token.`,
		int(CannotChangeTypeOfDocument):                               `The type of a confirmed document cannot be changed.`,
		int(SalesPromotionSpecialPriceWithPurchasedAmountConflict):    `Sales Promotion - Field "specialPrice" can only be used together with "purchasedAmount".`,
		int(SalesPromotionPercentageOffWithPurchasedAmountConflict):   `Sales Promotion - Fields "percentageOffMatchingItems" and "sumOffMatchingItems" are only allowed together with "purchasedAmount".`,
		int(SalesDocumentAccountNeedsUpdate):                          `Sales Document - For creating recurring billing invoices over API, the data model of your account needs an update.Please contact customer support.`,
		int(AccountLimitationOnIntegration):                           `This account uses customer registry microservice.This input field in this API call cannot be used; this is a limitation of the integration.`,
		int(AccountInputFieldIsNotAllowed):                            `This account uses customer registry microservice.This input field in this API call is not allowed to have that value; this is a limitation of the integration.`,
		int(GreekAccountsOnly):                                        `This field can only be used on Greek accounts.`,
		int(OnlyOneValueForSalesPromotion):                            `Sales Promotion - Flag "excludeDiscountedFromPercentageOffEntirePurchase" can only be set to 1 if you have specified "percentageOffEntirePurchase".`,
		int(WrongBillingAssociation):                                  `Sales Document - One or more billing readings on that row are not associated with the specified billing statement or are already associated with another invoice.`,
		int(SalesPromotionWrongReasonCode):                            `Sales Promotion - The purpose of the Reason Code must be "PROMOTION".`,
		int(SalesPromotionPurchasedProductWithSumOffConflict):         `Sales Promotion - Field "purchasedProductSubsidies" can only be used together with "purchasedProducts" and ("percentageOffMatchingItems" or "sumOffMatchingItems").`,
		int(SalesPromotionPurchasedProductSameAmountError):            `Sales Promotion - Field "purchasedProductSubsidies" must contain exactly the same number of elements as field "purchasedProducts".`,
		int(SalesPromotionAwardedProductAmountError):                  `Sales Promotion - Field "awardedProductSubsidies" must contain exactly the same number of elements as field "awardedProducts".`,
		int(NoAssortmentPossibleForLocation):                          `This location does not have an assortment.`,
		int(TooManyProducts):                                          `This location's assortment contains more than 10,000 products; API is not going to return the list.`,
		int(TooManyBillingIds):                                        `The number of billing statement IDs must not exceed 500.`,
		int(SalesPromotionSpecialUnitPurchasedAmountConflict):         `Sales Promotion - Field "specialUnitPrice" can only be used together with "purchasedAmount".`,
		int(SalesPromotionMaxItemsBiggerThanPurchasedAmount):          `Sales Promotion - Field "maxItemsWithSpecialUnitPrice" must be equal to or larger than "purchasedAmount".`,
		int(SalesPromotionPurchasedAmountMissingPurchasedProductData): `Sales Promotion - Field "purchasedAmount" can only be used together with "purchasedProductGroupID", "purchasedProductCategoryID", or "purchasedProducts".`,
		int(ExternalCouponRegistryService):                            `This account uses coupon registry microservice and this API call is not supported.`,
		int(ExternalCouponRegistryServiceWrongInputField):             `This account uses coupon registry microservice.This input field in this API call is not allowed to have that value; this is a limitation of the integration.`,
		int(SalesPromotionRedemptionLimitTooBig):                      `Sales Promotion - Field "redemptionLimit" is not allowed for promotions that give % off entire invoice, require reward points or apply to an unlimited number of items.`,
		int(SalesPromotionRedemptionLimitWithMaxItemsConflict):        `Sales Promotion - Field "redemptionLimit" can only be used together with "maxItemsWithSpecialUnitPrice" (for special unit price promotions).`,
		int(NoEmployeeRecord):                                         `You do not have an employee record.Please ask a manager to create an employee record for you.`,
		int(ComplianceAlreadyConfirmed):                               `You have already confirmed your compliance with the General Data Protection Regulation.`,
		int(NoAccessToCustomerData):                                   `You do not have access to customer data.Please contact your manager to receive an introduction to the General Data Protection Regulation, to get instructions about proper data protection procedures and to confirm that you will comply with them.`,
		int(NonEuCountry):                                             `Your account country is a non-EU country and the GDPR customer data processing log is not available.(Should you need the logging feature regardless, please let us know.)`,
		int(IntegrationSpecificFieldMissingInputParameter):            `If you attempt to add a integration-specific field to a payment, you also need to set input parameter "paymentServiceProvider".See the documentation of savePayment to find out what should be the appropriate input value for "paymentServiceProvider".`,
		int(ProductAlreadyAppearsInSpecialPriceList):                  `This product cannot be added to store price list.It already occurs in a Flyer or Manager's Special price list that is active during the same time period.`,
		int(PriceListOverlapsWithSpecialPriceList):                    `This store price list cannot be updated.After the update, the price list would overlap with one or more Flyer or Manager's Special price lists that contain the same products.`,
		int(WrongPurposeForReasonCode):                                `Inventory Registration - The purpose of the Reason Code must be "REGISTRATION".`,
		int(OldInventoryModule):                                       `This API call does not support the old inventory module.Please contact customer support to upgrade your inventory.`,
		int(CreateAccountError):                                       `Creating a new account is temporarily not possible.Please try again in 5 minutes.`,
		int(ValueLengthError):                                         `Value must not be longer than 100 characters.(Attribute "errorField" indicates the invalid input parameter.)`,
		int(BulkSubRequestDocumentError):                              `This sub-request in a bulk call cannot be executed.It refers to the special value "CURRENT_INVOICE_ID", but the preceding "saveSalesDocument" call returned an error code and no document was created.`,
		int(BulkSubRequestDuplicateError):                             `This sub-request in a bulk call cannot be executed.It refers to the special value "CURRENT_INVOICE_ID", but the preceding "saveSalesDocument" call was flagged as a duplicate and no document was created.`,
		int(EuAccountsField):                                          `This field can only be used on EU accounts.`,
		int(GiftCardVatRateWithPaymentTypeConflict):                   `The "giftCardVatRateID" field for a payment can be used only if the payment's type is "GIFTCARD".`,
		int(PosAppRequestType):                                        `This request is specific to POS applications, and cannot be called by other API clients.`,
		int(TooLongListOfElements):                                    `Provided list of elements is too long.(Attribute "errorField" indicates the invalid input parameter, documentation of exact call contains information about size limits.)`,
		int(CustomerCodeMissingInJWT):                                 `Customer code (clientCode) is missing from Json Web Token.`,
		int(NoRightsForBackOffice):                                    `User doesn't have rights for Back Office. Access to Back Office can be granted in Identity.`,
		int(MissingUsernameInJWT):                                     `Username is missing from Json Web Token.`,
		int(MissingUserName):                                          `There is no such username in Erply.`,
		int(PendingStatusForUser):                                     `This user exists but is in "pending" status.Manager has to add this user to a user group.`,
		int(UsernameAlreadyExists):                                    `A user with this username already exists.Cannot create a new user with this username.`,
		int(NotPossibleToExtendSessionForJWT):                         `This is a JWT-based session and therefore Erply does not have the authority to extend or delegate the session.`,
		int(SignupNotAllowedForCountry):                               `createInstallation() - New signup are not allowed for provided country.`,
		int(ContactPersonBelongsToAnotherContact):                     `This contact person cannot be used because they are another customer's contact.`,
		int(DateIsNotFutureError):                                     `addInvoiceAlgorithmChange() - The "date" parameter must specify a future date.`,
		int(WrongVersionNumberForInvoice):                             `addInvoiceAlgorithmChange() - Provided version number is not allowed for this account.`,
		int(StockTakingIsConfirmed):                                   `Provided stocktaking is already confirmed.`,
		int(SalesPromotionFlagExcludePromotionWrongValue):             `Sales Promotion - Flag "excludePromotionItemsFromPercentageOffEntirePurchase" can only be set to 1 if you have specified "percentageOffEntirePurchase".`,
		int(CDNIntegrationRequired):                                   `CDN integration has been enabled.This request should be done against CDN API.`,
		int(SalesPromotionMaxNrOfMatchingItemsConflictingValue):       `Sales Promotion - Field "maximumNumberOfMatchingItems" can only be used together with "percentageOffMatchingItems" or "sumOffMatchingItems".`,
		int(SalesPromotionMaxNrOfMatchingItemsWrongValue):             `Sales Promotion - Field "maximumNumberOfMatchingItems" must be equal to or larger than "purchasedAmount".`,
		int(StocktakingAlreadyConnectedToDocumentType):                `Inventory Registration, Inventory Write-Off - Provided Inventory Stocktaking is already connected to document of such type.`,
		int(ConfigurationCallIsMissing):                               `Configuration related to this call is missing, check call's documentation.`,
		int(ConfigurationCallIsInvalid):                               `Configuration related to this call is invalid, check call's documentation.`,
		int(WarehouseNoLocalQuickButtonEnabled):                       `POS Store Quick Buttons - Provided warehouse doesnt have local quick buttons enabled.`,
		int(WrongJWTAccount):                                          `Account in the given JWT is not valid for this request.This means that the token was generated for a different account than the request was made for.`,
		int(JWTDecodingFailure):                                       `JWT decoding failed.This means that the provided token is in incorrect format or decoding failed due to invalid fingerprints.`,
		int(JWTExpired):                                               `JWT expired.The used JWT ttl has passed and cannot be used.`,
		int(WrongLanguageCode):                                        `This language code is not supported.`,
		int(NotNewPassword):                                           `New password has already been used in the past.Password history checks feature has been enabled on this account.`,
	}
}

func (s ApiError) String() string {
	strVal, ok := errorsMap[int(s)]

	if !ok {
		strVal = ""
	}

	return fmt.Sprintf("[%d] %s", int(s), strVal)
}

type ErplyError struct {
	error
	Status  string
	Message string
	Code    ApiError
}

func (e *ErplyError) Error() string {
	return fmt.Sprintf("ERPLY API: %s, status: %s, code: %d", e.Message, e.Status, e.Code)
}

func NewErplyError(status, msg string, code ApiError) *ErplyError {
	return &ErplyError{Status: status, Message: msg, Code: code}
}

func NewErplyErrorf(status, msg string, code ApiError, args ...interface{}) *ErplyError {
	return &ErplyError{Status: status, Message: fmt.Sprintf(msg, args...), Code: code}
}

func NewFromResponseStatus(status *Status) *ErplyError {
	var s string
	if status.ErrorField != "" {
		s = fmt.Sprintf("%s, error field: %s", status.ErrorCode.String(), status.ErrorField)
	} else {
		s = status.ErrorCode.String()
	}
	m := status.Request + ": " + status.ResponseStatus
	return &ErplyError{Status: s, Message: m, Code: status.ErrorCode}
}

func NewFromError(msg string, err error, code ApiError) *ErplyError {
	if err != nil {
		return NewErplyError("Error", errors.Wrap(err, msg).Error(), code)
	}
	return NewErplyError("Error", msg, code)
}

