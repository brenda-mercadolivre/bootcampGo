package main

type Service interface {
	GetAll() ([]User, error)
	Create(Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
func (s *service) GetAll() ([]User, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Create(Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	user, err := s.repository.Create(lastID, Name, Surname, Email, Age, Height, Active, CreationDate)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
