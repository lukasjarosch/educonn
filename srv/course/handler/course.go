package handler

import (
	"context"
	"github.com/lukasjarosch/educonn/srv/course/database"
	pb "github.com/lukasjarosch/educonn/srv/course/proto/course"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
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
			Type:        course.Type,
		})
	}

	return nil
}
