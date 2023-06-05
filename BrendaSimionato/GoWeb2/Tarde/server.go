package main

type Service interface {
	GetAll() ([]User, error)
	Create(Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error)
	Update(Id int, Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error)
	UpdateSurnameAndAge(Id int, Surname string, age int) (User, error)
	Delete(id int) error
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
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
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

func (s service) Update(Id int, Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error) {
	user, err := s.repository.Update(Id, Name, Surname, Email, Age, Height, Active, CreationDate)

	return user, err
}

func (s service) UpdateSurnameAndAge(Id int, Surname string, Age int) (User, error) {
	user, err := s.repository.UpdateSurnameAndAge(Id, Surname, Age)

	return user, err

}

func (s service) Delete(Id int) error {
	err := s.repository.Delete(Id)

	return err
}
