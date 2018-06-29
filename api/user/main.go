package main

import (
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"context"
	api "github.com/micro/micro/api/proto"
	user "github.com/lukasjarosch/educonn/srv/user/proto/user"
)

type User struct {
	Client user.UserServiceClient
}

func (u *User) Create(ctx context.Context, req *api.Request, res *api.Response) {

	log.Error("ASDFASDFASDFASDF")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserServiceClient("go.micro.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
