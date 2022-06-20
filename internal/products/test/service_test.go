package products_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products"
	mock_products "github.com/douglmendes/mercado-fresco-round-go/internal/products/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func callMock(t *testing.T) (*mock_products.MockRepository, products.Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_products.NewMockRepository(ctrl)
	service := products.NewService(repository)

	return repository, service
}

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
		{
			name: "ConflictError",
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
					Return([]products.Product{expected}, nil)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("the product with code \"%s\" already exists", expected.ProductCode))

				assert.EqualValues(t, products.Product{}, result)
			},
		},
		{
			name: "CreateError",
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
					Return(products.Product{}, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.Error(t, err)

				assert.EqualValues(t, products.Product{}, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service := callMock(t)

			testCase.buildStubs(repository)

			result, err := service.Create(expected)
			testCase.checkResult(t, result, err)
		})
	}
}

func TestGetAll(t *testing.T) {
	expected := []products.Product{
		{
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
		},
		{
			Id:                             2,
			ProductCode:                    "xablau",
			Description:                    "description",
			Width:                          3.6,
			Height:                         3.2,
			Length:                         1.5,
			NetWeight:                      5.23,
			ExpirationRate:                 0.08,
			RecommendedFreezingTemperature: -3.4,
			FreezingRate:                   0.8,
			ProductTypeId:                  2,
			SellerId:                       3,
		},
	}

	testCases := []struct {
		name        string
		buildStubs  func(repository *mock_products.MockRepository)
		checkResult func(t *testing.T, result []products.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result []products.Product, err error) {
				assert.NoError(t, err)

				assert.ElementsMatch(t, expected, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]products.Product{}, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result []products.Product, err error) {
				assert.Error(t, err)

				assert.Empty(t, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service := callMock(t)

			testCase.buildStubs(repository)

			result, err := service.GetAll()
			testCase.checkResult(t, result, err)
		})
	}
}

func TestGetById(t *testing.T) {
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

	nonExistentId := 15

	testCases := []struct {
		name        string
		productId   int
		buildStubs  func(repository *mock_products.MockRepository)
		checkResult func(t *testing.T, result products.Product, err error)
	}{
		{
			name:      "OK",
			productId: expected.Id,
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					GetById(expected.Id).
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.NoError(t, err)

				assert.Equal(t, expected, result)
			},
		},
		{
			name:      "NotFound",
			productId: nonExistentId,
			buildStubs: func(repository *mock_products.MockRepository) {
				repository.
					EXPECT().
					GetById(nonExistentId).
					Times(1).
					Return(products.Product{}, fmt.Errorf("product (%d) not found", nonExistentId))
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("product (%d) not found", nonExistentId))

				assert.Equal(t, products.Product{}, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service := callMock(t)

			testCase.buildStubs(repository)

			result, err := service.GetById(testCase.productId)
			testCase.checkResult(t, result, err)
		})
	}
}

func TestUpdate(t *testing.T) {
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
					Update(expected).
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result products.Product, err error) {
				assert.NoError(t, err)

				assert.Equal(t, expected, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service := callMock(t)

			testCase.buildStubs(repository)

			result, err := service.Update(expected)
			testCase.checkResult(t, result, err)
		})
	}
}
