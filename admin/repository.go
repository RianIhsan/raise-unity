package admin

import (
	"github.com/RianIhsan/raise-unity/user"
	"gorm.io/gorm"
)

type Repository interface {
	GetTotalUsers() (int64, error)
	GetPaginatedUsers(offset, limit int) ([]user.User, error)
	SearchUserByName(name string) ([]user.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTotalUsers() (int64, error) {
	var total int64
	if err := r.db.Model(&user.User{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
func (r *repository) GetPaginatedUsers(offset, limit int) ([]user.User, error) {
	var users []user.User
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) SearchUserByName(name string) ([]user.User, error) {
	var users []user.User
	err := r.db.Where("name LIKE ?", "%"+name+"%").
		Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}
