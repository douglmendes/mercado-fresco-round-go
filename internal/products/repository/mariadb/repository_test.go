package mariadb

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
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
	noProducts = []domain.Product{}
)

func TestMariaDB_GetAll(t *testing.T) {
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
		{
			name: "Fail",
			buildStubs: func() {
				mock.ExpectQuery(GetAllQuery).WillReturnError(sql.ErrConnDone)
			},
			checkResult: func(t *testing.T, result []domain.Product, err error) {
				assert.Error(t, err)
				assert.Equal(t, noProducts, result)
			},
		},
		{
			name: "Scan Fail",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"product_code",
				}).AddRow(
					firstProduct.Id,
					firstProduct.ProductCode,
				).AddRow(
					secondProduct.Id,
					secondProduct.ProductCode,
				)

				mock.ExpectQuery(GetAllQuery).WillReturnRows(rows)
			},
			checkResult: func(t *testing.T, result []domain.Product, err error) {
				assert.Error(t, err)
				assert.Equal(t, noProducts, result)
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

func TestMariaDB_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		id          int
		checkResult func(t *testing.T, result domain.Product, err error)
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
				)

				mock.ExpectQuery(GetByIdQuery).WillReturnRows(rows)
			},
			id: firstProduct.Id,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.NoError(t, err)
				assert.Equal(t, firstProduct, result)
			},
		},
		{
			name: "Scan Fail",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"id",
					"product_code",
				}).AddRow(
					firstProduct.Id,
					firstProduct.ProductCode,
				)

				mock.ExpectQuery(GetByIdQuery).WillReturnRows(rows)
			},
			id: firstProduct.Id,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.Equal(t, emptyProduct, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.GetById(testCase.id)

			testCase.checkResult(t, result, err)
		})
	}
}

func TestMariaDB_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		product     domain.Product
		checkResult func(t *testing.T, result domain.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(
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
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			product: firstProduct,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.NoError(t, err)
				assert.Equal(t, firstProduct, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(
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
					).
					WillReturnError(sql.ErrConnDone)
			},
			product: firstProduct,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.Equal(t, emptyProduct, result)
			},
		},
		{
			name: "Last ID Error",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(CreateQuery)).
					WithArgs(
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
					).
					WillReturnResult(driver.ResultNoRows)
			},
			product: firstProduct,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.Equal(t, emptyProduct, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.Create(testCase.product)

			testCase.checkResult(t, result, err)
		})
	}
}

func TestMariaDB_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		product     domain.Product
		checkResult func(t *testing.T, result domain.Product, err error)
	}{
		{
			name: "OK",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(
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
						firstProduct.Id,
					).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			product: firstProduct,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.NoError(t, err)
				assert.Equal(t, firstProduct, result)
			},
		},
		{
			name: "Fail",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(UpdateQuery)).
					WithArgs(
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
						firstProduct.Id,
					).
					WillReturnError(sql.ErrConnDone)
			},
			product: firstProduct,
			checkResult: func(t *testing.T, result domain.Product, err error) {
				assert.Error(t, err)
				assert.Equal(t, emptyProduct, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.Update(testCase.product)

			testCase.checkResult(t, result, err)
		})
	}
}

func TestMariaDB_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		id          int
		checkResult func(t *testing.T, err error)
	}{
		{
			name: "OK",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WithArgs(firstProduct.Id).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			id: firstProduct.Id,
			checkResult: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Fail",
			buildStubs: func() {
				mock.
					ExpectExec(regexp.QuoteMeta(DeleteQuery)).
					WithArgs(firstProduct.Id).
					WillReturnError(sql.ErrConnDone)
			},
			id: firstProduct.Id,
			checkResult: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			err := repository.Delete(testCase.id)

			testCase.checkResult(t, err)
		})
	}
}
