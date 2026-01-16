package main

import (
	"log"
	"os"
	"rubix-store/internal/config"
	"rubix-store/internal/handlers"
	"rubix-store/internal/middleware"
	"rubix-store/internal/repository"
	"rubix-store/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	r := setupRouter(authHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}

func setupRouter(authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	// API versioning
	v1 := r.Group("/api")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", authHandler.GetProfile)
			protected.PUT("/profile", authHandler.UpdateProfile)
		}

		// Admin routes (for future)
		// admin := v1.Group("/admin")
		// admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		// {
		// 	admin.GET("/users", ...)
		// }

		// Products routes (for future)
		// products := v1.Group("/products")
		// {
		// 	products.GET("/", ...)
		// 	products.GET("/:id", ...)
		// }

		// Cart routes (for future)
		// cart := v1.Group("/cart")
		// cart.Use(middleware.AuthMiddleware())
		// {
		// 	cart.GET("/", ...)
		// 	cart.POST("/items", ...)
		// }

		// Orders routes (for future)
		// orders := v1.Group("/orders")
		// orders.Use(middleware.AuthMiddleware())
		// {
		// 	orders.POST("/", ...)
		// 	orders.GET("/", ...)
		// }
	}

	return r
}
