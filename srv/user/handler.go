package main

import (
	"context"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
)

type service struct {
	repo Respository
	tokenService *TokenService
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
	    return err
	}
	res.User = user
	return nil
}

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

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.UserResponse) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}