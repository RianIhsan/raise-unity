package user

import (
	"time"
)

type User struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Occupation string    `json:"occupation" gorm:"type:varchar(255)"`
	Email      string    `json:"email" gorm:"type:varchar(255);unique"`
	Password   string    `json:"password" gorm:"type:varchar(255)"`
	Avatar     string    `json:"avatar" form:"avatar" gorm:"type:varchar(255)"`
	Role       string    `json:"role" gorm:"type:varchar(25)"`
	IsVerified bool      `json:"is_verified" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OTP struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     int       `json:"user_id" gorm:"index;unique"`
	User       User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	OTP        string    `json:"otp" gorm:"type:varchar(255)"`
	ExpiredOTP int64     `json:"expired_otp" gorm:"type:bigint"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
