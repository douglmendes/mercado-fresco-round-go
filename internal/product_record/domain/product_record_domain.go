package domain

type ProductRecord struct {
	Id             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

type ProductRecordRepository interface {
	GetById(id int) ([]ProductRecord, error)
	Create(arg ProductRecord) (ProductRecord, error)
}
