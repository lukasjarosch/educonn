package main

import (
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	pb "github.com/lukasjarosch/educonn/srv/user/proto/user"
	log "github.com/sirupsen/logrus"
	"github.com/satori/go.uuid"
	"time"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Respository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) (*pb.User, error)
	GetByEmail(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db         *sql.DB
	userStruct *sqlbuilder.Struct
}

// GetAll users where deleted_at is NULL
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "email", "first_name", "last_name", "password")
	sb.From("users")
	sb.Where(sb.IsNull("deleted_at"))

	sql, args := sb.Build()
	rows, err := repo.db.Query(sql, args...)
	defer rows.Close()
	if err != nil {
	    return nil, err
	}

	for rows.Next() {
		user := &pb.User{}
		users = append(users, user)
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
		    log.Warnf("Error iterating result rows: %v", err)
		    return nil, err
		}
	}

	log.Debugf("Fetched %d users from database", len(users))

	return users, nil
}

// Get a user by ID
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	user := &pb.User{}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "email", "first_name", "last_name", "password")
	sb.From("users")
	sb.Where(
		sb.Equal("id", id),
		sb.IsNull("deleted_at"))
	sb.Limit(1)

	sql, args := sb.Build()
	row, err := repo.db.Query(sql, args...)
	if err != nil {
	    log.Warnf("Error querying database: %v", err)
	    return nil, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
	    log.Warnf("Error scanning result: %v", err)
	}
	log.Debugf("Fetched user '%s' from database", user.Id)

	return user, nil
}

func (repo *UserRepository) GetByEmail(user *pb.User) (*pb.User, error) {
	fetchedUser := &pb.User{}


	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "email", "first_name", "last_name", "password")
	sb.From("users")
	sb.Where(
		sb.Equal("email", user.Email),
		sb.IsNull("deleted_at"))
	sb.Limit(1)


	sql, args := sb.Build()
	log.Debugf("SQL: %v", sql)
	row, err := repo.db.Query(sql, args...)
	if err != nil {
	    log.Warnf("Error querying database: %v", err)
	    return nil, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&fetchedUser.Id, &fetchedUser.Email, &fetchedUser.FirstName, &fetchedUser.LastName, &fetchedUser.Password)
	if err != nil {
		log.Warnf("Error scanning result: %v", err)
	}

	log.Debugf("Fetched user '%s' by mail '%s' from database", fetchedUser.Id, fetchedUser.Email)

	return fetchedUser, nil
}

func (repo *UserRepository) Create(user *pb.User) (*pb.User, error) {
	uuid := uuid.NewV4()
	u := User{
		Id:        uuid.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: mysql.NullTime{
			Valid: false,
			Time: time.Now(),
		},
	}
	ib := repo.userStruct.InsertInto("users", u)

	sql, args := ib.Build()

	log.Debugf("Executing SQL: %v", sql)

	_, err := repo.db.Exec(sql, args...)

	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		log.Error(err)
	}

	if mysqlErr.Number == MYSQL_KEY_EXISTS {
		err = errors.New("the email address is already taken by another user")
	}

	user.Id = u.Id

	return user, err
}
