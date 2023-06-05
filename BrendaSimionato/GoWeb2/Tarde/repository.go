package main

import "fmt"

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
	Update(Id int, Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error)
	UpdateSurnameAndAge(Id int, Surname string, age int) (User, error)
	Delete(id int) error
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

func (repository) Update(Id int, Name, Surname, Email string, Age int, Height float64, Active bool, CreationDate string) (User, error) {
	u := User{Name: Name, Surname: Surname, Email: Email, Age: Age, Height: Height, Active: Active, CreationDate: CreationDate}
	updated := false
	for i := range us {
		if us[i].Id == Id {
			u.Id = Id
			us[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuário %d não encontrado", Id)
	}
	return u, nil
}

func (repository) UpdateSurnameAndAge(Id int, Surname string, Age int) (User, error) {
	var u User
	updated := false
	for i := range us {
		if us[i].Id == Id {
			us[i].Surname = Surname
			us[i].Age = Age
			updated = true
			u = us[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuário %d não encontrado", Id)
	}
	return u, nil

}

func (repository) Delete(Id int) error {
	deleted := false
	var index int
	for i := range us {
		if us[i].Id == Id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Usuário %d não encontrado", Id)
	}

	us = append(us[:index], us[index+1:]...)
	return nil
}
