package gortc_models

import (
	"time"
)

type InviteCode struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Code       string    `gorm:"not null;unique;size:6" json:"code"`
	IsActive   bool      `gorm:"default:true" json:"isActive"`
	CreatedBy  uint      `json:"createdBy"`
	RedeemedBy uint      `json:"redeemedBy"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	User       User
}
