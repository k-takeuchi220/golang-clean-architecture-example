package presenters

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IErrorPresenter interface {
	PresentBadRequest(c echo.Context, message string) error
	PresentInternalServerError(c echo.Context, err error) error
}

type ErrorPresenter struct{}

func NewErrorPresenter() IErrorPresenter {
	return &ErrorPresenter{}
}

func (p *ErrorPresenter) PresentBadRequest(c echo.Context, message string) error {
	response := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

	return c.JSON(http.StatusBadRequest, response)
}

func (p *ErrorPresenter) PresentInternalServerError(c echo.Context, err error) error {
	response := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	return c.JSON(http.StatusInternalServerError, response)
}
