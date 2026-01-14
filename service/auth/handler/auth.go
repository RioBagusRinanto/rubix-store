package handler

import (
	"net/http"

	"rubix-store/service/auth/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	registerUC *usecase.RegisterUsecase
	loginUC    *usecase.LoginUsecase
}

func NewAuthHandler(registerUC *usecase.RegisterUsecase, loginUC *usecase.LoginUsecase) *AuthHandler {
	return &AuthHandler{registerUC: registerUC, loginUC: loginUC}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Roles    string `json:"roles"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.registerUC.Execute(c.Request.Context(), req.Email, req.Password, req.Roles)
	if err != nil {
		switch err {
		case usecase.ErrUserExists:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case usecase.ErrHashPassword:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.loginUC.Execute(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		switch err {
		case usecase.ErrUserNotFound, usecase.ErrInvalidPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}
