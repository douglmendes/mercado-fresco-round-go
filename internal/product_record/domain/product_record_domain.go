package domain

import "context"

type ProductRecord struct {
	Id             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

type ProductRecordCount struct {
	ProductId    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}

type ProductRecordRepository interface {
	GetByProductId(ctx context.Context, productId int) ([]ProductRecordCount, error)
	Create(ctx context.Context, arg ProductRecord) (ProductRecord, error)
}

type ProductRecordService interface {
	GetByProductId(ctx context.Context, productId int) ([]ProductRecordCount, error)
	Create(ctx context.Context, arg ProductRecord) (ProductRecord, error)
}
