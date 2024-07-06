package presenters

import (
	"golang-clean-architecture-example/api"
	"golang-clean-architecture-example/usecases/dto/output"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserPresenter interface {
	PresentUpdateUserName(c echo.Context, output *output.UpdateUserNameOutput) error
}

type UserPresenter struct{}

func NewUserPresenter() IUserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) PresentUpdateUserName(c echo.Context, output *output.UpdateUserNameOutput) error {
	response := api.UpdateUserNameResponse{
		ID:   output.User.GetID(),
		Name: output.User.GetName(),
	}

	return c.JSON(http.StatusOK, response)
}
