package entities

import "golang-clean-architecture-example/domain/models"

type User struct {
	id   string
	name string
}

func FromDataModel(m *models.User) *User {
	return &User{
		id:   m.ID,
		name: m.Name,
	}
}

func (u *User) ToDataModel() *models.User {
	return &models.User{
		ID:   u.id,
		Name: u.name,
	}
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
