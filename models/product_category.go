package models

import "time"

type ProductCategory struct {
	ID        string     `json:"id" gorm:"primary_key;not null;type:varchar(36);index"`
	Name      string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Status    Status     `json:"status" gorm:"not null;type:varchar(10);index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy string     `json:"created_by" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updated_by,omitempty" gorm:"type:varchar(150);default:null"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}

type ProductCategoryRegister struct {
	Name   string `json:"name" binding:"required,min=3"`
	Status Status `json:"status"`
}

type ListProductCategory struct {
	Page              int               `json:"page"`
	Limit             int               `json:"limit"`
	Total             int               `json:"total"`
	TotalPage         int               `json:"total_page"`
	ProductCategories []ProductCategory `json:"product_categories"`
}

type ProductCategoryUpdate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    Status `json:"status"`
	UpdatedBy string `json:"updated_by"`
}
