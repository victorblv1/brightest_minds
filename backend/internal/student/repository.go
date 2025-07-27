package student

import "errors"

var ErrNotFound = errors.New("student not found")

// Repository defines contract for student data storage
type Repository interface {
	GetAll() ([]Student, error)
	GetByID(id string) (Student, error)
}

// InMemoryRepository is a mock implementation for development
type InMemoryRepository struct {
	students []Student
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		students: []Student{
			{
				ID: "1", Name: "Alice", Surname: "Johnson", Gender: GenderFemale,
				Nationality: "USA", Ethnicity: "Caucasian", College: "MIT",
				AcademicStatus: "PhD", Area: "Physics", StudyMajor: "Quantum Mechanics",
				Project: "Revolutionizing Quantum Computing",
			},
			{
				ID: "2", Name: "John", Surname: "Doe", Gender: GenderMale,
				Nationality: "UK", Ethnicity: "Asian", College: "Oxford",
				AcademicStatus: "Bachelor", Area: "Politics", StudyMajor: "NGOs",
				Project: "Global Charity Network",
			},
		},
	}
}

func (r *InMemoryRepository) GetAll() ([]Student, error) {
	return r.students, nil
}

func (r *InMemoryRepository) GetByID(id string) (Student, error) {
	for _, s := range r.students {
		if s.ID == id {
			return s, nil
		}
	}
	return Student{}, ErrNotFound
}
