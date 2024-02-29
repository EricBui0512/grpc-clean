package usecase

import (
	"errors"

	"github.com/EricBui0512/grpc-clean/internal/models"
	interfaces "github.com/EricBui0512/grpc-clean/pkg/v1"
	"gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &UseCase{repo}
}

func (uc *UseCase) Create(user models.User) (models.User, error) {
	if _, err := uc.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("the email is already associated with another user")
	}

	return uc.repo.Create(user)
}

func (uc *UseCase) Get(id string) (models.User, error) {
	var user models.User
	var err error

	if user, err = uc.repo.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("no such user with the id supplied")
		}

		return models.User{}, err
	}
	return user, nil
}
