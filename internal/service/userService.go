package service

import (
	"driveshare_backend/internal/database/read"
	"driveshare_backend/internal/models"
	"errors"
	"log"
	"net/http"
)

func GetProfileById(id int) (models.User, error) {
	return read.GetUserById(id)
}

func GetProfileByEmail(email string) (models.User, error) {
	return read.GetUserByEmail(email)
}

func SignIn(credentials models.LoginInfo, w http.ResponseWriter) (models.Payload, error) {
	user, _ := GetProfileByEmail(credentials.Username)
	if !CheckPasswordHash(credentials.Password, user.PasswordHash) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("user not found"))
		if err != nil {
			return models.Payload{}, err
		}
		return models.Payload{}, errors.New("incorrect login or password")
	}
	token := CreateJwt(credentials)
	log.Println(token)
	return models.Payload{
		Id:        user.Id,
		TokenType: "Bearer",
		AuthToken: token,
	}, nil
}
