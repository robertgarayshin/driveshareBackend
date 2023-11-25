package service

import (
	"driveshare_backend/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// https://www.sohamkamani.com/golang/jwt-authentication/

type Claims struct {
	Id       int
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJwt(credentials models.LoginInfo) string {
	user, _ := GetProfileByEmail(credentials.Username)
	mySigningKey := []byte("AllYourBase")

	// Create claims with multiple fields populated
	claims := Claims{
		user.Id,
		user.Email,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)
	return ss
}
