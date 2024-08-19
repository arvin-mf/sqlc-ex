package jwt

import (
	"errors"
	"fmt"
	"os"
	"sj/internal/db/sqlc"
	"sj/internal/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(payload sqlc.User) (string, error) {
	expStr := os.Getenv("JWT_EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 1
	}
	tokenJwtTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.NewUserClaims(payload.ID, payload.Role, exp))
	tokenJwt, err := tokenJwtTemp.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}

func DecodeToken(signedToken string, ptrClaims jwt.Claims, KEY string) error {
	token, err := jwt.ParseWithClaims(signedToken, ptrClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", errors.New("wrong signging method")
		}
		return []byte(KEY), nil
	})
	if err != nil {
		return fmt.Errorf("token has been tampered with")
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
