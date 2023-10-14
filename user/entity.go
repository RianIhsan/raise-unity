package user

import "time"

type User struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Occupation string    `json:"occupation" gorm:"type:varchar(255)"`
	Email      string    `json:"email" gorm:"type:varchar(255)"`
	Password   string    `json:"password" gorm:"type:varchar(255)"`
	Avatar     string    `json:"avatar" gorm:"type:varchar(255)"`
	Role       string    `json:"role" gorm:"type:varchar(25)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
