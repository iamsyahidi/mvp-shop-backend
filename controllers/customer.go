package controllers

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"
	"mvp-shop-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerController struct {
	customerService services.CustomerServiceInterface
}

type CustomerControllerInterface interface {
	CreateCustomer(c *gin.Context)
	GetCustomerById(c *gin.Context)
	GetCustomers(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func NewCustomerController(customerService services.CustomerServiceInterface) CustomerControllerInterface {
	return &customerController{
		customerService: customerService,
	}
}

// CreateCustomer godoc
// @Summary Create a customer
// @Description Create a customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.CustomerRegister true "Customer"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 302 {object} models.Response
// @Router /customers [post]
func (cc *customerController) CreateCustomer(c *gin.Context) {
	var customerRegister models.CustomerRegister
	if err := c.ShouldBindJSON(&customerRegister); err != nil {
		middleware.Response(c, customerRegister, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	customer := models.Customer{
		Email:     customerRegister.Email,
		Name:      customerRegister.Name,
		Password:  customerRegister.Password,
		CreatedBy: customerRegister.Name,
	}

	response, err := cc.customerService.CreateCustomer(&customer)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, customerRegister, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, customerRegister, *response)
}

// GetCustomerById godoc
// @Summary Get a customer by id
// @Description Get a customer by id
// @Tags customers
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /customers/{id} [get]
func (cc *customerController) GetCustomerById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
		return
	}

	response, err := cc.customerService.GetCustomerById(id)
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

// GetCustomers godoc
// @Summary Get list customer
// @Description Get list customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param collection query []string false "string collection" collectionFormat(multi)
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /customers [get]
func (cc *customerController) GetCustomers(c *gin.Context) {
	filter := c.Request.URL.Query()
	response, err := cc.customerService.GetCustomers(filter)
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

// UpdateCustomer godoc
// @Summary Update a customer
// @Description Update a customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body models.CustomerUpdate true "Customer"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 302 {object} models.Response
// @Router /customers/{id} [put]
func (cc *customerController) UpdateCustomer(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		middleware.Response(c, "", models.Response{
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

	var customerUpdate models.CustomerUpdate
	if err := c.ShouldBindJSON(&customerUpdate); err != nil {
		middleware.Response(c, customerUpdate, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	customerUpdate.ID = id
	customerUpdate.UpdatedBy = v.(*models.CustomerClaims).Name
	response, err := cc.customerService.UpdateCustomer(&customerUpdate)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, customerUpdate, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, customerUpdate, *response)
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Delete a customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /customers/{id} [delete]
func (cc *customerController) DeleteCustomer(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		middleware.Response(c, "", models.Response{
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

	customerClaims, ok := v.(*models.CustomerClaims)
	if !ok {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to cast customer claims",
			Data:    nil,
		})
		return
	}

	customerDelete := models.CustomerUpdate{
		ID:        id,
		UpdatedBy: customerClaims.Name,
	}
	response, err := cc.customerService.DeleteCustomer(&customerDelete)
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
