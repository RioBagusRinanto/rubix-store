package config

import (
	"fmt"
	"log"
	"os"
	"rubix-store/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	if err := db.AutoMigrate(
		&models.User{},
		// Add more models as you develop:
		// &models.Product{},
		// &models.Category{},
		// &models.Cart{},
		// &models.CartItem{},
		// &models.Order{},
		// &models.OrderItem{},
		// &models.Address{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connection established and migrations completed")
	return db
}
