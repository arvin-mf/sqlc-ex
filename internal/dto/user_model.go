package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserCreateReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	ID   string `json:"id" binding:"required"`
	Role string `json:"role" binding:"required"`
	jwt.RegisteredClaims
}

func NewUserClaims(id string, role string, exp time.Duration) UserClaims {
	return UserClaims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
