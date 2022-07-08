package mariadb

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/products/domain"
	"github.com/stretchr/testify/assert"
)

var (
	emptyProduct = domain.Product{}
	firstProduct = domain.Product{
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
	secondProduct = domain.Product{
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
	}
	allProducts = []domain.Product{
		firstProduct,
		secondProduct,
	}
)

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		checkResult func(t *testing.T, result []domain.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"product_code",
					"description",
					"width",
					"height",
					"length",
					"net_weight",
					"expiration_rate",
					"recommended_freezing_temperature",
					"freezing_rate",
					"product_type_id",
					"seller_id",
				}).AddRow(
					firstProduct.Id,
					firstProduct.ProductCode,
					firstProduct.Description,
					firstProduct.Width,
					firstProduct.Height,
					firstProduct.Length,
					firstProduct.NetWeight,
					firstProduct.ExpirationRate,
					firstProduct.RecommendedFreezingTemperature,
					firstProduct.FreezingRate,
					firstProduct.ProductTypeId,
					firstProduct.SellerId,
				).AddRow(
					secondProduct.Id,
					secondProduct.ProductCode,
					secondProduct.Description,
					secondProduct.Width,
					secondProduct.Height,
					secondProduct.Length,
					secondProduct.NetWeight,
					secondProduct.ExpirationRate,
					secondProduct.RecommendedFreezingTemperature,
					secondProduct.FreezingRate,
					secondProduct.ProductTypeId,
					secondProduct.SellerId,
				)

				mock.ExpectQuery(GetAllQuery).WillReturnRows(rows)
			},
			checkResult: func(t *testing.T, result []domain.Product, err error) {
				assert.NoError(t, err)
				assert.Equal(t, allProducts, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.GetAll()

			testCase.checkResult(t, result, err)
		})
	}
}
