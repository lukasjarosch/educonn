package database

type DataStore interface {
	Insert(course Course) error
	FindById(id string) (Course, error)
	GetAll() ([]Course, error)
	Update(course Course) error
	Delete(id string) error
}

