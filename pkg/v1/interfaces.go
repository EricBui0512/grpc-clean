package v1

import "github.com/EricBui0512/grpc-clean/internal/models"

// this is an interface for repo methods
type RepoInterface interface {
	//creates a user with data supplied
	Create(models.User) error

	Get(id string) (models.User, error)

	Update(models.User) error

	Delete(id string) error
}

// this is an interface for usecase methods
type UseCaseInterface interface {
	Create(models.User) error

	Get(id string) (models.User, error)

	Update(models.User) error

	Delete(id string) error
}
