package util

import (
    gomail "gopkg.in/mail.v2"
)

func SendMail(target string) error {
    message := gomail.NewMessage()

    message.SetHeader("From", "togelbos91@gmail.com")
    message.SetHeader("To", target)
    message.SetHeader("Subject", "Hello from Mikhael")
    message.SetBody("text/plain", "This is the Test Body")

    dialer := gomail.NewDialer("smtp.gmail.com", 587, "togelbos91@gmail.com", "Development123")

    if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
