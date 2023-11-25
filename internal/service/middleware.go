package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Authorization"] != nil {
			tokenString := strings.Replace(request.Header["Authorization"][0], "Bearer ", "", 1)
			token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("AllYourBase"), nil
			})
			if err != nil {
				log.Fatal(err)
			} else if claims, ok := token.Claims.(*Claims); ok {
				fmt.Println(claims.Id, claims.RegisteredClaims.Issuer)
			} else {
				log.Fatal("unknown claims type, cannot proceed")
			}
			// if there's a token
			if token.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}
		// response for if there's no token header
	}
}
