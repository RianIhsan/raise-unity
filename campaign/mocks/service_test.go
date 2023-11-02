package mocks

import (
	"errors"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestFindCampaigns(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	userID := 1
	campaigns := []campaign.Campaign{}

	t.Run("Valid find campaigns by user ID", func(t *testing.T) {

		repository.On("FindByUserID", userID).Return(campaigns, nil).Once()

		actualCampaigns, err := service.FindCampaigns(userID)

		assert.NoError(t, err)
		assert.Equal(t, campaigns, actualCampaigns)
		repository.AssertExpectations(t)
	})

	t.Run("Valid find campaigns without user ID", func(t *testing.T) {

		repository.On("FindAll").Return(campaigns, nil).Once()

		actualCampaigns, err := service.FindCampaigns(0)

		assert.NoError(t, err)
		assert.Equal(t, campaigns, actualCampaigns)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid find campaigns by user ID - user not found", func(t *testing.T) {

		repository.On("FindByUserID", userID).Return(nil, errors.New("no user found with that ID")).Once()

		actualCampaigns, err := service.FindCampaigns(userID)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "no user found with that ID")
		assert.Empty(t, actualCampaigns)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid find campaigns without user ID - error finding campaigns", func(t *testing.T) {

		repository.On("FindAll").Return(nil, errors.New("error finding campaigns")).Once()

		actualCampaigns, err := service.FindCampaigns(0)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "error finding campaigns")
		assert.Empty(t, actualCampaigns)
		repository.AssertExpectations(t)
	})
}

func TestFindCampaignByID(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	input := campaign.GetCampaignDetailInput{
		ID: 1,
	}
	c := campaign.Campaign{}

	t.Run("Valid find c by ID", func(t *testing.T) {

		repository.On("FindByID", input.ID).Return(c, nil).Once()

		actualCampaign, err := service.FindCampaignByID(input)

		assert.NoError(t, err)
		assert.Equal(t, c, actualCampaign)
		repository.AssertExpectations(t)
	})

	t.Run("Invalid find c by ID - c not found", func(t *testing.T) {

		repository.On("FindByID", input.ID).Return(c, errors.New("c not found")).Once()

		actualCampaign, err := service.FindCampaignByID(input)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "c not found")
		assert.Equal(t, c, actualCampaign)
		repository.AssertExpectations(t)
	})
}

func TestCreateCampaign(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	input := campaign.CreateCampaignInput{
		Name:             "Test Campaign",
		ShortDescription: "Short Description",
		Description:      "Campaign Description",
		GoalAmount:       1000,
		Perks:            "Perks",
		User:             user.User{ID: 1},
	}

	t.Run("Valid campaign creation", func(t *testing.T) {
		var c campaign.Campaign
		repository.On("Save", mock.Anything).Return(c, nil).Once()
		_, err := service.CreateCampaign(input)

		assert.NoError(t, err)
		assert.Equal(t, "unachieved", "unachieved")
		repository.AssertExpectations(t)
	})

	t.Run("Error during campaign creation", func(t *testing.T) {
		var c campaign.Campaign
		repository.On("Save", mock.Anything).Return(c, errors.New("error saving campaign")).Once()
		newCampaign, err := service.CreateCampaign(input)
		assert.Error(t, err)
		assert.Empty(t, newCampaign)
		repository.AssertExpectations(t)
	})
}

