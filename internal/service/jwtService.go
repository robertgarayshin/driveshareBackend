package service

import (
	"driveshare_backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var JwtKey = []byte("key")

// https://www.sohamkamani.com/golang/jwt-authentication/

type Claims struct {
	Id       int
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJwt(credentials models.LoginInfo, w http.ResponseWriter) *jwt.Token {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Id:       1,
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token
}
