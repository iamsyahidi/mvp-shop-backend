package database

import (
	"context"
	"fmt"
	"log"
	"mvp-shop-backend/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm(ctx context.Context) (db *gorm.DB, err error) {

	db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				`host=%s port=%s user=%s dbname=%s sslmode=%s password=%s`,
				os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_PASSWORD"),
			),
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	//TODO add other db/model migrate
	db.AutoMigrate(
		&models.Customer{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderDetail{},
	)

	return db, nil
}
