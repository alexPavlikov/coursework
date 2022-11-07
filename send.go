package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func send(email string, pass string, name string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", cfg.Email)
	m.SetHeader("To", email)

	m.SetHeader("Subject", "Интернет магазин прикалдесов MyInvention.ru")

	// Set the email body. You can set plain text or html with text/html
	message := fmt.Sprintf(`Приветствуем, %s, поздравляем Вас с успешной регистрацией, в честь этого дарим вам скидочный купон 5 процентов на любой товар - 'NewUser'! Не забудьте ваш пароль - %s`, name, pass)
	m.SetBody("text/plain", message)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "a.pavlikov2002@gmail.com", "isei dkte iiwl wior")

	// This is only needed when the SSL/TLS certificate is not valid on the server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
