package student

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllStudents() ([]Student, error) {
	return s.repo.GetAll()
}

func (s *Service) GetStudentByID(id string) (Student, error) {
	return s.repo.GetByID(id)
}