func TestUpdateCampaign(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	inputID := campaign.GetCampaignDetailInput{ID: 1}
	existingCampaign := campaign.Campaign{
		ID:               1,
		Name:             "Original Campaign",
		ShortDescription: "Original Description",
		Description:      "Original Description",
		GoalAmount:       1000,
		Perks:            "Original Perks",
		UserID:           1,
		Status:           "unachieved",
	}

	updatedData := campaign.CreateCampaignInput{
		Name:             "Updated Campaign",
		ShortDescription: "Updated Description",
		Description:      "Updated Description",
		GoalAmount:       2000,
		Perks:            "Updated Perks",
		User:             user.User{ID: 1},
	}

	t.Run("Valid campaign update", func(t *testing.T) {
		var updatedCampaign campaign.Campaign
		repository.On("FindByID", inputID.ID).Return(existingCampaign, nil).Once()
		repository.On("Update", mock.Anything).Return(updatedCampaign, nil).Once()

		updatedCampaign, err := service.UpdateCampaign(inputID, updatedData)

		assert.NoError(t, err)
		assert.NotEqual(t, updatedData.Name, updatedCampaign.Name)
		assert.NotEqual(t, updatedData.ShortDescription, updatedCampaign.ShortDescription)
		assert.NotEqual(t, updatedData.Description, updatedCampaign.Description)
		assert.NotEqual(t, updatedData.GoalAmount, updatedCampaign.GoalAmount)
		assert.NotEqual(t, updatedData.Perks, updatedCampaign.Perks)
		repository.AssertExpectations(t)
	})

	t.Run("Error during campaign update - campaign not found", func(t *testing.T) {
		repository.On("FindByID", inputID.ID).Return(campaign.Campaign{}, errors.New("campaign not found")).Once()

		updatedCampaign, err := service.UpdateCampaign(inputID, updatedData)

		assert.Error(t, err)
		assert.Empty(t, updatedCampaign)
		repository.AssertExpectations(t)
	})

	t.Run("Error during campaign update - not the owner", func(t *testing.T) {
		updatedData.User.ID = 2
		repository.On("FindByID", inputID.ID).Return(campaign.Campaign{}, errors.New("Not an owner of the campaign")).Once()

		updatedCampaign, err := service.UpdateCampaign(inputID, updatedData)

		assert.Error(t, err)
		assert.Equal(t, "Not an owner of the campaign", err.Error())
		assert.Empty(t, updatedCampaign)
		repository.AssertExpectations(t)
	})
}

func TestSaveCampaignImage(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	input := campaign.CreateCampaignImageInput{
		CampaignID: 1,
		IsPrimary:  true,
	}

	file := "test-image.jpg"

	t.Run("Valid campaign image creation - with marking all images as non-primary", func(t *testing.T) {
		repository.On("MarkAllImagesAsNonPrimary", input.CampaignID).Return(input.IsPrimary, nil).Once()

		newCampaignImage := campaign.CampaignImage{
			ID:         1,
			CampaignID: input.CampaignID,
			IsPrimary:  1,
			FileName:   file,
		}
		repository.On("CreateImage", mock.Anything).Return(newCampaignImage, nil).Once()

		createdImage, err := service.SaveCampaignImage(input, file)

		assert.NoError(t, err)
		assert.Equal(t, newCampaignImage, createdImage)

		repository.AssertExpectations(t)
	})

	t.Run("Valid campaign image creation - without marking all images as non-primary", func(t *testing.T) {
		repository.On("MarkAllImagesAsNonPrimary", input.CampaignID).Return(input.IsPrimary, nil).Once()

		newCampaignImage := campaign.CampaignImage{
			ID:         2,
			CampaignID: input.CampaignID,
			IsPrimary:  1,
			FileName:   file,
		}
		repository.On("CreateImage", mock.Anything).Return(newCampaignImage, nil).Once()

		createdImage, err := service.SaveCampaignImage(input, file)

		assert.NoError(t, err)
		assert.Equal(t, newCampaignImage, createdImage)

		repository.AssertExpectations(t)
	})

	t.Run("Error during campaign image creation - mark all images as non-primary", func(t *testing.T) {
		repository.On("MarkAllImagesAsNonPrimary", input.CampaignID).Return(input.IsPrimary, errors.New("error marking images as non-primary")).Once()

		createdImage, err := service.SaveCampaignImage(input, file)

		assert.Error(t, err)
		assert.Empty(t, createdImage)

		repository.AssertExpectations(t)
	})

	t.Run("Error during campaign image creation - create image error", func(t *testing.T) {
		repository.On("MarkAllImagesAsNonPrimary", input.CampaignID).Return(input.IsPrimary, nil).Once()

		repository.On("CreateImage", mock.Anything).Return(campaign.CampaignImage{}, errors.New("error creating campaign image")).Once()

		createdImage, err := service.SaveCampaignImage(input, file)

		assert.Error(t, err)
		assert.Empty(t, createdImage)

		repository.AssertExpectations(t)
	})
}

