package service

import (
	"context"
	"fmt"
	carrierRepo "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	warehouseRepo "github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
)

type service struct {
	carrierRepository   carrierRepo.CarrierRepository
	warehouseRepository warehouseRepo.WarehouseRepository
}

func NewService(r carrierRepo.CarrierRepository) carrierRepo.CarrierService {
	return &service{
		carrierRepository: r,
	}
}

func (s service) CreateCarrier(
	ctx context.Context,
	cid,
	companyName,
	address,
	telephone string,
	localityId int,
) (carrierRepo.Carrier, error) {

	carrierList, err := s.carrierRepository.GetAll(ctx)
	if err != nil {
		return carrierRepo.Carrier{}, err
	}

	for _, carrier := range carrierList {
		if carrier.Cid == cid {
			return carrierRepo.Carrier{}, fmt.Errorf("already exists a carrier with this cid: %s", cid)
		}
	}

	carrier, err := s.carrierRepository.Create(ctx, cid, companyName, address, telephone, localityId)
	if err != nil {
		return carrierRepo.Carrier{}, err
	}

	// aqui precisa verificar sobre locality_id, se existe ou nao, deve criar um carrier se a localidade existir, caso contrario retornar 409
	// agaurdar implementar o dominio de localidade - discutir a melhor forma de fazer isso

	return carrier, nil
}
