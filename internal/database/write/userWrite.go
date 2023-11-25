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
	read.DbMock = append(read.DbMock, res)
}

func ConfirmEmail(id int) {
	for i := range read.DbMock {
		if read.DbMock[i].Id == id {
			read.DbMock[i].IsConfirmed = true
		}
	}
}

func EditProfile(id int, user models.User) {
	for i := range read.DbMock {
		if read.DbMock[i].Id == id {
			read.DbMock[i].Name = user.Name
			read.DbMock[i].Email = user.Email
			read.DbMock[i].Surname = user.Surname
		}
	}
}

func DeleteProfile(id int) {
	for i := range read.DbMock {
		if read.DbMock[i].Id == id {
			read.DbMock[i] = read.DbMock[len(read.DbMock)-1]  // Copy last element to index i.
			read.DbMock[len(read.DbMock)-1] = read.UserRead{} // Erase last element (write zero value).
			read.DbMock = read.DbMock[:len(read.DbMock)-1]    // Truncate slice.
		}
	}
}
