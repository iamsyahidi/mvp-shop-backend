package controllers

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"
	"mvp-shop-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productCategoryController struct {
	productCategoryService services.ProductCategoryServiceInterface
}

type ProductCategoryControllerInterface interface {
	CreateProductCategory(c *gin.Context)
	GetProductCategories(c *gin.Context)
	GetProductCategoryById(c *gin.Context)
	UpdateProductCategory(c *gin.Context)
	DeleteProductCategory(c *gin.Context)
}

func NewProductCategoryController(productCategoryService services.ProductCategoryServiceInterface) ProductCategoryControllerInterface {
	return &productCategoryController{
		productCategoryService: productCategoryService,
	}
}

// CreateProductCategory godoc
// @Summary Create a productCategory
// @Description Create a productCategory
// @Tags productCategories
// @Accept  json
// @Produce  json
// @Param productCategory body models.ProductCategoryRegister true "ProductCategory"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products/categories [post]
func (pc *productCategoryController) CreateProductCategory(c *gin.Context) {
	var productCategoryRegister models.ProductCategoryRegister
	if err := c.ShouldBindJSON(&productCategoryRegister); err != nil {
		middleware.Response(c, productCategoryRegister, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	productCategory := models.ProductCategory{
		Name: productCategoryRegister.Name,
	}

	response, err := pc.productCategoryService.CreateProductCategory(&productCategory)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, productCategoryRegister, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, productCategoryRegister, *response)
}

// GetProductCategories godoc
// @Summary List a productCategory
// @Description List a productCategory
// @Tags productCategories
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param collection query []string false "string collection" collectionFormat(multi)
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products/categories [get]
func (pc *productCategoryController) GetProductCategories(c *gin.Context) {
	filter := c.Request.URL.Query()
	response, err := pc.productCategoryService.GetProductCategories(filter)
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

// GetProductCategoryById godoc
// @Summary Get a productCategory by id
// @Description Get a productCategory by id
// @Tags productCategories
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products/categories/{id} [get]
func (pc *productCategoryController) GetProductCategoryById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
		return
	}

	response, err := pc.productCategoryService.GetProductCategoryById(id)
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

// UpdateProductCategory godoc
// @Summary Update a productCategory
// @Description Update a productCategory
// @Tags productCategories
// @Accept  json
// @Produce  json
// @Param productCategory body models.ProductCategoryUpdate true "ProductCategory"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 302 {object} models.Response
// @Router /products/categories/{id} [put]
func (pc *productCategoryController) UpdateProductCategory(c *gin.Context) {
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

	var productCategoryUpdate models.ProductCategoryUpdate
	if err := c.ShouldBindJSON(&productCategoryUpdate); err != nil {
		middleware.Response(c, productCategoryUpdate, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	productCategoryUpdate.ID = id
	productCategoryUpdate.UpdatedBy = v.(*models.CustomerClaims).Name
	response, err := pc.productCategoryService.UpdateProductCategory(&productCategoryUpdate)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, productCategoryUpdate, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, productCategoryUpdate, *response)
}

// DeleteProductCategory godoc
// @Summary Delete a productCategory
// @Description Delete a productCategory
// @Tags productCategories
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products/categories/{id} [delete]
func (pc *productCategoryController) DeleteProductCategory(c *gin.Context) {
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

	productCategoryDelete := models.ProductCategoryUpdate{
		ID:        id,
		UpdatedBy: v.(*models.CustomerClaims).Name,
	}
	response, err := pc.productCategoryService.DeleteProductCategory(&productCategoryDelete)
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
