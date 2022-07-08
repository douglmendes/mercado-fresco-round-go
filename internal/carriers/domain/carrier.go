package domain

import (
	"context"
	"github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
)

type Carrier struct {
	Id          int    `json:"id,omitempty"`
	Cid         string `json:"cid,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	LocalityId  int    `json:"locality_id,omitempty"`
}

type CarrierRepository interface {
	GetAll(ctx context.Context) ([]Carrier, error)
	GetLocal(id int) (domain.Locality, error)
	Create(ctx context.Context, cid, companyName, address, telephone string, localityId int) (Carrier, error)
}

type CarrierService interface {
	CreateCarrier(ctx context.Context, cid, companyName, address, telephone string, localityId int) (Carrier, error)
}
