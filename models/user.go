package gortc_models

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type userRole string

const (
	RoleSuperAdmin userRole = "super-admin"
	RoleAdmin      userRole = "admin"
	Roleuser       userRole = "user"
)

func (p *userRole) Scan(value interface{}) error {
	*p = userRole(value.([]byte))
	return nil
}

func (p userRole) Value() (driver.Value, error) {
	return string(p), nil
}

type User struct {
	ID         uint         `gorm:"primarykey" json:"id"`
	Name       string       `gorm:"not null" json:"name"`
	Email      string       `gorm:"not null; unique; index" json:"email"`
	Password   string       `gorm:"not null" json:"-"`
	Salt       string       `gorm:"not null; size:12" json:"-" `
	IsVerified bool         `gorm:"default:false" json:"isVerified"`
	Role       userRole     `gorm:"type:enum('super-admin', 'admin', 'user');default:'user'" json:"role"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
	DeletedAt  sql.NullTime `gorm:"index" json:"-"`
}
