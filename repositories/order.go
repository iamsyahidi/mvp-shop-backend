package repositories

import (
	"fmt"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"
	"sync"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

type OrderRepositoryInterface interface {
	TransactionOrder(order *models.Order, orderDetail *[]models.OrderDetail, wg *sync.WaitGroup, mu *sync.Mutex) error
}

func NewOrderRepository(db *gorm.DB) OrderRepositoryInterface {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) TransactionOrder(order *models.Order, orderDetail *[]models.OrderDetail, wg *sync.WaitGroup, mu *sync.Mutex) error {
	defer wg.Done()

	// Acquire a lock to ensure safe concurrent access
	mu.Lock()
	defer mu.Unlock()

	// Begin a transaction
	tx := or.db.Begin()
	defer tx.Rollback()

	// Perform database operations within the transaction (use 'tx' from this point)
	if err := tx.Create(order).Error; err != nil {
		return fmt.Errorf("error creating order, %v", err)
	}

	if err := tx.Create(orderDetail).Error; err != nil {
		return fmt.Errorf("error creating order detail, %v", err)
	}

	for _, v := range *orderDetail {
		if err := tx.Model(&models.Product{}).Where(&models.Product{ID: v.ProductID}).Updates(map[string]interface{}{"stock": gorm.Expr("stock - ?", v.Qty)}).Error; err != nil {
			return fmt.Errorf("error updating stock product, %v", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction, %v", err)
	}

	return nil
}

func (or *orderRepository) GetMyOrders(pagination utils.Pagination, where map[string]string, userId string) ([]models.Order, int64, error) {
	var count int64
	var err error
	var sortField, sortDirection string
	var orders []models.Order

	queryBuilder := or.db.Model(&models.Order{}).Where("customer_id = ?", userId)

	if id, ok := where["id"]; ok && id != "" {
		queryBuilder = queryBuilder.Where("id = ?", id)
	}

	if invoice, ok := where["invoice"]; ok && invoice != "" {
		invoice := fmt.Sprintf("%%%s%%", invoice)
		queryBuilder = queryBuilder.Where(`"invoice" ILIKE ?`, invoice)
	}

	if pagination.SortField != "" {
		if pagination.SortField == "invoice" {
			sortField = `INITCAP("invoice")`
		} else {
			sortField = "created_at"
		}
	} else {
		sortField = "created_at"
	}

	if pagination.SortDirection != "" {
		sortDirection = pagination.SortDirection
	} else {
		sortDirection = models.SortDirectionDESC.String()
	}

	err = queryBuilder.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	orderBy := fmt.Sprintf("%s %s", sortField, sortDirection)
	limitBuilder := queryBuilder.Limit(pagination.Limit).Offset(offset).Order(orderBy)

	result := limitBuilder.Find(&orders)
	if result.Error != nil {
		return nil, count, result.Error
	}

	return orders, count, nil
}

func (or *orderRepository) GetMyOrderDetails(invoice string) ([]models.OrderDetail, error) {
	var orderDetail []models.OrderDetail

	result := or.db.Model(&models.OrderDetail{}).Where(&models.OrderDetail{Invoice: invoice}).Find(&orderDetail)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderDetail, nil
}
