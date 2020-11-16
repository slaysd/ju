package mail

import (
	"fmt"
	"net/smtp"

	"github.com/go-playground/validator"
)

var (
	validate *validator.Validate
)

type MailSender struct {
	Host     string `validate:"required,hostname"`
	Port     string `validate:"required"`
	Username string `validate:"required,email"`
	Password string `validate:"required"`
}

func (sender *MailSender) Send(title string, message string, reference string) {

	validate = validator.New()
	validate.Struct(sender)

	from := sender.Username
	to := sender.Username
	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + title + "\r\n\r\n" +
		message + "\r\n\r\n" +
		reference

	auth := smtp.PlainAuth("", sender.Username, sender.Password, sender.Host)
	err := smtp.SendMail(sender.Host+":"+sender.Port, auth, from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}
