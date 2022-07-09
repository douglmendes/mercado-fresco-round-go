package mariadb

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/douglmendes/mercado-fresco-round-go/internal/product_record/domain"
	"github.com/stretchr/testify/assert"
)

var (
	firstProductRecordsCount = domain.ProductRecordCount{
		ProductId:    2,
		Description:  "Chocolate",
		RecordsCount: 3,
	}
	secondProductRecordsCount = domain.ProductRecordCount{
		ProductId:    4,
		Description:  "Ice Cream",
		RecordsCount: 1,
	}
	allProductRecordsCount = []domain.ProductRecordCount{
		firstProductRecordsCount,
		secondProductRecordsCount,
	}
	noProductRecordsCount = []domain.ProductRecordCount{}
)

func TestMariaDB_GetByProductId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	testsCases := []struct {
		name        string
		buildStubs  func()
		productId   int
		checkResult func(t *testing.T, result []domain.ProductRecordCount, err error)
	}{
		{
			name: "OK_GetAll",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"product_id",
					"description",
					"records_count",
				}).AddRow(
					firstProductRecordsCount.ProductId,
					firstProductRecordsCount.Description,
					firstProductRecordsCount.RecordsCount,
				).AddRow(
					secondProductRecordsCount.ProductId,
					secondProductRecordsCount.Description,
					secondProductRecordsCount.RecordsCount,
				)

				mock.ExpectQuery(regexp.QuoteMeta(GetAllGroupByProductIdQuery)).WillReturnRows(rows)
			},
			productId: 0,
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.NoError(t, err)
				assert.Equal(t, allProductRecordsCount, result)
			},
		},
		{
			name: "Fail_GetAll",
			buildStubs: func() {
				mock.
					ExpectQuery(regexp.QuoteMeta(GetAllGroupByProductIdQuery)).
					WillReturnError(sql.ErrConnDone)
			},
			productId: 0,
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.Error(t, err)
				assert.Equal(t, noProductRecordsCount, result)
			},
		},
		{
			name: "OK_GetByProductId",
			buildStubs: func() {
				rows := sqlmock.NewRows([]string{
					"product_id",
					"description",
					"records_count",
				}).AddRow(
					firstProductRecordsCount.ProductId,
					firstProductRecordsCount.Description,
					firstProductRecordsCount.RecordsCount,
				)

				mock.ExpectQuery(regexp.QuoteMeta(GetAllGroupByProductIdWhereIdQuery)).WillReturnRows(rows)
			},
			productId: firstProductRecordsCount.ProductId,
			checkResult: func(t *testing.T, result []domain.ProductRecordCount, err error) {
				assert.NoError(t, err)
				assert.Equal(t, []domain.ProductRecordCount{firstProductRecordsCount}, result)
			},
		},
	}

	for _, testCase := range testsCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.buildStubs()

			repository := NewRepository(db)

			result, err := repository.GetByProductId(testCase.productId)

			testCase.checkResult(t, result, err)
		})
	}
}
