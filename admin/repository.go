package admin

import (
	"github.com/RianIhsan/raise-unity/campaign"
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
	GetUserById(userId int) (user.User, error)
	GetCampaignById(campaignId int) (campaign.Campaign, error)
	FindUserById(userId int) (user.User, error)
	FindCampaignById(id int) (campaign.Campaign, error)
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

func (r *repository) GetUserById(userId int) (user.User, error) {
	var u user.User
	err := r.db.Where("id = ?", userId).Delete(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) FindUserById(ID int) (user.User, error) {
	var u user.User
	err := r.db.Where("id = ?", ID).Find(&u).Error
	if err != nil {
		return u, err
	}
	return u, err
}

func (r *repository) GetCampaignById(campaignId int) (campaign.Campaign, error) {
	var c campaign.Campaign
	err := r.db.Where("id = ?", campaignId).Delete(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}

func (r *repository) FindCampaignById(id int) (campaign.Campaign, error) {
	var c campaign.Campaign
	err := r.db.Where("id = ?", id).Find(&c).Error
	if err != nil {
		return c, err
	}

	return c, nil
}
