package main

type User struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creationDate"`
}

var us []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Create(Id int, Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return us, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Create(Id int, Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error) {
	u := User{Id, Name, Surname, Email, Age, Height, Active, CreationDate}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}
