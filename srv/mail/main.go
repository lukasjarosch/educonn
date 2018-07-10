package main

import (
	"context"
	"fmt"
	"os"

	user "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/rabbitmq"
	log "github.com/sirupsen/logrus"
)

type service struct {
	config SmtpConfig
}

type Sub struct{}

func (s *Sub) Process(ctx context.Context, event *user.UserCreatedEvent) error {
	log.Debugf("Received event: '%s': %+v", UserCreatedTopic, event)

	emailChan <- event
	log.Debugf("Queue ADD: userCreated: %s", event.Email)

	return nil
}

var (
	ExchangeName     = "user"
	UserCreatedTopic = "user.events.created"
	UserCreatedQueue = "user-created-queue"
	emailChan        = make(chan *user.UserCreatedEvent)
)

func main() {

	log.SetLevel(log.InfoLevel)
	defer close(emailChan)

	smtpConfig := SmtpConfig{}

	srv := micro.NewService(
		micro.Name("educonn.mail"),
		micro.Version("1.0.0"),
		micro.Flags(
			cli.StringFlag{
				Name:   "smtp_host",
				EnvVar: "SMTP_HOST",
				Usage:  "The host where the STMP server is running",
			},
			cli.IntFlag{
				Name:   "smtp_port",
				EnvVar: "SMTP_PORT",
				Usage:  "Port of the SMTP server",
			},
			cli.StringFlag{
				Name:   "smtp_password",
				EnvVar: "SMTP_PASSWORD",
				Usage:  "Password for the SMTP server",
			},
			cli.StringFlag{
				Name:   "smtp_username",
				EnvVar: "SMTP_USERNAME",
				Usage:  "Username for the SMTP server",
			},
			cli.BoolFlag{
				Name:   "debug",
				EnvVar: "DEBUG",
				Usage:  "Enable debug mode. Disabled by default",
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

	// SMTP
	log.Infof("SMTP is configured at: %s:%d", smtpConfig.Hostname, smtpConfig.Port)
	NewDialer(&smtpConfig)
	log.Debugf("SMTP Dialer created")

	// RabbitMQ
	broker := srv.Server().Options().Broker
	if err := broker.Init(rabbitmq.Exchange(ExchangeName)); err != nil {
		log.Fatalf("Unable to init broker: %v", err)
		os.Exit(1)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Unable to connect to broker: %v", err)
		os.Exit(1)
	}
	log.Debugf("Connected to RabbitMQ exchange '%s'", ExchangeName)

	micro.RegisterSubscriber(UserCreatedTopic, srv.Server(), new(Sub), server.SubscriberQueue(UserCreatedQueue))
	micro.Broker(broker)

	go func() {
		for email := range emailChan {
			SendMail(smtpConfig, email)
		}
	}()
	//// ----- ////

	/*
		pb.RegisterEmailServiceHandler(srv.Server(), &service{smtpConfig})

		log.Debugf("Starting mail queue...")

	*/

	//// ----- ////

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
