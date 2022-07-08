package service

import (
	"context"
	"fmt"
	carrierRepo "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	localityRepo "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
)

type service struct {
	carrierRepository  carrierRepo.CarrierRepository
	localityRepository localityRepo.LocalityRepository
}

func NewService(r carrierRepo.CarrierRepository, l localityRepo.LocalityRepository) carrierRepo.CarrierService {
	return &service{
		carrierRepository:  r,
		localityRepository: l,
	}
}

func (s *service) CreateCarrier(
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

	locality, err := s.localityRepository.GetById(ctx, localityId)
	if locality.Id == 0 {
		return carrierRepo.Carrier{}, fmt.Errorf("locality %d not found", localityId)
	}

	carrier, err := s.carrierRepository.Create(ctx, cid, companyName, address, telephone, localityId)
	if err != nil {
		return carrierRepo.Carrier{}, err
	}
	return carrier, nil
}
