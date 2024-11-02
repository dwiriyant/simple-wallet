package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

type JWTClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint, username string) (string, error) {
	claims := &JWTClaims{
		ID:       userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
