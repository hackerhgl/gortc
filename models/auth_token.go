package gortc_models

import (
	"time"
)

type AuthToken struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"userId"`
	Token     string    `gorm:"not null" json:"token"`
	IsActive  bool      `gorm:"default:true" json:"isActive"`
	CreatedAt time.Time `json:"-"`
	User      User
}
