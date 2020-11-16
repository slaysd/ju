package mail

import (
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

func (s *MailSender) Send(title string, message string, reference string) error {

	validate = validator.New()
	validate.Struct(s)

	from := s.Username
	to := s.Username
	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + title + "\r\n\r\n" +
		message + "\r\n\r\n" +
		reference

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err := smtp.SendMail(s.Host+":"+s.Port, auth, from, []string{to}, []byte(msg))

	return err
}
