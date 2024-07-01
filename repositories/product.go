package repositories

import (
	"fmt"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepositoryInterface interface {
	CreateProduct(product *models.Product) error
	GetProducts(pagination utils.Pagination, where map[string]string) ([]models.ProductView, int64, error)
	GetProductById(id string) (models.ProductView, error)
	UpdateProduct(product *models.ProductUpdate) error
	DeleteProduct(product *models.ProductUpdate) (err error)
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) CreateProduct(product *models.Product) error {
	return pr.db.Create(product).Error
}

func (pr *productRepository) GetProducts(pagination utils.Pagination, where map[string]string) ([]models.ProductView, int64, error) {
	var count int64
	var err error
	var sortField, sortDirection string
	var products []models.ProductView

	queryBuilder := pr.db.
		Table("products").Select("products.*, product_categories.name as category_name").
		Joins("left join product_categories on products.category_id = product_categories.id").
		Where("products.status <> ?", models.StatusDeleted)

	if id, ok := where["id"]; ok && id != "" {
		queryBuilder = queryBuilder.Where(`products.id = ?`, id)
	}

	if name, ok := where["name"]; ok && name != "" {
		name := fmt.Sprintf("%%%s%%", name)
		queryBuilder = queryBuilder.Where(`products."name" ILIKE ?`, name)
	}

	if categoryId, ok := where["category_id"]; ok && categoryId != "" {
		queryBuilder = queryBuilder.Where(`products.category_id = ?`, categoryId)
	}

	if categoryName, ok := where["category_name"]; ok && categoryName != "" {
		categoryName := fmt.Sprintf("%%%s%%", categoryName)
		queryBuilder = queryBuilder.Where(`product_categories."name" ILIKE ?`, categoryName)
	}

	if pagination.SortField != "" {
		if pagination.SortField == "name" {
			sortField = `INITCAP(products."name")`
		} else if pagination.SortField == "price" {
			sortField = "products.price"
		} else if pagination.SortField == "stock" {
			sortField = "products.stock"
		} else {
			sortField = "products.created_at"
		}
	} else {
		sortField = "products.created_at"
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

	result := limitBuilder.Scan(&products)
	if result.Error != nil {
		return nil, count, result.Error
	}

	return products, count, nil
}

func (pr *productRepository) GetProductById(id string) (models.ProductView, error) {
	var product models.ProductView
	queryBuilder := pr.db.
		Table("products").Select("products.*, product_categories.name as category_name").
		Joins("left join product_categories on products.category_id = product_categories.id").
		Where("products.status <> ?", models.StatusDeleted)
	if err := queryBuilder.Where("products.id = ?", id).Scan(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (pr *productRepository) UpdateProduct(product *models.ProductUpdate) error {
	return pr.db.
		Model(&models.Product{ID: product.ID}).
		Updates(
			map[string]interface{}{
				"name":        product.Name,
				"price":       product.Price,
				"stock":       product.Stock,
				"category_id": product.CategoryID,
				"status":      product.Status,
				"updated_at":  gorm.Expr("now()"),
				"updated_by":  product.UpdatedBy,
			},
		).Error
}

func (pr *productRepository) DeleteProduct(product *models.ProductUpdate) (err error) {
	return pr.db.
		Model(&models.Product{ID: product.ID}).
		Updates(
			map[string]interface{}{
				"status":     models.StatusDeleted.String(),
				"updated_at": gorm.Expr("now()"),
				"updated_by": product.UpdatedBy,
			},
		).Error
}
