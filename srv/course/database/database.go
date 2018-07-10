package database

import "gopkg.in/mgo.v2"
import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "courses"

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type DB struct {
	DB *mgo.Session
	DbName string
}

func NewDB(config *DbConfig) (*DB, error) {
	connString := fmt.Sprintf("%s:%d/%s",
		config.Host,
		config.Port,
		config.Database,
	)
	log.Debugf("Connected to: %s", connString)
	session, err := mgo.Dial(connString)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: session,
		DbName: config.Database,
	}, nil
}

func (s *DB) Insert(course Course) error {
	sess := s.DB.Clone()
	defer sess.Close()

	err := sess.DB(s.DbName).C(COLLECTION).Insert(course)
	if err != nil {
	    return err
	}

	return nil
}

func (s *DB) FindById(id string) (Course, error) {
	sess := s.DB.Clone()
	defer sess.Close()

	course := Course{}

	err := sess.DB(s.DbName).C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&course)
	log.Debug(course)
	if err != nil {
	    return Course{}, err
	}

	return course, nil
}

func (s *DB) GetAll() ([]Course, error) {
	sess := s.DB.Clone()
	defer sess.Close()

	var courses []Course

	err := sess.DB(s.DbName).C(COLLECTION).Find(bson.M{}).All(&courses)
	if err != nil {
		return courses, err
	}
	return courses, nil
}

func (s *DB) Delete(id string) error {
	sess := s.DB.Clone()
	defer sess.Close()

	err := sess.DB(s.DbName).C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
	    return err
	}
	return nil
}

func (s *DB) Update(course Course) error {
	sess := s.DB.Clone()
	defer sess.Close()

	err := sess.DB(s.DbName).C(COLLECTION).UpdateId(course.ID, course)
	if err != nil {
	    return err
	}

	return nil
}
