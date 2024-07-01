package controllers

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"
	"mvp-shop-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService services.ProductServiceInterface
}

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProductById(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

func NewProductController(productService services.ProductServiceInterface) ProductControllerInterface {
	return &productController{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.ProductRegister true "Product"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products [post]
func (pc *productController) CreateProduct(c *gin.Context) {
	var productRegister models.ProductRegister
	if err := c.ShouldBindJSON(&productRegister); err != nil {
		middleware.Response(c, productRegister, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	product := models.Product{
		Name:       productRegister.Name,
		Price:      productRegister.Price,
		Stock:      productRegister.Stock,
		CategoryID: productRegister.CategoryID,
		Status:     productRegister.Status,
	}

	response, err := pc.productService.CreateProduct(&product)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, productRegister, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, productRegister, *response)
}

// GetProducts godoc
// @Summary List a product
// @Description List a product
// @Tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param collection query []string false "string collection" collectionFormat(multi)
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products [get]
func (pc *productController) GetProducts(c *gin.Context) {
	filter := c.Request.URL.Query()
	response, err := pc.productService.GetProducts(filter)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, filter, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, filter, *response)
}

// GetProductById godoc
// @Summary Get a product by id
// @Description Get a product by id
// @Tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products/{id} [get]
func (pc *productController) GetProductById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
		return
	}

	response, err := pc.productService.GetProductById(id)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, id, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, id, *response)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.ProductUpdate true "Product"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 302 {object} models.Response
// @Router /products/{id} [put]
func (pc *productController) UpdateProduct(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		c.JSON(401, models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
		return
	}

	var productUpdate models.ProductUpdate
	if err := c.ShouldBindJSON(&productUpdate); err != nil {
		middleware.Response(c, productUpdate, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	productUpdate.ID = id
	productUpdate.UpdatedBy = v.(*models.CustomerClaims).Name
	response, err := pc.productService.UpdateProduct(&productUpdate)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, productUpdate, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, productUpdate, *response)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products/{id} [delete]
func (pc *productController) DeleteProduct(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		c.JSON(401, models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	id := c.Param("id")
	if id == "" {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
		return
	}

	productDelete := models.ProductUpdate{
		ID:        id,
		UpdatedBy: v.(*models.CustomerClaims).Name,
	}
	response, err := pc.productService.DeleteProduct(&productDelete)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, id, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, id, *response)
}
