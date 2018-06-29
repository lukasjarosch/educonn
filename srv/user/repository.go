package main

import (
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	log "github.com/sirupsen/logrus"
	"github.com/satori/go.uuid"
)

type Respository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) (*pb.User, error)
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db         *sql.DB
	userStruct *sqlbuilder.Struct
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	/*
		var users []*pb.User
		if err := repo.db.Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	*/
	return nil, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	/*
		var user *pb.User
		user.Id = id
		if err := repo.db.First(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
	*/
	return nil, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	/*
		if err := repo.db.First(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
	*/
	return nil, nil
}

func (repo *UserRepository) Create(user *pb.User) (*pb.User, error) {
	uuid, _ := uuid.NewV4()
	u := User{
		Id:        uuid.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
	ib := repo.userStruct.InsertInto("users", u)

	sql, args := ib.Build()

	log.Debugf("Executing SQL: %v", sql)

	_, err := repo.db.Exec(sql, args...)

	user.Id = u.Id

	return user, err
}
