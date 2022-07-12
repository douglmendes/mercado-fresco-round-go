package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	mock_domain "github.com/douglmendes/mercado-fresco-round-go/internal/products/domain/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func callMock(t *testing.T) (*mock_domain.MockProductRepository, domain.ProductService) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mock_domain.NewMockProductRepository(ctrl)
	service := NewService(repository)

	return repository, service
}

func TestCreate(t *testing.T) {
	expected := domain.Product{
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
		buildStubs  func(repository *mock_domain.MockProductRepository)
		checkResult func(t *testing.T, result domain.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, nil)

				repository.
					EXPECT().
					Create(expected).
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.NoError(t, err)

				assert.EqualValues(t, expected, result)
			},
		},
		{
			name: "GetAllError",
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)

				assert.EqualValues(t, domain.Product{}, result)
			},
		},
		{
			name: "ConflictError",
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{expected}, nil)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("the product with code \"%s\" already exists", expected.ProductCode))

				assert.EqualValues(t, domain.Product{}, result)
			},
		},
		{
			name: "CreateError",
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, nil)

				repository.
					EXPECT().
					Create(expected).
					Times(1).
					Return(domain.Product{}, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)

				assert.EqualValues(t, domain.Product{}, result)
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
	expected := []domain.Product{
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
		buildStubs  func(repository *mock_domain.MockProductRepository)
		checkResult func(t *testing.T, result []domain.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result []domain.Product, err error) {
				assert.NoError(t, err)

				assert.ElementsMatch(t, expected, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, os.ErrPermission)
			},
			checkResult: func(t *testing.T, result []domain.Product, err error) {
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
	expected := domain.Product{
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
		buildStubs  func(repository *mock_domain.MockProductRepository)
		checkResult func(t *testing.T, result domain.Product, err error)
	}{
		{
			name:      "OK",
			productId: expected.Id,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(expected.Id).
					Times(1).
					Return(expected, nil)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.NoError(t, err)

				assert.Equal(t, expected, result)
			},
		},
		{
			name:      "NotFound",
			productId: nonExistentId,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(nonExistentId).
					Times(1).
					Return(domain.Product{}, fmt.Errorf("product (%d) not found", nonExistentId))
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("product (%d) not found", nonExistentId))

				assert.Equal(t, domain.Product{}, result)
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
	firstProduct := domain.Product{
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

	secondProduct := domain.Product{
		Id:                             2,
		ProductCode:                    "xablau",
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

	allProducts := []domain.Product{firstProduct, secondProduct}

	updatedProduct := domain.Product{
		Id:                             1,
		ProductCode:                    "xpto",
		Description:                    "xpto description",
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

	conflictingUpdatedProduct := domain.Product{
		Id:                             1,
		ProductCode:                    "xablau",
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
		name           string
		updatedProduct domain.Product
		buildStubs     func(repository *mock_domain.MockProductRepository)
		checkResult    func(t *testing.T, result domain.Product, err error)
	}{
		{
			name:           "OK",
			updatedProduct: updatedProduct,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(updatedProduct.Id).
					Times(1).
					Return(firstProduct, nil)

				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return(allProducts, nil)

				repository.
					EXPECT().
					Update(updatedProduct).
					Times(1).
					Return(updatedProduct, nil)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.NoError(t, err)

				assert.Equal(t, updatedProduct, result)
			},
		},
		{
			name:           "NotFound",
			updatedProduct: updatedProduct,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(updatedProduct.Id).
					Times(1).
					Return(domain.Product{}, fmt.Errorf("firstProduct (%d) not found", updatedProduct.Id))
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("firstProduct (%d) not found", updatedProduct.Id))

				assert.Equal(t, domain.Product{}, result)
			},
		},
		{
			name:           "Fail",
			updatedProduct: updatedProduct,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(updatedProduct.Id).
					Times(1).
					Return(firstProduct, nil)

				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return(allProducts, nil)

				repository.
					EXPECT().
					Update(updatedProduct).
					Times(1).
					Return(domain.Product{}, os.ErrClosed)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)

				assert.Equal(t, domain.Product{}, result)
			},
		},
		{
			name:           "Conflict",
			updatedProduct: conflictingUpdatedProduct,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(conflictingUpdatedProduct.Id).
					Times(1).
					Return(firstProduct, nil)

				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return(allProducts, nil)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Equal(
					t,
					fmt.Errorf(
						"the product with code \"%s\" already exists",
						conflictingUpdatedProduct.ProductCode,
					),
					err,
				)

				assert.Equal(t, domain.Product{}, result)
			},
		},
		{
			name:           "Fail2",
			updatedProduct: conflictingUpdatedProduct,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					GetById(conflictingUpdatedProduct.Id).
					Times(1).
					Return(firstProduct, nil)

				repository.
					EXPECT().
					GetAll().
					Times(1).
					Return([]domain.Product{}, os.ErrClosed)
			},
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)

				assert.Equal(t, domain.Product{}, result)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service := callMock(t)

			testCase.buildStubs(repository)

			result, err := service.Update(testCase.updatedProduct)
			testCase.checkResult(t, result, err)
		})
	}
}

func TestDelete(t *testing.T) {
	existentId := 1
	nonExistentId := 99

	testCases := []struct {
		name        string
		productId   int
		buildStubs  func(repository *mock_domain.MockProductRepository)
		checkResult func(t *testing.T, err error)
	}{
		{
			name:      "OK",
			productId: existentId,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					Delete(existentId).
					Times(1).
					Return(nil)
			},
			checkResult: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:      "NotFound",
			productId: nonExistentId,
			buildStubs: func(repository *mock_domain.MockProductRepository) {
				repository.
					EXPECT().
					Delete(nonExistentId).
					Times(1).
					Return(fmt.Errorf("product (%d) not found", nonExistentId))
			},
			checkResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("product (%d) not found", nonExistentId))
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository, service := callMock(t)

			testCase.buildStubs(repository)

			err := service.Delete(testCase.productId)
			testCase.checkResult(t, err)
		})
	}
}
