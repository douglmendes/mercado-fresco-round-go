package domain

import (
	"context"
)

type ProductBatch struct {
	Id                 int    `json:"id,omitempty"`
	BatchNumber        int    `json:"batch_number,omitempty"`
	CurrentQuantity    int    `json:"current_quantity,omitempty"`
	CurrentTemperature int    `json:"current_temperature,omitempty"`
	DueDate            string `json:"due_date,omitempty"`
	InitialQuantity    int    `json:"initial_quantity,omitempty"`
	ManufacturingDate  string `json:"manufacturing_date,omitempty"`
	ManufacturingHour  int    `json:"manufacturing_hour,omitempty"`
	MinimumTemperature int    `json:"minimum_temperature,omitempty"`
	ProductId          int    `json:"product_id,omitempty"`
	SectionId          int    `json:"section_id,omitempty"`
}

type SectionRecords struct {
	SectionId     int `json:"section_id,omitempty"`
	SectionNumber int `json:"section_number,omitempty"`
	ProductsCount int `json:"products_count,omitempty"`
}

//go:generate mockgen -source=./carrier.go -destination=./mock/carrier_mock.go
type ProductBatchesRepository interface {
	GetAll(ctx context.Context) ([]ProductBatch, error)
	Create(ctx context.Context, batchNumber, currentQuantity, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour, minimumTemperature, productId, sectionId int) (*ProductBatch, error)
	GetBySectionId(ctx context.Context, sectionId int) ([]SectionRecords, error)
}

type ProductBatchesService interface {
	Create(ctx context.Context, batchNumber, currentQuantity, currentTemperature int, dueDate string, initialQuantity int, manufacturingDate string, manufacturingHour, minimumTemperature, productId, sectionId int) (*ProductBatch, error)
	GetBySectionId(ctx context.Context, sectionId int) ([]SectionRecords, error)
}
