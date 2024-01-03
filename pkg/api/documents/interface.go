package documents

import "context"

type Manager interface {
	GetPurchaseDocuments(ctx context.Context, filters map[string]string) ([]PurchaseDocument, error)
	GetPurchaseDocumentsWithStatus(ctx context.Context, filters map[string]string) (GetPurchaseDocumentsResponse, error)
	GetPurchaseDocumentsBulk(ctx context.Context, bulkRequest []map[string]interface{}, baseFilters map[string]string) (GetPurchaseDocumentResponseBulk, error)
}
