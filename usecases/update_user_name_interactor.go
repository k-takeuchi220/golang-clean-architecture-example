package usecases

import (
	"context"
	"golang-clean-architecture-example/domain/repositories"
	"golang-clean-architecture-example/usecases/dto/input"
	"golang-clean-architecture-example/usecases/dto/output"
)

type IUpdateUserNameInteractor interface {
	Execute(ctx context.Context, r *input.UpdateUserNameInput) (*output.UpdateUserNameOutput, error)
}

type UpdateUserNameInteractor struct {
	userRepository repositories.IUserRepository
}

func NewUpdateUserNameInteractor(
	userRepository repositories.IUserRepository,
) IUpdateUserNameInteractor {
	return &UpdateUserNameInteractor{
		userRepository: userRepository,
	}
}

func (i *UpdateUserNameInteractor) Execute(ctx context.Context, input *input.UpdateUserNameInput) (*output.UpdateUserNameOutput, error) {
	user, err := i.userRepository.GetUser(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	user.SetName(input.NewName)

	err = i.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	output := &output.UpdateUserNameOutput{
		User: user,
	}

	return output, nil
}
