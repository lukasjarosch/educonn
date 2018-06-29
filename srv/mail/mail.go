package main

import (
	"gopkg.in/gomail.v2"
	user "github.com/lukasjarosch/educonn/srv/user/proto/user"
	log "github.com/sirupsen/logrus"
	"fmt"
)

type  SmtpConfig struct {
	Hostname string
	Port int
	Username string
	Password string
	Dialer *gomail.Dialer
}

func SendMail(config SmtpConfig, event *user.UserCreatedEvent) {

	from := "hallo@educonn.de"
	subject := "Willkommen auf EduConn, %s"
	message := "Something great is about to happen..."

	mail := gomail.NewMessage()

	mail.SetHeader("From", from)
	mail.SetHeader("To", event.Email)
	mail.SetHeader("Subject", fmt.Sprintf(subject, event.FirstName))
	mail.SetBody("text/html", message)

	if err := config.Dialer.DialAndSend(mail); err != nil {
		log.Warn(err)
	}

	log.Debugf("Sent user-created-email to '%s'", event.Email)
}

func NewDialer(config* SmtpConfig) {
	dialer := gomail.NewDialer(config.Hostname, config.Port, config.Username, config.Password)

	config.Dialer = dialer
}

