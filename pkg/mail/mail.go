package mail

import (
	"fmt"
	"net/smtp"
)

const ()

type MailSender struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (sender *MailSender) Send(title string, message string, reference string) {
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
