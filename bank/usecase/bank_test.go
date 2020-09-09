package usecase_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/OLTeam-go/sea-store-backend-transactions/bank/usecase"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"

	"github.com/stretchr/testify/mock"

	mocks "github.com/OLTeam-go/sea-store-backend-transactions/mocks/domain"
)

func Test_Bank(t *testing.T) {
	mockRepo := new(mocks.BankRepository)
	mockBank1 := models.Bank{
		Name:   "BNI",
		Active: true,
	}
	var mockList []*models.Bank
	mockList = append(mockList, &mockBank1)
	mockRepo.On("Fetch", mock.Anything).Return(mockList, nil)
	t.Run("success fetch bank", func(t *testing.T) {
		bu := usecase.New(mockRepo, time.Second*3)
		res, err := bu.Fetch(context.TODO())

		assert.NoError(t, err)
		assert.True(t, reflect.DeepEqual(mockBank1, *(res[0])))
	})
}
