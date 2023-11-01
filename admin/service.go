package admin

import (
	"errors"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"math"
)

type Service interface {
	GetUsersPagination(page, pageSize int) ([]user.User, int, int, int, int, error)
	SearchUserByName(name string) ([]user.User, error)
	GetTransactionsPagination(page, pageSize int) ([]transaction.Transaction, int, int, int, int, error)
	SearchTransactionByUsername(name string) ([]transaction.Transaction, error)
	DeleteUserById(id int) (user.User, error)
	DeleteCampaignById(id int) (campaign.Campaign, error)
	FindUserById(id int) (user.User, error)
	FindCampaignById(id int) (campaign.Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUsersPagination(page, pageSize int) ([]user.User, int, int, int, int, error) {
	totalUsers, err := s.repository.GetTotalUsers()
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalUsers) / float64(pageSize)))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}
	offset := (page - 1) * pageSize
	users, err := s.repository.GetPaginatedUsers(offset, pageSize)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	var nextPage, prevPage int
	if page < totalPages {
		nextPage = page + 1
	}
	if page > 1 {
		prevPage = page - 1
	}
	return users, totalPages, page, nextPage, prevPage, nil
}

func (s *service) SearchUserByName(name string) ([]user.User, error) {
	userByName, err := s.repository.SearchUserByName(name)
	if err != nil {
		return userByName, err
	}
	return userByName, nil
}

func (s *service) GetTransactionsPagination(page, pageSize int) ([]transaction.Transaction, int, int, int, int, error) {
	totalTransactions, err := s.repository.GetTotalTransactions()
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalTransactions) / float64(pageSize)))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}
	offset := (page - 1) * pageSize
	transactions, err := s.repository.GetPaginatedTransactions(offset, pageSize)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	var nextPage, prevPage int
	if page < totalPages {
		nextPage = page + 1
	}
	if page > 1 {
		prevPage = page - 1
	}
	return transactions, totalPages, page, nextPage, prevPage, nil
}

func (s *service) SearchTransactionByUsername(name string) ([]transaction.Transaction, error) {
	userTransaction, err := s.repository.SearchTransactionByUsername(name)
	if err != nil {
		return userTransaction, err
	}
	return userTransaction, nil
}

func (s *service) FindUserById(id int) (user.User, error) {
	user, err := s.repository.FindUserById(id)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}

	return user, err
}

func (s *service) DeleteUserById(id int) (user.User, error) {
	deleteUser, err := s.repository.GetUserById(id)
	if err != nil {
		return user.User{}, err
	}

	return deleteUser, nil
}

func (s *service) DeleteCampaignById(id int) (campaign.Campaign, error) {
	deleteCampaign, err := s.repository.GetCampaignById(id)
	if err != nil {
		return campaign.Campaign{}, err
	}
	return deleteCampaign, err
}

func (s *service) FindCampaignById(id int) (campaign.Campaign, error) {
	campaign, err := s.repository.FindCampaignById(id)
	if err != nil {
		return campaign, err
	}
	if campaign.ID == 0 {
		return campaign, errors.New("no campaign found with that id")
	}
	return campaign, nil
}
