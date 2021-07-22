package gortc_models

import (
	"database/sql/driver"

	"gorm.io/gorm"
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
	gorm.Model
	Name       string   `gorm:"not null"`
	Email      string   `gorm:"not null"`
	Password   string   `json:"-" gorm:"not null"`
	Salt       string   `gorm:"not null; size:12" json:"-" `
	IsVerified bool     `gorm:"default:false"`
	Role       userRole `gorm:"type:enum('super-admin', 'admin', 'user');default:'user'"`
}
