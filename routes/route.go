package routes

import (
	"mvp-shop-backend/controllers"
	"mvp-shop-backend/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(customerController controllers.CustomerControllerInterface, authController controllers.AuthControllerInterface, productCategoryController controllers.ProductCategoryControllerInterface, productController controllers.ProductControllerInterface, cartController controllers.CartControllerInterface) *gin.Engine {
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

	//* products/categories
	productCategories := baseRouter.Group("/products/categories")
	productCategories.POST("", productCategoryController.CreateProductCategory)
	productCategoriesWithAuth := baseRouter.Group("/products/categories")
	productCategoriesWithAuth.Use(middleware.AuthMiddleware())
	productCategoriesWithAuth.GET("", productCategoryController.GetProductCategories)
	productCategoriesWithAuth.GET("/:id", productCategoryController.GetProductCategoryById)
	productCategoriesWithAuth.PUT("/:id", productCategoryController.UpdateProductCategory)
	productCategoriesWithAuth.DELETE("/:id", productCategoryController.DeleteProductCategory)

	//* products
	products := baseRouter.Group("/products")
	products.POST("", productController.CreateProduct)
	productsWithAuth := baseRouter.Group("/products")
	productsWithAuth.Use(middleware.AuthMiddleware())
	productsWithAuth.GET("", productController.GetProducts)
	productsWithAuth.GET("/:id", productController.GetProductById)
	productsWithAuth.PUT("/:id", productController.UpdateProduct)
	productsWithAuth.DELETE("/:id", productController.DeleteProduct)

	//* carts
	cartsWithAuth := baseRouter.Group("/carts")
	cartsWithAuth.Use(middleware.AuthMiddleware())
	cartsWithAuth.POST("", cartController.CreateCart)
	cartsWithAuth.GET("", cartController.GetCartByCustomerID)
	cartsWithAuth.DELETE("/:id", cartController.DeleteCart)

	return router
}
