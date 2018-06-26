package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	log "github.com/sirupsen/logrus"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"os"
)

func main() {

	var dbCfg DbConfig

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("1.0.0"),
		micro.Flags(
			cli.BoolFlag{
				Name: "debug",
				EnvVar: "DEBUG",
				Usage: "Enable debug mode",
			},
			cli.StringFlag{
				Name: "db_host",
				EnvVar: "DB_HOST",
				Usage: "The database host",
			},
			cli.StringFlag{
				Name: "db_port",
				EnvVar: "DB_PORT",
				Usage: "The database port",
			},
			cli.StringFlag{
				Name: "db_username",
				EnvVar: "DB_USERNAME",
				Usage: "The database user",
			},
			cli.StringFlag{
				Name: "db_password",
				EnvVar: "DB_PASSWORD",
				Usage: "The database password",
			},
			cli.StringFlag{
				Name: "db_name",
				EnvVar: "DB_NAME",
				Usage: "The database name to use",
			},
		),
	)

	// Parse command line flags
	srv.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("debug")	 {
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

	db, err := CreateConnection(dbCfg)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	log.Debug("Database connection established")

	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db: db}
	tokenService := &TokenService{repo: repo}

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
