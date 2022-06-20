package products_test

import (
	"os"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	mock_products "github.com/douglmendes/mercado-fresco-round-go/internal/products/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	expected := products.Product{
		Id:                             1,
		ProductCode:                    "xpto",
		Description:                    "description",
		Width:                          6.3,
		Height:                         2.3,
		Length:                         5.1,
		NetWeight:                      23.5,
		ExpirationRate:                 0.8,
		RecommendedFreezingTemperature: -4.3,
		FreezingRate:                   0.4,
		ProductTypeId:                  3,
		SellerId:                       5,
	}

	testCases := []struct {
		name        string
		buildStubs  func(repository *mock_products.MockRepository)
		checkResult func(t *testing.T, result products.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					LastID().
					Times(1).
					Return(0, nil)

				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]products.Product{}, nil)

				repository.
					EXPECT().
					Create(expected).
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.NoError(t, err)

				assert.EqualValues(t, expected, result)
			},
		},
		{
			name: "LastIdError",
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					LastID().
					Times(1).
					Return(0, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.Error(t, err)

				assert.EqualValues(t, products.Product{}, result)
			},
		},
		{
			name: "GetAllError",
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					LastID().
					Times(1).
					Return(0, nil)

				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]products.Product{}, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.Error(t, err)

				assert.EqualValues(t, products.Product{}, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repository := mock_products.NewMockRepository(ctrl)
			service := products.NewService(repository)

			testCase.buildStubs(repository)

			result, err := service.Create(expected)
			testCase.checkResult(t, result, err)
		})
	}
}