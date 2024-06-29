package models

import "time"

type Cart struct {
	ID         string     `json:"id" gorm:"primary_key;not null;type:varchar(36);index"`
	CustomerID string     `json:"customer_id" gorm:"not null;type:varchar(36);index"`
	ProductID  string     `json:"product_id" gorm:"not null;type:varchar(36);index"`
	Qty        float64    `json:"qty" gorm:"index"`
	Price      float64    `json:"price" gorm:"index"`
	Amount     float64    `json:"amount" gorm:"index"`
	Status     Status     `json:"status" gorm:"not null;type:varchar(10);index"`
	CreatedAt  time.Time  `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy  string     `json:"created_by" gorm:"not null;type:varchar(150)"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
	UpdatedBy  *string    `json:"updated_by,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Cart) TableName() string {
	return "carts"
}

type CartRegister struct {
	CustomerID string  `json:"customer_id" binding:"required"`
	ProductID  string  `json:"product_id" binding:"required"`
	Qty        float64 `json:"qty"`
	Price      float64 `json:"price"`
	Amount     float64 `json:"amount"`
	Status     Status  `json:"status"`
}

type CartUpdate struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	ProductID  string  `json:"product_id"`
	Qty        float64 `json:"qty"`
	Price      float64 `json:"price"`
	Amount     float64 `json:"amount"`
	Status     Status  `json:"status"`
	UpdatedBy  string  `json:"updated_by"`
}

type ProductCartView struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Qty       float64    `json:"qty"`
	Price     float64    `json:"price"`
	Amount    float64    `json:"amount"`
	Status    Status     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}

type CartView struct {
	Products    []ProductCartView `json:"products"`
	TotalAmount float64           `json:"total_amount"`
}
