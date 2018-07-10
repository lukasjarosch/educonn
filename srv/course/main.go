package main

import (
	"context"
	"errors"
	"os"

	"github.com/lukasjarosch/educonn/srv/course/database"
	"github.com/lukasjarosch/educonn/srv/course/handler"
	course "github.com/lukasjarosch/educonn/srv/course/proto/course"
	auth "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

const ServiceName = "educonn.course"
const ServiceVersion = "1.0.0"

var (
	service micro.Service

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
	service = micro.NewService(
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
		micro.WrapHandler(AuthWrapper),
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

	log.Debugf("Database %s", dbCfg.Database)

	// Setup database
	db, err := database.NewDB(&dbCfg)
	if err != nil {
		log.Errorf("Unable to connect to storage: %v", err)
		return
	}

	// Register Handler
	course.RegisterCourseHandler(service.Server(), &handler.Service{DB: db})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("topic.educonn.course", service.Server(), new(subscriber.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, res)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New( "no auth metadata found in request")
		}

		if meta["Token"] == "" {
			return errors.New("no token header found in request")
		}

		token := meta["Token"]
		log.Debugf("authenticating with token %s", token)

		authClient := auth.NewAuthClient("educonn.user", service.Client())
		_, err := authClient.ValidateToken(ctx, &auth.Token{
			Token: token,
		})
		if err != nil {
			return err
		}

		err = fn(ctx, req, res)
		return err
	}
}
