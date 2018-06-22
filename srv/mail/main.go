package main

import (
	pb "github.com/lukasjarosch/educonn/srv/mail/proto"
	"github.com/micro/go-micro"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/micro/cli"
	"os"
	"gopkg.in/gomail.v2"
)

type service struct {
	config SmtpConfig
}

func (s *service) SendEmail(ctx context.Context, req *pb.EmailRequest, res *pb.Response) error {

	log.Debugf("EmailRequest received: TO='%s' SUBJECT='%s'", req.To, req.Subject)

	mail := gomail.NewMessage()
	mail.SetHeader("From", req.From)
	mail.SetHeader("To", req.To)
	mail.SetHeader("Subject", req.Subject)
	mail.SetBody("text/html", req.Message)

	if err := DialAndSendEmail(&s.config, *mail); err != nil {
		res.Code = 500
		res.Message = "EMAIL_SEND_FAIL"
		log.Warn(err)

		return nil
	}

	res.Code = 200
	res.Message = "EMAIL_SEND_SUCCESS"
	log.Debug("EmailRequest handled successfully, mail was sent!")

	return nil
}

func main() {

	log.SetLevel(log.InfoLevel)

	smtpConfig := SmtpConfig{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.mail"),
		micro.Version("1.0.0"),
		micro.Flags(
			cli.StringFlag{
				Name: "smtp_host",
				EnvVar: "SMTP_HOST",
				Usage: "The host where the STMP server is running",
			},
			cli.IntFlag{
				Name: "smtp_port",
				EnvVar: "SMTP_PORT",
				Usage: "Port of the SMTP server",
			},
			cli.StringFlag{
				Name: "smtp_password",
				EnvVar: "SMTP_PASSWORD",
				Usage: "Password for the SMTP server",
			},
			cli.StringFlag{
				Name: "smtp_username",
				EnvVar: "SMTP_USERNAME",
				Usage: "Username for the SMTP server",
			},
			cli.BoolFlag{
				Name: "debug",
				EnvVar: "DEBUG",
				Usage: "Enable debug mode. Disabled by default",
			},
		),
	)


	srv.Init(
		micro.Action(func(c *cli.Context) {
			smtpHost := c.String("smtp_host")
			smtpPort := c.Int("smtp_port")

			if c.Bool("debug") {
				log.SetLevel(log.DebugLevel)
				log.Debug("DEBUG ENABLED")
			}


			if smtpHost == "" {
				log.Errorf("SMTP_HOST not set. Cannot continue!")
				os.Exit(1)
			}

			if smtpPort == 0 {
				log.Debugf("SMTP_PORT not specified. Falling back to default: 465")
				smtpPort = 465
			}

			smtpConfig.Hostname = smtpHost
			smtpConfig.Port = smtpPort
			smtpConfig.Username = c.String("smtp_username")
			smtpConfig.Password = c.String("smtp_password")
		}),
	)

	log.Infof("SMTP is configured at: %s:%d", smtpConfig.Hostname, smtpConfig.Port)

	pb.RegisterEmailServiceHandler(srv.Server(), &service{smtpConfig})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
