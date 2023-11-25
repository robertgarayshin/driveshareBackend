package read

import (
	"driveshare_backend/internal/models"
	"errors"
)

type UserRead struct {
	Id                int
	Email             string
	PasswordHash      string
	ConfirmationToken string
	IsConfirmed       bool
	Name              string
	Surname           string
}

var UserDbMock = []UserRead{
	{
		Id:                1,
		Email:             "rigarayshin@edu.hse.ru",
		PasswordHash:      "$2a$14$tjFK4aUP21sgJ2L/n7p5ne0Tzvr2ZZAvyIjHhxQbfUfaDcbEq.AGS",
		ConfirmationToken: "testtokentest",
		IsConfirmed:       false,
		Name:              "Robert",
		Surname:           "Garayshin",
	},
	{
		Id:                2,
		Email:             "test@user.org",
		PasswordHash:      "$2a$14$tjFK4aUP21sgJ2L/n7p5ne0Tzvr2ZZAvyIjHhxQbfUfaDcbEq.AGS",
		ConfirmationToken: "najsdkh2yugebv3271nl928js9i",
		IsConfirmed:       false,
		Name:              "Test",
		Surname:           "User",
	},
	{
		Id:                3,
		Email:             "john@doe.com",
		PasswordHash:      "$2a$14$tjFK4aUP21sgJ2L/n7p5ne0Tzvr2ZZAvyIjHhxQbfUfaDcbEq.AGS",
		ConfirmationToken: "cfdsokfvhjdifusogh7832bkdjcd",
		IsConfirmed:       false,
		Name:              "John",
		Surname:           "Doe",
	},
}

func GetUserById(id int) (models.User, error) {
	for i := range UserDbMock {
		if UserDbMock[i].Id == id {
			return models.User{
				Id:                UserDbMock[i].Id,
				Email:             UserDbMock[i].Email,
				Name:              UserDbMock[i].Name,
				Surname:           UserDbMock[i].Surname,
				ConfirmationToken: UserDbMock[i].ConfirmationToken,
				PasswordHash:      UserDbMock[i].PasswordHash,
			}, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func GetUserByEmail(email string) (models.User, error) {
	for i := range UserDbMock {
		if UserDbMock[i].Email == email {
			return models.User{
				Id:                UserDbMock[i].Id,
				Email:             UserDbMock[i].Email,
				Name:              UserDbMock[i].Name,
				Surname:           UserDbMock[i].Surname,
				ConfirmationToken: UserDbMock[i].ConfirmationToken,
				PasswordHash:      UserDbMock[i].PasswordHash,
			}, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func GetAllUsers() []models.User {
	var result []models.User
	for i := range UserDbMock {
		result = append(result, models.User{
			Id:                UserDbMock[i].Id,
			Email:             UserDbMock[i].Email,
			Name:              UserDbMock[i].Name,
			Surname:           UserDbMock[i].Surname,
			ConfirmationToken: UserDbMock[i].ConfirmationToken,
			PasswordHash:      UserDbMock[i].PasswordHash,
		})
	}
	return result
}
