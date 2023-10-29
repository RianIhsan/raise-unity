package admin

import (
	"github.com/RianIhsan/raise-unity/user"
	"math"
)

type Service interface {
	GetUsersPagination(page, pageSize int) ([]user.User, int, int, int, int, error)
	SearchUserByName(name string) ([]user.User, error)
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
