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
	var users []*pb.User

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "email", "first_name", "last_name")
	sb.From("users")

	sql, args := sb.Build()
	rows, err := repo.db.Query(sql, args...)
	defer rows.Close()
	if err != nil {
	    return nil, err
	}

	for rows.Next() {
		user := &pb.User{}
		users = append(users, user)
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
		    log.Warnf("Error iterating result rows: %v", err)
		    return nil, err
		}
	}

	log.Debugf("Fetched %d users from database", len(users))

	return users, nil
}

func (repo *UserRepository) Get(id string) (*pb.User, error) {
	user := &pb.User{}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "email", "first_name", "last_name")
	sb.From("users")
	sb.Where(sb.Equal("id", id))
	sb.Limit(1)

	sql, args := sb.Build()
	row, err := repo.db.Query(sql, args...)
	defer row.Close()
	if err != nil {
	    log.Warnf("Error querying database: %v", err)
	    return nil, err
	}

	row.Next()
	err = row.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
	    log.Warnf("Error scanning result: %v", err)
	}
	log.Debugf("Fetched user '%s' from database", user.Id)

	return user, nil
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
