package user

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	SaveOTP(otp OTP) (OTP, error)
	FindValidOTP(userID int, otp string) (OTP, error)
	UpdateUser(user User) (User, error)
	DeleteOTP(otp OTP) error
	DeleteUserOTP(userID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SaveOTP(otp OTP) (OTP, error) {
	err := r.db.Create(&otp).Error
	if err != nil {
		return otp, err
	}
	return otp, nil
}

func (r *repository) FindValidOTP(userID int, otp string) (OTP, error) {
	var validOTP OTP
	err := r.db.Where("user_id = ? AND otp = ? AND expired_otp > ?", userID, otp, time.Now().Unix()).Find(&validOTP).Error
	if err != nil {
		return validOTP, err
	}

	return validOTP, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Model(&user).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteOTP(otp OTP) error {
	if err := r.db.Delete(&otp).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteUserOTP(userID int) error {
	err := r.db.Where("user_id = ?", userID).Delete(&OTP{}).Error
	if err != nil {
		return err
	}
	return nil
}
