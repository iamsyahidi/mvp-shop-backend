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

type productService struct {
	productRepository repositories.ProductRepositoryInterface
}

type ProductServiceInterface interface {
	CreateProduct(product *models.Product) (res *models.Response, err error)
	GetProducts(filter map[string][]string) (res *models.Response, err error)
	GetProductById(id string) (res *models.Response, err error)
	UpdateProduct(product *models.ProductUpdate) (res *models.Response, err error)
	DeleteProduct(product *models.ProductUpdate) (res *models.Response, err error)
}

func NewProductService(productRepository repositories.ProductRepositoryInterface) ProductServiceInterface {
	return &productService{
		productRepository: productRepository,
	}
}

func (ps *productService) CreateProduct(product *models.Product) (res *models.Response, err error) {
	product.ID = uuid.New().String()
	product.CreatedBy = "admin"
	product.Status = models.StatusActive
	err = ps.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
	}, nil
}

func (ps *productService) GetProducts(filter map[string][]string) (res *models.Response, err error) {

	pagination, search := utils.GeneratePaginationFromRequest(filter)
	products, count, err := ps.productRepository.GetProducts(pagination, search)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}, nil
	}

	data := models.ListProduct{
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Total:     int(count),
		TotalPage: int(math.Ceil(float64(count) / float64(pagination.Limit))),
		Products:  products,
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Product list successfully",
		Data:    data,
	}, nil
}

func (ps *productService) GetProductById(id string) (res *models.Response, err error) {

	product, err := ps.productRepository.GetProductById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Response{
				Code:    http.StatusNotFound,
				Message: "Product not exist",
			}, nil
		}
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Product get successfully",
		Data:    product,
	}, nil
}

func (ps *productService) UpdateProduct(product *models.ProductUpdate) (res *models.Response, err error) {
	err = ps.productRepository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Product updated successfully",
	}, nil
}

func (ps *productService) DeleteProduct(product *models.ProductUpdate) (res *models.Response, err error) {
	err = ps.productRepository.DeleteProduct(product)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Product deleted successfully",
	}, nil
}
