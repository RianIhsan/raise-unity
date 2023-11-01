package admin

import (
	"errors"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/transaction"
	"testing"
	"time"

	"github.com/RianIhsan/raise-unity/admin/mocks"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersPagination(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	var userList = []user.User{
		{
			ID:         1,
			Name:       "Rian",
			Occupation: "programmer",
			Email:      "rianganteng@gmail.com",
			Password:   "rian12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: false,
		},
		{
			ID:         2,
			Name:       "Ihsan",
			Occupation: "programmer",
			Email:      "ihsanganteng@gmail.com",
			Password:   "ihsan12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: false,
		},
	}

	t.Run("Success get user pagination", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Get total user failed", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(0), errors.New("get total users error")).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get total users error")
		assert.Nil(t, users)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Get paginated users failed", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(nil, errors.New("get paginated users error")).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get paginated users error")
		assert.Nil(t, users)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
	})

	t.Run("page = 0", func(t *testing.T) {
		var page = 0
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 5
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 1
		var pageSize = 1

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > 1", func(t *testing.T) {
		var page = 2
		var pageSize = 1

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 1, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := service.GetUsersPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 1, prevPage)
		repository.AssertExpectations(t)
	})
}

func TestGetTransactionsPagination(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	var mockTransaction = []transaction.Transaction{
		{
			ID:         1,
			UserID:     3,
			CampaignID: 1,
			Amount:     10000,
			Status:     "paid",
			Code:       "48011",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/7a45765e-a0f3-4bfb-9b3f-a58560f09611",
			User:       user.User{Name: "user"},
			Campaign:   campaign.Campaign{Name: "Dana"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         2,
			UserID:     3,
			CampaignID: 1,
			Amount:     10000,
			Status:     "paid",
			Code:       "48011",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/7a45765e-a0f3-4bfb-9b3f-a58560f09611",
			User:       user.User{Name: "user"},
			Campaign:   campaign.Campaign{Name: "Dana"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	t.Run("Succes get all transaction users", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalTransactions").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransactions", 0, pageSize).Return(mockTransaction, nil).Once()
		transactions, totalPages, currentPage, nextPage, prevPage, err := service.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Failed Get total transaction user ", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalTransactions").Return(int64(0), errors.New("get total users transactions error")).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := service.GetTransactionsPagination(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get total users transactions error")
		assert.Nil(t, transactions)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page = 0", func(t *testing.T) {
		var page = 0
		var pageSize = 10

		repository.On("GetTotalTransactions").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransactions", 0, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := service.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 5
		var pageSize = 10

		repository.On("GetTotalTransactions").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransactions", 0, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := service.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 1
		var pageSize = 1

		repository.On("GetTotalTransactions").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransactions", 0, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := service.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > 1", func(t *testing.T) {
		var page = 2
		var pageSize = 1

		repository.On("GetTotalTransactions").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransactions", 1, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := service.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 1, prevPage)
		repository.AssertExpectations(t)
	})

}

func TestSearchUserByName(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	var userList = []user.User{
		{
			ID:         1,
			Name:       "Rian",
			Occupation: "programmer",
			Email:      "rianganteng@gmail.com",
			Password:   "rian12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: false,
		},
	}

	t.Run("Success get user", func(t *testing.T) {
		repository.On("SearchUserByName", "Rian").Return(userList, nil).Once()

		result, err := service.SearchUserByName("Rian")
		assert.Nil(t, err)
		assert.Equal(t, userList, result)
		repository.AssertExpectations(t)
	})

	t.Run("get user failed", func(t *testing.T) {
		repository.On("SearchUserByName", "Rian").Return(nil, errors.New("get user error")).Once()

		result, err := service.SearchUserByName("Rian")
		assert.Error(t, err)
		assert.EqualError(t, err, "get user error")
		assert.Nil(t, result)
		repository.AssertExpectations(t)
	})
}

func TestSearchTransactionByUsername(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	var transactionList = []transaction.Transaction{
		{
			ID:         1,
			UserID:     3,
			CampaignID: 1,
			Amount:     10000,
			Status:     "paid",
			Code:       "48011",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/7a45765e-a0f3-4bfb-9b3f-a58560f09611",
			User:       user.User{Name: "user"},
			Campaign:   campaign.Campaign{Name: "Dana"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	t.Run("Success get transaction list by username", func(*testing.T) {
		repository.On("SearchTransactionByUsername", "user").Return(transactionList, nil).Once()

		result, err := service.SearchTransactionByUsername("user")
		assert.Nil(t, err)
		assert.Equal(t, transactionList, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed get transaction list by username", func(*testing.T) {
		repository.On("SearchTransactionByUsername", "user").Return(nil, errors.New("Failed Get transaction user by user name")).Once()

		result, err := service.SearchTransactionByUsername("user")
		assert.Error(t, err)
		assert.EqualError(t, err, "Failed Get transaction user by user name")
		assert.Nil(t, result)
		repository.AssertExpectations(t)
	})
}

func TestDeleteUserById(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	var user = user.User{
		ID:         1,
		Name:       "Rian",
		Occupation: "programmer",
		Email:      "rianganteng@gmail.com",
		Password:   "rian12345",
		Avatar:     "www.cloudinary.com/avatar",
		Role:       "user",
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	t.Run("success delete user", func(t *testing.T) {
		repository.On("GetUserById", user.ID).Return(user, nil).Once()
		result, err := service.DeleteUserById(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)
		repository.AssertExpectations(t)

	})

	t.Run("Failed to delete user", func(t *testing.T) {
		repository.On("GetUserById", user.ID).Return(user, errors.New("User not found")).Once()

		_, err := service.DeleteUserById(user.ID)
		assert.Error(t, err)
		assert.Equal(t, "User not found", err.Error())
		repository.AssertExpectations(t)

	})
}

func TestDeleteCampaignId(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	var MockCampaign = campaign.Campaign{
		ID:               1,
		UserID:           1,
		Name:             "Sample Campaign",
		ShortDescription: "A short description",
		Description:      "A long description",
		Perks:            "Perks for backers",
		BackerCount:      100,
		GoalAmount:       10000,
		CurrentAmount:    5000,
		Status:           "active",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		CampaignImages: []campaign.CampaignImage{
			{
				ID:         1,
				CampaignID: 1,
				IsPrimary:  1,
				FileName:   "gambar.png",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
		},
		User: user.User{
			ID:         1,
			Name:       "Rian",
			Occupation: "programmer",
			Email:      "rianganteng@gmail.com",
			Password:   "rian12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	t.Run("Succes Delete campaign", func(t *testing.T) {
		repository.On("GetCampaignById", MockCampaign.ID).Return(MockCampaign, nil).Once()
		result, err := service.DeleteCampaignById(MockCampaign.ID)
		assert.NoError(t, err)
		assert.Equal(t, MockCampaign.ID, result.ID)
	})

	t.Run("Failed to delete campaign", func(t *testing.T) {
		repository.On("GetCampaignById", MockCampaign.ID).Return(MockCampaign, errors.New("Campaign not found")).Once()
		_, err := service.DeleteCampaignById(MockCampaign.ID)
		assert.Error(t, err)
		assert.Equal(t, "Campaign not found", err.Error())
		repository.AssertExpectations(t)

	})
}

func TestFindUserById(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	mockUser := user.User{
		ID:         1,
		Name:       "Rian",
		Occupation: "programmer",
		Email:      "rianganteng@gmail.com",
		Password:   "rian12345",
		Avatar:     "www.cloudinary.com/avatar",
		Role:       "user",
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	t.Run("Success get user id", func(t *testing.T) {
		repository.On("FindUserById", mockUser.ID).Return(mockUser, nil).Once()
		result, err := service.FindUserById(mockUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, result.ID)
		repository.AssertExpectations(t)

	})

	t.Run("Failed get user id", func(t *testing.T) {
		repository.On("FindUserById", mockUser.ID).Return(mockUser, errors.New("no user found with that ID")).Once()
		_, err := service.FindUserById(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, 1, mockUser.ID, "no user found with that ID")
		repository.AssertExpectations(t)
	})

}

func TestFindCampaignById(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = NewService(repository)

	mockCampaign := campaign.Campaign{
		ID:               1,
		UserID:           1,
		Name:             "Sample Campaign",
		ShortDescription: "A short description",
		Description:      "A long description",
		Perks:            "Perks for backers",
		BackerCount:      100,
		GoalAmount:       10000,
		CurrentAmount:    5000,
		Status:           "active",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		CampaignImages: []campaign.CampaignImage{
			{
				ID:         1,
				CampaignID: 1,
				IsPrimary:  1,
				FileName:   "gambar.png",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
		},
		User: user.User{
			ID:         1,
			Name:       "Rian",
			Occupation: "programmer",
			Email:      "rianganteng@gmail.com",
			Password:   "rian12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	t.Run("Success get campaign id", func(t *testing.T) {
		repository.On("FindCampaignById", mockCampaign.ID).Return(mockCampaign, nil).Once()
		result, err := service.FindCampaignById(mockCampaign.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockCampaign.ID, result.ID)
		repository.AssertExpectations(t)

	})

	t.Run("Failed get campaign id", func(t *testing.T) {
		repository.On("FindCampaignById", mockCampaign.ID).Return(mockCampaign, errors.New("no campaign found with that id")).Once()
		_, err := service.FindCampaignById(mockCampaign.ID)
		assert.Error(t, err)
		assert.Equal(t, 1, mockCampaign.ID, "no campaign found with that id")
		repository.AssertExpectations(t)
	})

}
