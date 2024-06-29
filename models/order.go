package models

import "time"

type Order struct {
	Invoice    string     `json:"invoice" gorm:"primary_key;not null;type:varchar(100);index"`
	CustomerID string     `json:"customer_id" gorm:"not null;type:varchar(100);index"`
	Amount     float64    `json:"amount" gorm:"index"`
	Payment    bool       `json:"payment" gorm:"not null;index;default:false"`
	Status     Status     `json:"status" gorm:"not null;type:varchar(10);index"`
	CreatedAt  time.Time  `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy  string     `json:"created_by" gorm:"not null;type:varchar(150)"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
	UpdatedBy  *string    `json:"updated_by,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderDetail struct {
	Invoice   string     `json:"invoice" gorm:"not null;type:varchar(100);index"`
	ProductID string     `json:"product_id" gorm:"not null;type:varchar(100);index"`
	Qty       float64    `json:"qty" gorm:"index"`
	Price     float64    `json:"price" gorm:"index"`
	Amount    float64    `json:"amount" gorm:"index"`
	Status    Status     `json:"status" gorm:"not null;type:varchar(10);index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy string     `json:"created_by" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updated_by,omitempty" gorm:"type:varchar(150);default:null"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}

type OrderRegister struct {
	Payment  bool           `json:"payment" binding:"required"`
	Products []OrderProduct `json:"products" binding:"required"`
}

type OrderProduct struct {
	ProductID string  `json:"product_id" binding:"required"`
	Qty       float64 `json:"qty" binding:"required"`
}
