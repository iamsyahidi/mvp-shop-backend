package services

import (
	"math"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"
	"mvp-shop-backend/repositories"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productCategoryService struct {
	productCategoryRepository repositories.ProductCategoryRepositoryInterface
}

type ProductCategoryServiceInterface interface {
	CreateProductCategory(productCategory *models.ProductCategory) (res *models.Response, err error)
	GetProductCategories(filter map[string][]string) (res *models.Response, err error)
	GetProductCategoryById(id string) (res *models.Response, err error)
	UpdateProductCategory(productCategory *models.ProductCategoryUpdate) (res *models.Response, err error)
	DeleteProductCategory(productCategory *models.ProductCategoryUpdate) (res *models.Response, err error)
}

func NewProductCategoryService(productCategoryRepository repositories.ProductCategoryRepositoryInterface) ProductCategoryServiceInterface {
	return &productCategoryService{
		productCategoryRepository: productCategoryRepository,
	}
}

func (ps *productCategoryService) CreateProductCategory(productCategory *models.ProductCategory) (res *models.Response, err error) {
	productCategory.ID = uuid.New().String()
	productCategory.CreatedBy = "admin"
	productCategory.Status = models.StatusActive
	err = ps.productCategoryRepository.CreateProductCategory(productCategory)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusCreated,
		Message: "ProductCategory created successfully",
	}, nil
}

func (ps *productCategoryService) GetProductCategories(filter map[string][]string) (res *models.Response, err error) {

	pagination, search := utils.GeneratePaginationFromRequest(filter)
	productCategories, count, err := ps.productCategoryRepository.GetProductCategories(pagination, search)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}, nil
	}

	data := models.ListProductCategory{
		Page:              pagination.Page,
		Limit:             pagination.Limit,
		Total:             int(count),
		TotalPage:         int(math.Ceil(float64(count) / float64(pagination.Limit))),
		ProductCategories: productCategories,
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "ProductCategory list successfully",
		Data:    data,
	}, nil
}

func (ps *productCategoryService) GetProductCategoryById(id string) (res *models.Response, err error) {

	productCategory, err := ps.productCategoryRepository.GetProductCategoryById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Response{
				Code:    http.StatusNotFound,
				Message: "ProductCategory not exist",
			}, nil
		}
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "ProductCategory get successfully",
		Data:    productCategory,
	}, nil
}

func (ps *productCategoryService) UpdateProductCategory(productCategory *models.ProductCategoryUpdate) (res *models.Response, err error) {
	err = ps.productCategoryRepository.UpdateProductCategory(productCategory)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "ProductCategory updated successfully",
	}, nil
}

func (ps *productCategoryService) DeleteProductCategory(productCategory *models.ProductCategoryUpdate) (res *models.Response, err error) {
	err = ps.productCategoryRepository.DeleteProductCategory(productCategory)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "ProductCategory deleted successfully",
	}, nil
}
