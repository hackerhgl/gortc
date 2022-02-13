package gortc_models

import (
	"time"
)

type UserResetPasswordOTP struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null;index:user_code,unique" json:"userId"`
	Code      string    `gorm:"not null;index:user_code,unique; size:6" json:"code"`
	IsActive  bool      `gorm:"default:true" json:"isActive"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	// ExpireAt  time.Time `json:"-"`
	User User
}
