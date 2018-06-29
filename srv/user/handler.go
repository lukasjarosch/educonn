package main

import (
	"context"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"strings"
)

type service struct {
	repo         Respository
	tokenService *TokenService
	pubCreated   micro.Publisher
}

var (
	errorDuplicateEmail = "Email is already in use!"
)

// Get a specific user
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
	user, err := srv.repo.GetByEmail(req)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
	    return err
	}

	res.Token = token
	return nil
}

// Create a new user and publish an UserCreatedEvent. If the user could not be created, the error will be stuffed in the
// UserResponse
func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.UserResponse) error {

	if req.Email == "" {
		return errors.New("No email provided")
	}
	if req.Password == "" {
		return errors.New("No password provided")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
	    return err
	}
	req.Password = string(hashedPassword)
	user, err := srv.repo.Create(req)

	if err != nil {
		log.Info(err.Error())

		if strings.Contains(strings.ToLower(err.Error()), "email") &&
			strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				res.Errors = []*pb.Error{{
					Code: 500,
					Description: errorDuplicateEmail,
				}}

		} else {
			res.Errors = []*pb.Error{{
				Code:        500,
				Description: err.Error(),
			}}
		}
	}

	res.User = user
	log.Debugf("Created user '%s'", user.Id)
	PublishUserCreated(srv.pubCreated, *user)

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
	    return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true
	return nil
}
