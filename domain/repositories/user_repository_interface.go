package repositories

import (
	"context"
	"golang-clean-architecture-example/domain/entities"
)

type IUserRepository interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
}