func TestGetPaginatedCampaigns(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	var mockCampaign = []campaign.Campaign{
		{
			ID:               1,
			UserID:           1,
			Name:             "Campaign Name",
			ShortDescription: "Short Description",
			Description:      "Campaign Description",
			Perks:            "Campaign Perks",
			BackerCount:      100,
			GoalAmount:       10000,
			CurrentAmount:    5000,
			Status:           "unachieved",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			CampaignImages: []campaign.CampaignImage{
				{
					ID:         1,
					CampaignID: 1,
					IsPrimary:  1,
					FileName:   "image1.jpg",
				},
				{
					ID:         2,
					CampaignID: 1,
					IsPrimary:  0,
					FileName:   "image2.jpg",
				},
			},
			User: user.User{
				ID:   1,
				Name: "rian",
			},
		},
		{
			ID:               2,
			UserID:           2,
			Name:             "Campaign Name",
			ShortDescription: "Short Description",
			Description:      "Campaign Description",
			Perks:            "Campaign Perks",
			BackerCount:      100,
			GoalAmount:       20000,
			CurrentAmount:    3000,
			Status:           "unachieved",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			CampaignImages: []campaign.CampaignImage{
				{
					ID:         1,
					CampaignID: 2,
					IsPrimary:  1,
					FileName:   "image1.jpg",
				},
				{
					ID:         2,
					CampaignID: 2,
					IsPrimary:  0,
					FileName:   "image2.jpg",
				},
			},
			User: user.User{
				ID:   2,
				Name: "ihsan",
			},
		},
	}

	t.Run("Success get pagination campaign", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalCampaigns").Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaigns", 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Get total campaign failed", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalCampaigns").Return(int64(0), errors.New("get total campaign failed")).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get total campaign failed")
		assert.Nil(t, campaigns)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Get paginated campaign failed", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalCampaigns").Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaigns", 0, pageSize).Return(nil, errors.New("get paginated campaigns error")).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get paginated campaigns error")
		assert.Nil(t, campaigns)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
	})

	t.Run("page = 0", func(t *testing.T) {
		var page = 0
		var pageSize = 10

		repository.On("GetTotalCampaigns").Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaigns", 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 5
		var pageSize = 10

		repository.On("GetTotalCampaigns").Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaigns", 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 1
		var pageSize = 1

		repository.On("GetTotalCampaigns").Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaigns", 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > 1", func(t *testing.T) {
		var page = 2
		var pageSize = 1

		repository.On("GetTotalCampaigns").Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaigns", 1, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaigns(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 1, prevPage)
		repository.AssertExpectations(t)
	})
}

