package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn/srv/video/handler"
	"github.com/lukasjarosch/educonn/srv/video/subscriber"

	example "github.com/lukasjarosch/educonn/srv/video/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.video"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.video", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.video", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
