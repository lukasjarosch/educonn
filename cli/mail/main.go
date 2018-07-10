package main

import (
	"context"
	pb "github.com/lukasjarosch/educonn/srv/mail/proto"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)

	client := pb.NewEmailServiceClient("educonn.mail", microclient.DefaultClient)

	service := micro.NewService(
		micro.Name("go.micro.cli.mail"),
		micro.Version("latest"),

		// Define some flags to set
		micro.Flags(
			cli.StringFlag{
				Name:   "to",
				EnvVar: "TO",
				Usage:  "Recieving address",
			},
			cli.StringFlag{
				Name:  "subject",
				Usage: "Email subject",
			},
			cli.StringFlag{
				Name:  "from",
				Usage: "From email",
			},
			cli.StringFlag{
				Name:  "body",
				Usage: "Email message body (html enabled)",
			},
		),

		// Evaluate and fire EmailRequest
		micro.Action(func(c *cli.Context) {
			to := c.String("to")
			from := c.String("from")
			subject := c.String("subject")
			body := c.String("body")

			_, err := client.SendEmail(context.TODO(), &pb.EmailRequest{
				To:      to,
				From:    from,
				Subject: subject,
				Message: body,
			})

			if err != nil {
				log.Fatalf("Unable to send mail: %v", err)
			}
			os.Exit(0)
		}),
	)

	// Go!
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
