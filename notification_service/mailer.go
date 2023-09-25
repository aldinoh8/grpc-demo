package main

import (
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Dialer *gomail.Dialer
}

func NewMailer(host, username, key string) Mailer {
	d := gomail.NewDialer(
		host,
		587,
		username,
		key,
	)

	return Mailer{Dialer: d}
}

func (m Mailer) GenerateMail(receiver, subject, body string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "ftgo@no-reply.com")
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	err := m.Dialer.DialAndSend(mail)
	return err
}
