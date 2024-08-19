package handler

import (
	"database/sql"
	"sj/internal/db/sqlc"
	"sj/internal/repository"
	"sj/internal/service"

	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine, uh *UserHandler) {
	r.GET("/auth/google/login-w-google", uh.LoginWithGoogle)
	r.GET("/auth/google/callback", uh.GetGoogleDetails)
	r.POST("/register", uh.RegisterUser)
	r.POST("/login", uh.Login)
}

func InitHandler(db *sql.DB) *UserHandler {
	queries := sqlc.New(db)

	userRepo := repository.NewUserRepository(queries)
	userServ := service.NewUserService(userRepo)
	userHand := NewUserHandler(userServ)

	transactionRepo := 
	
	return userHand
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	uh := InitHandler(db)
	route(r, uh)
}
