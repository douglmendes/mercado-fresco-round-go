package domain

import (
	"context"
)

type Carrier struct {
	Id          int    `json:"id,omitempty"`
	Cid         string `json:"cid,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	LocalityId  int    `json:"locality_id,omitempty"`
}

//go:generate mockgen -source=./carrier.go -destination=./mock/carrier_mock.go
type CarrierRepository interface {
	GetAll(ctx context.Context) ([]Carrier, error)
	Create(ctx context.Context, cid, companyName, address, telephone string, localityId int) (Carrier, error)
}

type CarrierService interface {
	CreateCarrier(ctx context.Context, cid, companyName, address, telephone string, localityId int) (Carrier, error)
}
