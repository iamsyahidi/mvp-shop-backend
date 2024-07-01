package controllers

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"
	"mvp-shop-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartController struct {
	cartService services.CartServiceInterface
}

type CartControllerInterface interface {
	CreateCart(c *gin.Context)
	UpdateCart(c *gin.Context)
	GetCartByCustomerID(c *gin.Context)
	DeleteCart(c *gin.Context)
}

func NewCartController(cartService services.CartServiceInterface) CartControllerInterface {
	return &cartController{
		cartService: cartService,
	}
}

// CreateCart godoc
// @Summary Create a cart
// @Description Create a cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param cart body models.CartRegister true "Cart"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 302 {object} models.Response
// @Router /carts [post]
func (cc *cartController) CreateCart(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		middleware.Response(c, "", models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	customer := v.(*models.CustomerClaims)
	var cartRegister models.CartRegister
	if err := c.ShouldBindJSON(&cartRegister); err != nil {
		middleware.Response(c, cartRegister, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	cart := models.Cart{
		CustomerID: customer.ID,
		ProductID:  cartRegister.ProductID,
		Qty:        cartRegister.Qty,
		Price:      cartRegister.Price,
		Status:     cartRegister.Status,
		CreatedBy:  customer.Email,
	}

	response, err := cc.cartService.CreateCart(&cart)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, cartRegister, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, cartRegister, *response)
}

// GetCartByCustomerID godoc
// @Summary Get a cart by id
// @Description Get a cart by id
// @Tags carts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Cart ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /carts/{id} [get]
func (cc *cartController) GetCartByCustomerID(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		middleware.Response(c, "", models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	customer := v.(*models.CustomerClaims)

	response, err := cc.cartService.GetCartByCustomerID(customer.ID)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, customer.ID, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, customer.ID, *response)
}

// UpdateCart godoc
// @Summary Update a cart
// @Description Update a cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param cart body models.CartUpdate true "Cart"
// @Success 201 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 302 {object} models.Response
// @Router /carts/{id} [put]
func (cc *cartController) UpdateCart(c *gin.Context) {
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

	var cartUpdate models.CartUpdate
	if err := c.ShouldBindJSON(&cartUpdate); err != nil {
		middleware.Response(c, cartUpdate, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	cartUpdate.ID = id
	cartUpdate.UpdatedBy = v.(*models.CustomerClaims).Name
	response, err := cc.cartService.UpdateCart(&cartUpdate)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, cartUpdate, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, cartUpdate, *response)
}

// DeleteCart godoc
// @Summary Delete a cart
// @Description Delete a cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Cart ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /carts/{id} [delete]
func (cc *cartController) DeleteCart(c *gin.Context) {
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

	customer, ok := v.(*models.CustomerClaims)
	if !ok {
		middleware.Response(c, id, models.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to cast cart claims",
			Data:    nil,
		})
		return
	}

	cartDelete := models.CartUpdate{
		ID:        id,
		UpdatedBy: customer.Name,
	}
	response, err := cc.cartService.DeleteCart(&cartDelete)
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
