package controllers

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"
	"mvp-shop-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService services.OrderServiceInterface
}

type OrderControllerInterface interface {
	CreateOrder(c *gin.Context)
}

func NewOrderController(orderService services.OrderServiceInterface) OrderControllerInterface {
	return &orderController{
		orderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create an order
// @Description Creates a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.OrderRegister true "Order"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /orders [post]
func (oc *orderController) CreateOrder(c *gin.Context) {
	v, ok := c.Get("customer")
	if !ok {
		middleware.Response(c, "", models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	var orderRegister models.OrderRegister
	if err := c.ShouldBindJSON(&orderRegister); err != nil {
		middleware.Response(c, orderRegister, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	order := models.Order{
		Payment:    orderRegister.Payment,
		CustomerID: v.(*models.CustomerClaims).ID,
		CreatedBy:  v.(*models.CustomerClaims).Name,
	}

	orderDetail := make([]models.OrderDetail, len(orderRegister.Products))
	for i, e := range orderRegister.Products {
		orderDetail[i] = models.OrderDetail{
			ProductID: e.ProductID,
			Qty:       e.Qty,
		}
	}

	response, err := oc.orderService.CreateOrder(&order, &orderDetail)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, orderRegister, models.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    nil,
		})
		return
	}

	middleware.Response(c, orderRegister, *response)
}
