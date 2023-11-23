package service

import "driveshare_backend/internal/models"

func GetProfileById(id int) models.User {
	return models.User{
		Id:       0,
		Email:    "email",
		Password: "password",
		Name:     "name",
		Surname:  "surname",
	}
}

func GetProfileByEmail(email string) models.User {
	return models.User{
		Id:       0,
		Email:    "email",
		Password: "password",
		Name:     "name",
		Surname:  "surname",
	}
}
