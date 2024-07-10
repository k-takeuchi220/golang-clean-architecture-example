package repositories

import (
	"context"
	"database/sql"
	"golang-clean-architecture-example/domain/entities"
	"golang-clean-architecture-example/domain/repositories"
	"golang-clean-architecture-example/infrastructure/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*entities.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE id = ?", id)
	user := &models.User{}
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return entities.FromDataModel(user), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	model := user.ToDataModel()
	_, err := r.db.ExecContext(ctx, "UPDATE users SET name = ? WHERE id = ?", model.Name, model.ID)
	return err
}
