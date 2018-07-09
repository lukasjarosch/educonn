package main

import (
	"github.com/lukasjarosch/educonn/srv/course/handler"
	course "github.com/lukasjarosch/educonn/srv/course/proto/course"
	"github.com/micro/cli"
	log "github.com/sirupsen/logrus"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn/srv/course/database"
)

const ServiceName = "go.micro.srv.course"
const ServiceVersion = "1.0.0"

var (
	debugFlag = cli.BoolFlag{
		Name:   "debug",
		EnvVar: "DEBUG",
		Usage:  "Enable debug mode",
	}
	dbHostFlag = cli.StringFlag{
			Name:   "db_host",
			EnvVar: "DB_HOST",
			Usage:  "The database host",
	}
	dbPortFlag = cli.StringFlag{
		Name:   "db_port",
		EnvVar: "DB_PORT",
		Usage:  "The database port",
	}
	dbUserFlag = cli.StringFlag{
		Name:   "db_user",
		EnvVar: "DB_USER",
		Usage:  "The database user",
	}
	dbPassFlag = cli.StringFlag{
		Name:   "db_pass",
		EnvVar: "DB_PASS",
		Usage:  "The database password",
	}
	dbNameFlag = cli.StringFlag{
		Name:   "db_name",
		EnvVar: "DB_NAME",
		Usage:  "The database name",
	}
)

func main() {

	var dbCfg database.DbConfig

	// New Service
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version(ServiceVersion),
		micro.Flags(
			debugFlag,
			dbHostFlag,
			dbPortFlag,
			dbUserFlag,
			dbPassFlag,
			dbNameFlag,
		),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("debug") {
				log.SetLevel(log.DebugLevel)
				log.Debug("DEBUG enabled")
			} else {
				log.SetLevel(log.InfoLevel)
			}
			dbCfg.Host = c.String("db_host")
			dbCfg.Port = c.Int("db_port")
			dbCfg.Username = c.String("db_user")
			dbCfg.Password = c.String("db_pass")
			dbCfg.Database = c.String("db_name")
		}),
	)

	// Setup database
	db, err := database.NewDB(&dbCfg)
	if err != nil {
	    log.Errorf("Unable to connect to storage: %v", err)
	    return
	}

	// Register Handler
	course.RegisterCourseHandler(service.Server(), &handler.Service{DB: db})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("topic.go.micro.srv.course", service.Server(), new(subscriber.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
