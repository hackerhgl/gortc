package gortc_models

import (
	"time"

	"gopkg.in/nullbio/null.v4"
)

type InviteCode struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Code       string    `gorm:"not null;unique;size:6" json:"code"`
	IsActive   bool      `gorm:"default:true" json:"isActive"`
	CreatedBy  uint      `gorm:"not null" json:"createdBy"`
	RedeemedBy null.Uint `json:"redeemedBy"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Created    User      `json:"-" gorm:"foreignKey:CreatedBy"`
	Redeemed   User      `json:"-" gorm:"foreignKey:RedeemedBy"`
}
