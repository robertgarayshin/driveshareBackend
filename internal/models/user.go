package models

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	ConfirmationToken string
	PasswordHash      string
}
