package main

import (
	"gopkg.in/gomail.v2"
	"github.com/lukasjarosch/educonn/srv/mail/proto"
	log "github.com/sirupsen/logrus"
)

type  SmtpConfig struct {
	Hostname string
	Port int
	Username string
	Password string
	Dialer *gomail.Dialer
}

func SendMail(config SmtpConfig, req* go_micro_srv_email.EmailRequest) {

	log.Debug("SendMail::start")

	mail := gomail.NewMessage()

	mail.SetHeader("From", req.From)
	mail.SetHeader("To", req.To)
	mail.SetHeader("Subject", req.Subject)
	mail.SetBody("text/html", req.Message)

	if err := config.Dialer.DialAndSend(mail); err != nil {
		log.Warn(err)
	}
	log.Debug("SendMail::end")
}

func NewDialer(config* SmtpConfig) {
	dialer := gomail.NewDialer(config.Hostname, config.Port, config.Username, config.Password)

	config.Dialer = dialer
}

