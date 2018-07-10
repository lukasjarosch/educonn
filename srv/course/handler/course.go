package handler

import (
	"context"
	"github.com/lukasjarosch/educonn/srv/course/database"
	pb "github.com/lukasjarosch/educonn/srv/course/proto/course"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Service struct {
	DB *database.DB
}

// Create a new course
func (srv *Service) Create(ctx context.Context, req *pb.CourseEntity, res *pb.CourseResponse) error {

	id := bson.NewObjectId()

	course := database.Course{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		Topics:      req.Topics,
	}

	err := srv.DB.Insert(course)
	if err != nil {
		error := &pb.Error{
			Code:    500,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		return err
	}

	log.Infof("created course %s", id.Hex())

	return nil
}

// Get an existing course
func (srv *Service) Get(ctx context.Context, req *pb.CourseEntity, res *pb.CourseResponse) error {

	if req.Id == "" {
		log.Debug("missing id")
		return errors.New("missing id")
	}

	course, err := srv.DB.FindById(req.Id)
	if err != nil {
		error := &pb.Error{
			Code:    404,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		log.Debugf("%s: %s", req.Id, err.Error())
		return err
	}

	log.Debugf("%s: %s", req.Id, course.Title)

	if course.ID == "" {
		error := &pb.Error{
			Code:    404,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		return err
	}

	res.Course = &pb.CourseEntity{
		Id:          course.ID.Hex(),
		Title:       course.Title,
		Description: course.Description,
		Topics:      course.Topics,
		Type:        course.Type,
	}

	return nil
}

// GetAll courses
func (srv *Service) GetAll(ctx context.Context, req *pb.Request, res *pb.CourseResponse) error {
	var courses []database.Course

	courses, err := srv.DB.GetAll()
	if err != nil {
		error := &pb.Error{
			Code:    500,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		return err
	}

	for _, course := range courses {
		res.Courses = append(res.Courses, &pb.CourseEntity{
			Id:          course.ID.Hex(),
			Title:       course.Title,
			Description: course.Description,
			Topics:      course.Topics,
			Type:        course.Type,
		})
	}

	return nil
}