func TestPaginatedCampaignsByUserID(t *testing.T) {
	repository := NewRepository(t)
	service := campaign.NewService(repository)

	var mockCampaign = []campaign.Campaign{
		{
			ID:               1,
			UserID:           1,
			Name:             "Campaign Name",
			ShortDescription: "Short Description",
			Description:      "Campaign Description",
			Perks:            "Campaign Perks",
			BackerCount:      100,
			GoalAmount:       10000,
			CurrentAmount:    5000,
			Status:           "unachieved",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			CampaignImages: []campaign.CampaignImage{
				{
					ID:         1,
					CampaignID: 1,
					IsPrimary:  1,
					FileName:   "image1.jpg",
				},
				{
					ID:         2,
					CampaignID: 1,
					IsPrimary:  0,
					FileName:   "image2.jpg",
				},
			},
			User: user.User{
				ID:   1,
				Name: "rian",
			},
		},
		{
			ID:               2,
			UserID:           2,
			Name:             "Campaign Name",
			ShortDescription: "Short Description",
			Description:      "Campaign Description",
			Perks:            "Campaign Perks",
			BackerCount:      100,
			GoalAmount:       20000,
			CurrentAmount:    3000,
			Status:           "unachieved",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			CampaignImages: []campaign.CampaignImage{
				{
					ID:         1,
					CampaignID: 2,
					IsPrimary:  1,
					FileName:   "image1.jpg",
				},
				{
					ID:         2,
					CampaignID: 2,
					IsPrimary:  0,
					FileName:   "image2.jpg",
				},
			},
			User: user.User{
				ID:   2,
				Name: "ihsan",
			},
		},
	}

	t.Run("Success get pagination campaign by user id", func(t *testing.T) {
		var page = 1
		var pageSize = 10
		var userId = 1
		var totalCampaigns = 20
		repository.On("GetTotalCampaignsByUserID", userId).Return(int64(totalCampaigns), nil).Once()
		repository.On("GetPaginatedCampaignsByUserID", userId, 0, pageSize).Return(mockCampaign, nil).Once()

		resultCampaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaignsByUserID(userId, page, pageSize)

		assert.NoError(t, err)
		assert.Equal(t, mockCampaign, resultCampaigns)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, page, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)

		repository.AssertExpectations(t)

	})

	t.Run("Error getting total campaigns count", func(t *testing.T) {
		var page = 1
		var pageSize = 10
		var userId = 1
		repository.On("GetTotalCampaignsByUserID", userId).Return(int64(0), errors.New("error getting total count")).Once()

		resultCampaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaignsByUserID(userId, page, pageSize)

		assert.Error(t, err)
		assert.Empty(t, resultCampaigns)
		assert.Equal(t, 0, totalPages)
		assert.Equal(t, 0, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)

		repository.AssertExpectations(t)
	})

	t.Run("page = 0", func(t *testing.T) {
		var page = 0
		var pageSize = 10
		var userID = 1

		repository.On("GetTotalCampaignsByUserID", userID).Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaignsByUserID", userID, 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaignsByUserID(userID, page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalpage", func(t *testing.T) {
		var page = 5
		var pageSize = 10
		var userID = 1

		repository.On("GetTotalCampaignsByUserID", userID).Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaignsByUserID", userID, 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaignsByUserID(userID, page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 1
		var pageSize = 1

		var userID = 1

		repository.On("GetTotalCampaignsByUserID", userID).Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaignsByUserID", userID, 0, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaignsByUserID(userID, page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > 1", func(t *testing.T) {
		var page = 2
		var pageSize = 1
		var userID = 1

		repository.On("GetTotalCampaignsByUserID", userID).Return(int64(2), nil).Once()
		repository.On("GetPaginatedCampaignsByUserID", userID, 1, pageSize).Return(mockCampaign, nil).Once()

		campaigns, totalPages, currentPage, nextPage, prevPage, err := service.GetPaginatedCampaignsByUserID(userID, page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, campaigns)
		assert.Equal(t, mockCampaign[0].Name, campaigns[0].Name)
		assert.Equal(t, mockCampaign[1].Name, campaigns[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 1, prevPage)
		repository.AssertExpectations(t)

	})
}

func TestSearchCampaignsByName(t *testing.T) {
	var repository = NewRepository(t)
	var service = campaign.NewService(repository)

	var mockCampaign = []campaign.Campaign{
		{
			ID:               1,
			UserID:           1,
			Name:             "Campaign Name",
			ShortDescription: "Short Description",
			Description:      "Campaign Description",
			Perks:            "Campaign Perks",
			BackerCount:      100,
			GoalAmount:       10000,
			CurrentAmount:    5000,
			Status:           "unachieved",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			CampaignImages: []campaign.CampaignImage{
				{
					ID:         1,
					CampaignID: 1,
					IsPrimary:  1,
					FileName:   "image1.jpg",
				},
				{
					ID:         2,
					CampaignID: 1,
					IsPrimary:  0,
					FileName:   "image2.jpg",
				},
			},
			User: user.User{
				ID:   1,
				Name: "rian",
			},
		},
	}

	t.Run("Success get campaign list by name", func(t *testing.T) {
		repository.On("SearchCampaignsByName", "rian").Return(mockCampaign, nil).Once()

		result, err := service.SearchCampaignsByName("rian")
		assert.Nil(t, err)
		assert.Equal(t, mockCampaign, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed get campaign list by username", func(*testing.T) {
		repository.On("SearchCampaignsByName", "rian").Return(nil, errors.New("Failed Get campaign user by name")).Once()

		result, err := service.SearchCampaignsByName("rian")
		assert.Error(t, err)
		assert.EqualError(t, err, "Failed Get campaign user by name")
		assert.Nil(t, result)
		repository.AssertExpectations(t)
	})
}
