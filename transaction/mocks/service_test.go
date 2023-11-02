package mocks

import (
	"errors"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/campaign/mocks"
	"github.com/RianIhsan/raise-unity/payment"
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTransactionByCampaignID(t *testing.T) {
	campaignRepository := mocks.NewRepository(t)
	repository := NewRepository(t)
	paymentService := payment.NewService()
	service := transaction.NewService(repository, campaignRepository, paymentService)

	campaignID := 1
	userID := 1
	campaignData := campaign.Campaign{ID: campaignID, UserID: userID}
	userData := user.User{ID: userID}

	transactions := []transaction.Transaction{
		{
			ID:         1,
			CampaignID: campaignID,
			UserID:     userID,
			Amount:     10000,
			Status:     "paid",
			Code:       "12345",
			PaymentURL: "https://sandbox.midtrans.com",
			User:       userData,
			Campaign:   campaignData,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	t.Run("Valid owner of the campaign", func(t *testing.T) {
		campaignRepository.On("FindByID", campaignID).Return(campaignData, nil).Once()

		repository.On("GetByCampaignID", campaignID).Return(transactions, nil).Once()
		input := transaction.GetCampaignTransactionInput{
			ID:   campaignID,
			User: userData,
		}

		result, err := service.GetTransactionByCampaignID(input)
		assert.NoError(t, err)
		assert.Equal(t, transactions, result)

		campaignRepository.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid owner of the campaign", func(t *testing.T) {
		campaignRepository.On("FindByID", campaignID).Return(campaignData, nil).Once()

		invalidUserData := user.User{ID: 2}
		input := transaction.GetCampaignTransactionInput{ID: campaignID, User: invalidUserData}
		result, err := service.GetTransactionByCampaignID(input)

		expectedError := errors.New("not owner of the campaign")
		assert.EqualError(t, err, expectedError.Error())
		assert.Empty(t, result)

		campaignRepository.AssertExpectations(t)

		repository.AssertExpectations(t)
	})

	t.Run("Error when finding campaign", func(t *testing.T) {
		expectedError := errors.New("error finding campaign")
		campaignRepository.On("FindByID", campaignID).Return(campaign.Campaign{}, expectedError).Once()

		input := transaction.GetCampaignTransactionInput{ID: campaignID, User: userData}
		result, err := service.GetTransactionByCampaignID(input)

		assert.EqualError(t, err, expectedError.Error())
		assert.Empty(t, result)

		repository.AssertExpectations(t)
	})
}

func TestGetTransactionByUserID(t *testing.T) {
	campaignRepository := mocks.NewRepository(t)
	repository := NewRepository(t)
	paymentService := payment.NewService()
	service := transaction.NewService(repository, campaignRepository, paymentService)

	campaignID := 1
	userID := 1
	campaignData := campaign.Campaign{ID: campaignID, UserID: userID}
	userData := user.User{ID: userID}

	transactions := []transaction.Transaction{
		{
			ID:         1,
			CampaignID: campaignID,
			UserID:     userID,
			Amount:     10000,
			Status:     "paid",
			Code:       "12345",
			PaymentURL: "https://sandbox.midtrans.com",
			User:       userData,
			Campaign:   campaignData,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
	t.Run("Valid user ID", func(t *testing.T) {
		repository.On("GetByUserID", userID).Return(transactions, nil).Once()
		result, err := service.GetTransactionByUserID(userID)

		assert.NoError(t, err)
		assert.Equal(t, transactions, result)

		repository.AssertExpectations(t)
	})

	t.Run("Error when finding transactions", func(t *testing.T) {
		expectedError := errors.New("error finding transactions")
		repository.On("GetByUserID", userID).Return(transactions, expectedError).Once()

		result, err := service.GetTransactionByUserID(userID)

		assert.EqualError(t, err, expectedError.Error())
		assert.NotEmpty(t, result)

		repository.AssertExpectations(t)
	})
}
