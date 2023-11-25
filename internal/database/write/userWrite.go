package write

import (
	"driveshare_backend/internal/database/read"
	"driveshare_backend/internal/models"
	"strconv"
	"time"
)

func CreateUser(user models.User) {
	token := strconv.Itoa(int(time.Now().Unix())) //TODO: добавить что-нибудь в base64
	res := read.UserRead{
		Id:                user.Id,
		Email:             user.Email,
		PasswordHash:      user.Password,
		ConfirmationToken: token,
		IsConfirmed:       false,
		Name:              user.Name,
		Surname:           user.Surname,
	}
	read.UserDbMock = append(read.UserDbMock, res)
}

func ConfirmEmail(id int) {
	for i := range read.UserDbMock {
		if read.UserDbMock[i].Id == id {
			read.UserDbMock[i].IsConfirmed = true
		}
	}
}

func EditProfile(id int, user models.User) {
	for i := range read.UserDbMock {
		if read.UserDbMock[i].Id == id {
			read.UserDbMock[i].Name = user.Name
			read.UserDbMock[i].Email = user.Email
			read.UserDbMock[i].Surname = user.Surname
		}
	}
}

func DeleteProfile(id int) {
	for i := range read.UserDbMock {
		if read.UserDbMock[i].Id == id {
			read.UserDbMock[i] = read.UserDbMock[len(read.UserDbMock)-1] // Copy last element to index i.
			read.UserDbMock[len(read.UserDbMock)-1] = read.UserRead{}    // Erase last element (write zero value).
			read.UserDbMock = read.UserDbMock[:len(read.UserDbMock)-1]   // Truncate slice.
		}
	}
}
