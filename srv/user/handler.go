package main

import (
	"context"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
)

type service struct {
	repo         Respository
	tokenService *TokenService
	pubCreated   micro.Publisher
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

// GetAll users
func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.UserResponse) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testing123"
	return nil
}

// Create a new user and publish an UserCreatedEvent. If the user could not be created, the error will be stuffed in the
// UserResponse
func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	user, err := srv.repo.Create(req)

	if err != nil {
		log.Warn(err.Error())

		res.Errors = []*pb.Error{{
			Code:        500,
			Description: err.Error(),
		}}
	}

	res.User = user
	log.Debugf("Created user '%s'", user.Id)
	PublishUserCreated(srv.pubCreated, *user)

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
