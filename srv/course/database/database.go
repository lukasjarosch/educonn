package database

import "gopkg.in/mgo.v2"
import (
	"fmt"
	log "github.com/sirupsen/logrus"
	course "github.com/lukasjarosch/educonn/srv/course/proto/course"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"errors"
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

// NewDB creates a new database root session
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

// Inserts adds a new course
func (s *DB) Insert(course Course) error {
	sess := s.DB.Clone()
	defer sess.Close()

	err := sess.DB(s.DbName).C(COLLECTION).Insert(course)
	if err != nil {
	    return err
	}

	return nil
}

// FindCourseById searches for a course with a given id
func (s *DB) FindCourseById(id string) (Course, error) {
	sess := s.DB.Clone()
	defer sess.Close()

	course := Course{}

	err := sess.DB(s.DbName).C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&course)
	if err != nil {
	    return Course{}, err
	}

	return course, nil
}

// FindCourseByModuleId returns a course with the desired module
func (s *DB) FindCourseByModuleId(id string) (Course, error) {
	sess := s.DB.Clone()
	defer sess.Close()

	log.Debugf("Searching for module: %s", id)
	course := Course{}

	err := sess.DB(s.DbName).C(COLLECTION).Find(
		bson.M{
			"modules._id": bson.ObjectIdHex(id),
		}).One(&course)
	if err != nil {
	    return Course{}, err
	}

	for _, module := range course.Modules {
		if module.ID.Hex() == id {
			return course, nil
		}
	}

	return Course{}, errors.New("module not found inside course")
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

func (s *DB) UpdateCourse(course Course) error {
	sess := s.DB.Clone()
	defer sess.Close()

	err := sess.DB(s.DbName).C(COLLECTION).UpdateId(course.ID, course)
	if err != nil {
	    return err
	}
	log.Debugf("updated course %s", course.ID.Hex())

	return nil
}

// AddModules adds a list of modules to a course. The module list will be replaced, not added.
func (s *DB) AddModules(course Course, modules []*course.Module) error {
	sess := s.DB.Clone()
	defer sess.Close()

	for _, module := range modules {
		order,  _ := strconv.Atoi(module.Order)

		// Create module
		m := Module{
			ID: bson.NewObjectId(),
			Name: module.Name,
			Description: module.Description,
			Order: order,
		}

		// Add module lessons
		err := s.addLessons(&m, module.Lessons)
		if err != nil {
			log.Error(err)
			return err
		}

		course.Modules = append(course.Modules, m)
	}

	err := sess.DB(s.DbName).C(COLLECTION).UpdateId(course.ID, course)

	if err != nil {
		return err
	}
	return nil
}

func (s *DB) addLessons(module *Module, lessons []*course.Lesson) error {
	sess := s.DB.Clone()
	defer sess.Close()

	for _, lesson := range lessons {
		order,  _ := strconv.Atoi(lesson.Order)
		l := Lesson{
			ID: bson.NewObjectId(),
			Name: lesson.Name,
			Order: order,
			VideoID: lesson.VideoId,
		}
		module.Lessons = append(module.Lessons, l)
	}
	return nil
}

