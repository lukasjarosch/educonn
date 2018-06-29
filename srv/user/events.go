package main

import (
	log "github.com/sirupsen/logrus"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"github.com/micro/go-micro"
	"context"
)

var (
	userCreatedTopic = "user.events.created"
)

func PublishUserCreated(publisher micro.Publisher, user pb.User) {
	event := &pb.UserCreatedEvent{
		Id: user.Id,
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}
	/*
	body, err := proto.Marshal(&event)
	if err != nil {
		log.Warnf("Could not marshal user message: %v", err)
	}

	msg := &broker.Message{
		Header: map[string]string{
			"user_id": user.Id,
		},
		Body: body,
	}
	*/

	if err := publisher.Publish(context.Background(), event); err != nil {
		log.Warnf("Unable to publish message: %v", err)
	}

	log.Debugf("UserCreatedEvent for user '%s' published to '%s'", user.Id, userCreatedTopic)
}