package admin

import (
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"gorm.io/gorm"
)

type Repository interface {
	GetTotalUsers() (int64, error)
	GetPaginatedUsers(offset, limit int) ([]user.User, error)
	SearchUserByName(name string) ([]user.User, error)
	GetTotalTransactions() (int64, error)
	GetPaginatedTransactions(offset, limit int) ([]transaction.Transaction, error)
	SearchTransactionByUsername(name string) ([]transaction.Transaction, error)
	GetTotalTransactionsByUsername(name string) (int64, error)
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

func (r *repository) GetTotalTransactions() (int64, error) {
	var total int64
	if err := r.db.Model(&transaction.Transaction{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetTotalTransactionsByUsername(name string) (int64, error) {
	var userTransaction []transaction.Transaction
	var total int64
	if err := r.db.Preload("User").
		Preload("Campaign").
		Joins("JOIN users ON transactions.user_id = users.id").
		Where("users.name LIKE ?", "%"+name+"%").
		Count(&total).
		Find(&userTransaction).Error; err != nil {
		return 0, err
	}
	return total, nil
}
func (r *repository) GetPaginatedTransactions(offset, limit int) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	err := r.db.Offset(offset).Limit(limit).Preload("User").Preload("Campaign").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) SearchTransactionByUsername(name string) ([]transaction.Transaction, error) {
	var userTransaction []transaction.Transaction
	err := r.db.Preload("User").Preload("Campaign").Joins("JOIN users ON transactions.user_id = users.id").Where("users.name LIKE ?", "%"+name+"%").Find(&userTransaction).Error
	if err != nil {
		return userTransaction, err
	}
	return userTransaction, nil
}
