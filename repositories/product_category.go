package repositories

import (
	"fmt"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"

	"gorm.io/gorm"
)

type productCategoryRepository struct {
	db *gorm.DB
}

type ProductCategoryRepositoryInterface interface {
	CreateProductCategory(productCategory *models.ProductCategory) error
	GetProductCategories(pagination utils.Pagination, where map[string]string) ([]models.ProductCategory, int64, error)
	GetProductCategoryById(id string) (models.ProductCategory, error)
	UpdateProductCategory(productCategory *models.ProductCategoryUpdate) error
	DeleteProductCategory(productCategory *models.ProductCategoryUpdate) (err error)
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepositoryInterface {
	return &productCategoryRepository{
		db: db,
	}
}

func (pr *productCategoryRepository) CreateProductCategory(productCategory *models.ProductCategory) error {
	return pr.db.Create(productCategory).Error
}

func (pr *productCategoryRepository) GetProductCategories(pagination utils.Pagination, where map[string]string) ([]models.ProductCategory, int64, error) {
	var count int64
	var err error
	var sortField, sortDirection string
	var productCategorys []models.ProductCategory

	queryBuilder := pr.db.Model(&models.ProductCategory{}).Where("status <> ?", models.StatusDeleted)

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

	result := limitBuilder.Find(&productCategorys)
	if result.Error != nil {
		return nil, count, result.Error
	}

	return productCategorys, count, nil
}

func (pr *productCategoryRepository) GetProductCategoryById(id string) (models.ProductCategory, error) {
	var productCategory models.ProductCategory
	if err := pr.db.Where(&models.ProductCategory{ID: id}).First(&productCategory).Error; err != nil {
		return productCategory, err
	}
	return productCategory, nil
}

func (pr *productCategoryRepository) UpdateProductCategory(productCategory *models.ProductCategoryUpdate) error {
	return pr.db.
		Model(&models.ProductCategory{ID: productCategory.ID}).
		Updates(
			map[string]interface{}{
				"name":       productCategory.Name,
				"status":     productCategory.Status,
				"updated_at": gorm.Expr("now()"),
				"updated_by": productCategory.UpdatedBy,
			},
		).Error
}

func (pr *productCategoryRepository) DeleteProductCategory(productCategory *models.ProductCategoryUpdate) (err error) {
	return pr.db.
		Model(&models.ProductCategory{ID: productCategory.ID}).
		Updates(
			map[string]interface{}{
				"status":     models.StatusDeleted.String(),
				"updated_at": gorm.Expr("now()"),
				"updated_by": productCategory.UpdatedBy,
			},
		).Error
}
