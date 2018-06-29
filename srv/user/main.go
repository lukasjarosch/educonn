package main

import (
	"os"

	"github.com/huandu/go-sqlbuilder"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	log "github.com/sirupsen/logrus"
)

const ExchangeName = "user"

func main() {

	var dbCfg DbConfig

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("1.0.0"),
		micro.Flags(
			cli.BoolFlag{
				Name:   "debug",
				EnvVar: "DEBUG",
				Usage:  "Enable debug mode",
			},
			cli.StringFlag{
				Name:   "db_host",
				EnvVar: "DB_HOST",
				Usage:  "The database host",
			},
			cli.StringFlag{
				Name:   "db_port",
				EnvVar: "DB_PORT",
				Usage:  "The database port",
			},
			cli.StringFlag{
				Name:   "db_username",
				EnvVar: "DB_USERNAME",
				Usage:  "The database user",
			},
			cli.StringFlag{
				Name:   "db_password",
				EnvVar: "DB_PASSWORD",
				Usage:  "The database password",
			},
			cli.StringFlag{
				Name:   "db_name",
				EnvVar: "DB_NAME",
				Usage:  "The database name to use",
			},
		),
	)

	// Parse command line flags
	srv.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("debug") {
				log.SetLevel(log.DebugLevel)
				log.Debug("DEBUG enabled")
			} else {
				log.SetLevel(log.InfoLevel)
			}
			dbCfg.Host = c.String("db_host")
			dbCfg.Port = c.String("db_port")
			dbCfg.User = c.String("db_username")
			dbCfg.Password = c.String("db_password")
			dbCfg.Name = c.String("db_name")
		}),
	)

	// Database
	db, err := CreateConnection(dbCfg)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	log.Debug("Database connection established")

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

	pub1 := micro.NewPublisher(userCreatedTopic, srv.Client())

	userStruct := sqlbuilder.NewStruct(new(User))
	repo := &UserRepository{db: db, userStruct: userStruct}
	tokenService := &TokenService{repo: repo}

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, pub1})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
