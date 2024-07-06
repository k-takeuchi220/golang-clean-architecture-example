package controllers

import (
	"golang-clean-architecture-example/api"
	"golang-clean-architecture-example/presenters"
	"golang-clean-architecture-example/usecases"
	"golang-clean-architecture-example/usecases/dto/input"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	updateUserNameInteractor usecases.IUpdateUserNameInteractor
	userPresenter            presenters.IUserPresenter
	errorPresenter           presenters.IErrorPresenter
}

func NewUserController(
	updateUserNameInteractor usecases.IUpdateUserNameInteractor,
	userPresenter presenters.IUserPresenter,
	errorPresenter presenters.IErrorPresenter,
) *UserController {
	return &UserController{
		updateUserNameInteractor: updateUserNameInteractor,
		userPresenter:            userPresenter,
		errorPresenter:           errorPresenter,
	}
}

func (c *UserController) UpdateUserName(e echo.Context, userID string) error {
	req := api.UpdateUserNameRequest{}
	if err := e.Bind(&req); err != nil {
		return c.errorPresenter.PresentBadRequest(e, "invalid request")
	}

	input := &input.UpdateUserNameInput{
		UserID:  userID,
		NewName: req.Name,
	}
	ctx := e.Request().Context()

	output, err := c.updateUserNameInteractor.Execute(ctx, input)
	if err != nil {
		return c.errorPresenter.PresentInternalServerError(e, err)
	}

	return c.userPresenter.PresentUpdateUserName(e, output)
}
