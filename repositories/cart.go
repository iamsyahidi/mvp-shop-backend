package repositories

import (
	"mvp-shop-backend/models"

	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

type CartRepositoryInterface interface {
	CreateCart(cart *models.Cart) error
	GetCartByCustomerID(id string) (carts []models.ProductCartView, err error)
	DeleteCart(cart *models.CartUpdate) (err error)
}

func NewCartRepository(db *gorm.DB) CartRepositoryInterface {
	return &cartRepository{
		db: db,
	}
}

func (cr *cartRepository) CreateCart(cart *models.Cart) error {
	return cr.db.Create(cart).Error
}

func (cr *cartRepository) DeleteCart(cart *models.CartUpdate) (err error) {
	return cr.db.
		Model(&models.Cart{ID: cart.ID}).
		Updates(
			map[string]interface{}{
				"status":     models.StatusDeleted.String(),
				"updated_at": gorm.Expr("now()"),
				"updated_by": cart.UpdatedBy,
			},
		).Error
}

func (cr *cartRepository) GetCartByCustomerID(id string) (carts []models.ProductCartView, err error) {
	return carts, cr.db.
		Table("carts").Select("carts.*, products.name as name").
		Joins("left join products on carts.product_id = products.id").
		Where("carts.customer_id = ? and carts.status <> ?", id, models.StatusDeleted).Find(&carts).Error
}
