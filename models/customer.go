package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Customer struct {
	ID        string     `json:"id" gorm:"primary_key;not null;type:varchar(100);index"`
	Email     string     `json:"email" gorm:"not null;type:varchar(100);index"`
	Name      string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Password  string     `json:"password,omitempty" gorm:"type:varchar(150);index"`
	Status    Status     `json:"status" gorm:"not null;type:varchar(10);index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy string     `json:"created_by" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updated_by,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Customer) TableName() string {
	return "customer"
}

type CustomerRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3"`
	Password string `json:"password" binding:"required"`
}

type CustomerClaims struct {
	jwt.StandardClaims
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

type ListCustomer struct {
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
	Total     int        `json:"total"`
	TotalPage int        `json:"totalPage"`
	Customers []Customer `json:"customers"`
}

type CustomerUpdate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    Status `json:"status"`
	UpdatedBy string `json:"updated_by"`
}
