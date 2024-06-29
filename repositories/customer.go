package repositories

import (
	"fmt"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

type CustomerRepositoryInterface interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomerByEmail(email string) (models.Customer, error)
	GetCustomerById(id string) (models.Customer, error)
	UpdateCustomer(customer *models.CustomerUpdate) error
	GetCustomers(pagination utils.Pagination, where map[string]string) (customers []models.Customer, count int64, err error)
	DeleteCustomer(customer *models.CustomerUpdate) (err error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepositoryInterface {
	return &customerRepository{
		db: db,
	}
}

func (cr *customerRepository) CreateCustomer(customer *models.Customer) error {
	return cr.db.Create(customer).Error
}

func (cr *customerRepository) GetCustomerByEmail(email string) (models.Customer, error) {
	var customer models.Customer
	if err := cr.db.Where(&models.Customer{Email: email}).First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (cr *customerRepository) GetCustomerById(id string) (models.Customer, error) {
	var customer models.Customer
	if err := cr.db.
		Where(&models.Customer{ID: id}).
		Select("id", "email", "name", "status", "created_at", "created_by", "updated_at", "updated_by").
		First(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (cr *customerRepository) UpdateCustomer(customer *models.CustomerUpdate) error {
	return cr.db.
		Model(&models.Customer{ID: customer.ID}).
		Updates(
			map[string]interface{}{
				"name":       customer.Name,
				"status":     customer.Status,
				"updated_at": gorm.Expr("now()"),
				"updated_by": customer.UpdatedBy,
			},
		).Error
}

func (cr *customerRepository) GetCustomers(pagination utils.Pagination, where map[string]string) (customers []models.Customer, count int64, err error) {

	var (
		sortField, sortDirection string
	)

	queryBuilder := cr.db.
		Model(&models.Customer{}).
		Where("status <> ?", models.StatusDeleted).
		Select("id", "email", "name", "status", "created_at", "created_by", "updated_at", "updated_by")

	if id, ok := where["id"]; ok && id != "" {
		queryBuilder = queryBuilder.Where("id = ?", id)
	}

	if name, ok := where["name"]; ok && name != "" {
		name := fmt.Sprintf("%%%s%%", name)
		queryBuilder = queryBuilder.Where(`"name" ILIKE ?`, name)
	}

	if pagination.SortField != "" {
		if pagination.SortField == "name" {
			sortField = `INITCAP("name")`
		} else if pagination.SortField == "email" {
			sortField = `INITCAP("email")`
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

	result := limitBuilder.Find(&customers)
	if result.Error != nil {
		return nil, count, result.Error
	}

	return customers, count, nil
}

func (cr *customerRepository) DeleteCustomer(customer *models.CustomerUpdate) (err error) {
	return cr.db.
		Model(&models.Customer{ID: customer.ID}).
		Updates(
			map[string]interface{}{
				"status":     models.StatusDeleted.String(),
				"updated_at": gorm.Expr("now()"),
				"updated_by": customer.UpdatedBy,
			},
		).Error
}
