package service

import (
	"driveshare_backend/configs"
	"fmt"
	"net/smtp"
	"strconv"
)

func SendVerificationEmail(id int, token string, to []string) {
	message := []byte("Subject: Email Verification\r\n" +
		"\r\n" +
		"Here is verification link http://localhost:8080/verify/" + strconv.Itoa(id) + "/token=" + token)

	auth := smtp.PlainAuth("", configs.Email, configs.Password, configs.SmtpHost)

	err := smtp.SendMail(configs.SmtpHost+":"+configs.SmtpPort, auth, configs.Email, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
