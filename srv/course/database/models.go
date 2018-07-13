package database

import "gopkg.in/mgo.v2/bson"
import pb "github.com/lukasjarosch/educonn/srv/course/proto/course"

type Course struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Type        string        `bson:"type" json:"type"`
	Topics      []string      `bson:"topics" json:"topics"`
	Modules     []Module      `bson:"modules" json:"modules"`
}

type Module struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"title" json:"title"`
	Order       int           `bson:"order" json:"order"`
	Description string        `bson:"description" json:"description"`
	Lessons     []Lesson      `bson:"lessons" json:"lessons"`
}

type Lesson struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"title" json:"title"`
	Order   int           `bson:"order" json:"order"`
	VideoID string        `bson:"video_id" json:"video_id"`
}

func NewCourseProtoFromDb(course Course) *pb.CourseEntity {

	var pbModules []*pb.Module
	for _, m := range course.Modules {

		var pbLessons []*pb.Lesson
		for _, l := range m.Lessons {
			lesson := &pb.Lesson{
				Id:      l.ID.Hex(),
				Order:   string(l.Order),
				Name:    l.Name,
				VideoId: l.VideoID,
			}
			pbLessons = append(pbLessons, lesson)
		}

		mod := &pb.Module{
			Id:          m.ID.Hex(),
			Name:        m.Name,
			Order:       string(m.Order),
			Description: m.Description,
			Lessons:     pbLessons,
		}
		pbModules = append(pbModules, mod)
	}

	pbCourse := &pb.CourseEntity{
		Id:          course.ID.Hex(),
		Title:       course.Title,
		Description: course.Description,
		Topics:      course.Topics,
		Type:        course.Type,
		Modules:     pbModules,
	}

	return pbCourse
}
