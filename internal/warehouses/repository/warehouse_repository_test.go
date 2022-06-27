package repository

import (
	"encoding/json"
	"errors"
	"github.com/douglmendes/mercado-fresco-round-go/internal/warehouses/domain"
	"github.com/douglmendes/mercado-fresco-round-go/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetAll(t *testing.T) {
	t.Run("should return a warehouse list", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		input := []domain.Warehouse{
			{
				1,
				"Monroe 860",
				"47470000",
				"TSFK",
				10,
				10,
			},
			{
				2,
				"Rua do Teste 2",
				"555555555",
				"JJJ",
				10,
				2,
			},
		}

		dataJson, _ := json.Marshal(input)
		fileStoreMock := &store.Mock{
			Data: dataJson,
			Err:  nil,
		}

		fileStore.AddMock(fileStoreMock)
		result, _ := NewRepository(fileStore).GetAll()

		assert.Equal(t, result, input, "should be equal")
	})

	t.Run("should return err when Store returns an error", func(t *testing.T) {
		fileStore := store.New(store.FileType, "")

		expectedErr := errors.New("error on connect store / database")

		fileStoreMock := &store.Mock{
			Data: []byte{},
			Err:  expectedErr,
		}

		fileStore.AddMock(fileStoreMock)

		repository := NewRepository(fileStore)

		_, err := repository.GetAll()

		assert.Equal(t, err, expectedErr, "should be equal")
	})
}

func TestRepository_GetById(t *testing.T) {
	fileStore := store.New(store.FileType, "")

	wh := []domain.Warehouse{
		{
			1,
			"Monroe 860",
			"47470000",
			"TSFK",
			10,
			10,
		},
	}

	dataJson, _ := json.Marshal(wh)
	fileStoreMock := &store.Mock{
		Data: dataJson,
		Err:  nil,
	}

	fileStore.AddMock(fileStoreMock)

	result, err := NewRepository(fileStore).GetById(1)

	assert.Equal(t, result, wh[0])
	assert.Nil(t, err)

}

func TestRepository_GetById_NOK(t *testing.T) {
	fileStore := store.New(store.FileType, "")

	wh := []domain.Warehouse{
		{
			1,
			"Monroe 860",
			"47470000",
			"TSFK",
			10,
			10,
		},
	}

	dataJson, _ := json.Marshal(wh)
	fileStoreMock := &store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	fileStore.AddMock(fileStoreMock)

	result, err := NewRepository(fileStore).GetById(55)

	assert.Equal(t, result, domain.Warehouse{})
	assert.Error(t, err, "warehouse not found")
}

func TestRepository_GetById_ReadNotOK(t *testing.T) {
	fileStore := store.New(store.FileType, "")

	expectedErr := errors.New("error on connect store / database")
	fileStoreMock := &store.Mock{
		Data: []byte{},
		Err:  expectedErr,
	}
	fileStore.AddMock(fileStoreMock)

	repository := NewRepository(fileStore)
	_, err := repository.GetById(1)
	assert.Equal(t, err, expectedErr)
}
