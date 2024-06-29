package main

import (
	"context"
	"fmt"
	"log"
	"mvp-shop-backend/controllers"
	"mvp-shop-backend/pkg/database"
	"mvp-shop-backend/repositories"
	"mvp-shop-backend/routes"
	"mvp-shop-backend/services"
	"os"

	_ "mvp-shop-backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MVP Online Store API
// @description This is a small project for an online store server
// @version 1.0
// @host localhost:3001
// @BasePath /v1
// @schemes http
// @contact.name Ilham Syahidi
// @contact.email ilhamsyahidi66@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := database.InitGorm(ctx)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.DebugMode)

	// Repositories
	customerRepository := repositories.NewCustomerRepository(db)

	// Services
	customerService := services.NewCustomerService(customerRepository)
	authService := services.NewAuthService(customerRepository)

	// Controllers
	customerController := controllers.NewCustomerController(customerService)
	authController := controllers.NewAuthController(authService)

	router := routes.NewRouter(customerController, authController)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	router.Run(port)
}
