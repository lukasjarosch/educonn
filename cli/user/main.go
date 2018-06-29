package main

import (
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	microclient "github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"context"
	"os"
)

func main() {
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	firstName := "Lukas"
	lastName := "Jarosch"
	email := "lukas.jarosch@mail.com"
	password := "asdf"

	r, err := client.Create(context.TODO(), &pb.User{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
	})

	if  r.GetErrors() != nil {
		for _, err := range r.Errors {
			log.Error(err)
			return
		}
	}

	if err != nil {
	    log.Fatal(err)
	    os.Exit(1)
	}

	log.Infof("Created user: %s (%s)", r.User.Id, r.User.Email)
}
