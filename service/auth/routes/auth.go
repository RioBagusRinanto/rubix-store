package routes

import (
	"os"

	"rubix-store/service/auth/handler"
	"rubix-store/service/auth/middleware"
	"rubix-store/service/auth/repository"
	"rubix-store/service/auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRoutes(r *gin.Engine, db *sqlx.DB) {
	// Infrastructure
	userRepo := repository.NewPostgresUserRepository(db)

	// Use Cases
	registerUC := usecase.NewRegisterUsecase(userRepo)
	loginUC := usecase.NewLoginUsecase(userRepo, getJWTSecret())

	// Handler
	authHandler := handler.NewAuthHandler(registerUC, loginUC)

	// Routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Protected routes
	secret := getJWTSecret()
	meGroup := r.Group("/me")
	meGroup.Use(middleware.JWTMiddleware(secret))
	{
		meGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user_id": c.GetInt("user_id"),
				"email":   c.GetString("user_email"),
			})
		})
	}
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret"
	}
	return []byte(secret)
}
