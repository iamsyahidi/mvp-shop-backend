package routes

import (
	"mvp-shop-backend/controllers"
	"mvp-shop-backend/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(customerController controllers.CustomerControllerInterface, authController controllers.AuthControllerInterface) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	baseRouter := router.Group("/v1")

	//* customers
	customers := baseRouter.Group("/customers")
	customers.POST("", customerController.CreateCustomer)
	customersWithAuth := baseRouter.Group("/customers")
	customersWithAuth.Use(middleware.AuthMiddleware())
	customersWithAuth.GET("", customerController.GetCustomers)
	customersWithAuth.GET("/:id", customerController.GetCustomerById)
	customersWithAuth.PUT("/:id", customerController.UpdateCustomer)
	customersWithAuth.DELETE("/:id", customerController.DeleteCustomer)

	//* auth
	auth := baseRouter.Group("/auth")
	auth.POST("/login", authController.Login)

	return router
}
