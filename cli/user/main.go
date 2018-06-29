package main

import (
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	microclient "github.com/micro/go-micro/client"
	"context"
	"github.com/prometheus/common/log"
	"os"
)

func main() {
	client := pb.NewAuthClient("go.micro.srv.user", microclient.DefaultClient)

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

	authResp, err := client.Auth(context.TODO(), &pb.User{
		Email: email,
		Password: password,
	})
	if err != nil {
	    log.Warnf("Unable to authenticate user: %v", err)
	}

	log.Infof("Access token: %s", authResp.Token)

	/*
	r, err := client.GetAll(context.TODO(), &pb.Request{})
	if err != nil {
	    log.Fatal(err)
	}

	log.Warn(r)
	r, err = client.Get(context.TODO(), &pb.User{Id: r.User.Id})
	if err != nil {
		log.Fatal(err)
	}

	log.Debug(r)
	*/
}
