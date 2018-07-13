package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/lukasjarosch/educonn/srv/course/database"
	pb "github.com/lukasjarosch/educonn/srv/course/proto/course"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"strconv"
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

	// Add modules
	err = srv.DB.AddModules(course, req.Modules)
	if err != nil {
		error := &pb.Error{
			Code:    500,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		return err
	}

	log.Debug("created course %s", id.Hex())

	res.Course = database.NewCourseProtoFromDb(course)

	return nil
}

// Get an existing course
func (srv *Service) Get(ctx context.Context, req *pb.CourseEntity, res *pb.CourseResponse) error {

	if req.Id == "" {
		log.Debug("missing id")
		return errors.New("missing id")
	}

	course, err := srv.DB.FindCourseById(req.Id)
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

	res.Course = database.NewCourseProtoFromDb(course)

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
		res.Courses = append(res.Courses, database.NewCourseProtoFromDb(course))
	}

	return nil
}

// Add a module to a given course
func (srv *Service) AddModule(ctx context.Context, req *pb.AddModuleRequest, res *pb.AddModuleResponse) error {
	if req.CourseId == "" {
		error := &pb.Error{
			Code:    400,
			Message: "Missing course_id",
		}
		res.Errors = append(res.Errors, error)
		return errors.New("missing course_id")
	}

	// Fetch course
	course, err := srv.DB.FindCourseById(req.CourseId)
	if err != nil {
		error := &pb.Error{
			Code:    404,
			Message: fmt.Sprintf("Course with id %s not found", req.CourseId),
		}
		res.Errors = append(res.Errors, error)
		return errors.New(fmt.Sprintf("Course with id %s not found", req.CourseId))
	}

	// Append the module
	order, _ := strconv.Atoi(req.Module.Order)
	course.Modules = append(course.Modules, database.Module{
		ID:          bson.NewObjectId(),
		Name:        req.Module.Name,
		Description: req.Module.Description,
		Order:       order,
	})

	err = srv.DB.UpdateCourse(course)
	if err != nil {
		error := &pb.Error{
			Code:    500,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		log.Warn(err)
	}

	return nil
}

// Add a lesson to a given course and module
func (srv *Service) AddLesson(ctx context.Context, req *pb.AddLessonRequest, res *pb.AddLessonResponse) error {
	if req.GetModuleId() == "" {
		error := &pb.Error{
			Code:    400,
			Message: "missing module_id",
		}
		res.Errors = append(res.Errors, error)
		return errors.New("missing module_id")
	}

	// Find the corresponding course including the module
	course, err := srv.DB.FindCourseByModuleId(req.ModuleId)
	if err != nil {
		error := &pb.Error{
			Code:    500,
			Message: err.Error(),
		}
		res.Errors = append(res.Errors, error)
		log.Warnf("error fetching module: %v", err)
		return err
	}

	var lessonId bson.ObjectId

	// Add the lesson to the module with the correct ID
	for k, module := range course.Modules {
		if module.ID.Hex() == req.ModuleId {
			lessonId = bson.NewObjectId()
			order, _ := strconv.Atoi(req.Lesson.Order)
			course.Modules[k].Lessons = append(module.Lessons, database.Lesson{
				ID:      lessonId,
				Name:    req.Lesson.Name,
				VideoID: req.Lesson.VideoId,
				Order:   order,
			})
			break
		}
	}

	// Update the course document
	err = srv.DB.UpdateCourse(course)
	if err != nil {
		log.Error(err)
	}

	res.LessonId = lessonId.Hex()

	return nil
}
