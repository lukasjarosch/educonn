package main

import "gopkg.in/gomail.v2"

type  SmtpConfig struct {
	Hostname string
	Port int
	Username string
	Password string
}

func DialAndSendEmail(config* SmtpConfig, message gomail.Message) error {

	dialer := gomail.NewDialer(config.Hostname, config.Port, config.Username, config.Password)
	if err := dialer.DialAndSend(&message); err != nil {
		return err
	}
	return nil
}