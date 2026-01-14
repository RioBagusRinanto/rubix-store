package handler

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"rubix-store/service/auth/repository"
	"rubix-store/service/auth/usecase"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func setupDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	sqlxdb := sqlx.NewDb(db, "sqlmock")
	return sqlxdb, mock, func() { sqlxdb.Close() }
}

func TestRegister_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, mock, cleanup := setupDB(t)
	defer cleanup()

	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))

	userRepo := repository.NewPostgresUserRepository(db)
	registerUC := usecase.NewRegisterUsecase(userRepo)
	loginUC := usecase.NewLoginUsecase(userRepo, []byte("secret"))
	handler := NewAuthHandler(registerUC, loginUC)

	router := gin.New()
	router.POST("/auth/register", handler.Register)

	body := map[string]string{"email": "test@example.com", "password": "password123"}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/auth/register", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, 201, w.Code)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, mock, cleanup := setupDB(t)
	defer cleanup()

	pw := []byte("password123")
	hash, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

	rows := sqlmock.NewRows([]string{"id", "email", "password_hash", "roles"}).
		AddRow(1, "test@example.com", string(hash), "user")
	mock.ExpectQuery("SELECT id, email, password_hash, roles FROM users WHERE email=\\$1").
		WithArgs("test@example.com").
		WillReturnRows(rows)

	userRepo := repository.NewPostgresUserRepository(db)
	registerUC := usecase.NewRegisterUsecase(userRepo)
	loginUC := usecase.NewLoginUsecase(userRepo, []byte("secret"))
	handler := NewAuthHandler(registerUC, loginUC)

	router := gin.New()
	router.POST("/auth/login", handler.Login)

	body := map[string]string{"email": "test@example.com", "password": "password123"}
	b, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
	var resp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NotEmpty(t, resp["token"])
	require.NoError(t, mock.ExpectationsWereMet())
}
