package service

import (
	"context"
	"fmt"
	carrierRepo "github.com/douglmendes/mercado-fresco-round-go/internal/carriers/domain"
	localityRepo "github.com/douglmendes/mercado-fresco-round-go/internal/localities/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/logger"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
)

type service struct {
	carrierRepository  carrierRepo.CarrierRepository
	localityRepository localityRepo.LocalityRepository
}

func NewService(cr carrierRepo.CarrierRepository, lr localityRepo.LocalityRepository) carrierRepo.CarrierService {
	return &service{
		carrierRepository:  cr,
		localityRepository: lr,
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
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
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
		logger.Error(ctx, store.GetPathWithLine(), err.Error())
		return carrierRepo.Carrier{}, err
	}
	return carrier, nil
}
