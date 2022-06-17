package mocks

import (
	"github.com/douglmendes/mercado-fresco-round-go/internal/sellers"
	"github.com/stretchr/testify/mock"
)

type SellerRepository struct {
	mock.Mock
}

func (s *SellerRepository) GetAll() ([]sellers.Seller, error) {
	args := s.Called()

	var seller []sellers.Seller

	if rf, ok := args.Get(0).(func() []sellers.Seller); ok {
		seller = rf()
	} else {
		if args.Get(0) != nil {
			seller = args.Get(0).([]sellers.Seller)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return seller, err
	
}