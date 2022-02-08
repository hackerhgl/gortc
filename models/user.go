package gortc_models

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type UserRole string

const (
	RoleSuperAdmin UserRole = "super-admin"
	RoleAdmin      UserRole = "admin"
	RoleUser       UserRole = "user"
)

func (p *UserRole) Scan(value interface{}) error {
	*p = UserRole(value.([]byte))
	return nil
}

func (p UserRole) Value() (driver.Value, error) {
	return string(p), nil
}

var RolesArray []UserRole = []UserRole{RoleUser, RoleAdmin, RoleSuperAdmin}

type User struct {
	ID         uint         `gorm:"primarykey" json:"id"`
	Name       string       `gorm:"not null" json:"name"`
	Email      string       `gorm:"not null; unique; index" json:"email"`
	Password   string       `gorm:"not null" json:"-"`
	Salt       string       `gorm:"not null; size:12" json:"-" `
	IsActive   bool         `gorm:"default:true" json:"isActive"`
	IsVerified bool         `gorm:"default:false" json:"isVerified"`
	Role       UserRole     `gorm:"type:enum('super-admin', 'admin', 'user');default:'user'" json:"role"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
	DeletedAt  sql.NullTime `gorm:"index" json:"-"`
}
