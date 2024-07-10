package models

import "golang-clean-architecture-example/domain/entities"

type User struct {
	ID   string
	Name string
}

func FromDomainModel(m *entities.User) *User {
	return &User{
		ID:   m.GetID(),
		Name: m.GetName(),
	}
}

func (m *User) ToDomainModel() *entities.User {
	return entities.NewUser(m.ID, m.Name)
}
