package output

import "golang-clean-architecture-example/domain/entities"

type UpdateUserNameOutput struct {
	User *entities.User
}
