package infrastructure

import (
	"golang-clean-architecture-example/controllers"
	"golang-clean-architecture-example/infrastructure/repositories"
	"golang-clean-architecture-example/presenters"
	"golang-clean-architecture-example/usecases"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewServer)
	container.Provide(NewDB)

	// controllers
	container.Provide(controllers.NewUserController)

	// presenters
	container.Provide(presenters.NewUserPresenter)
	container.Provide(presenters.NewErrorPresenter)

	// usecases
	container.Provide(usecases.NewUpdateUserNameInteractor)

	// repositories
	container.Provide(repositories.NewUserRepository)

	return container
}
