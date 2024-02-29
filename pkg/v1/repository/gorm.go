package repo

import (
	"github.com/EricBui0512/grpc-clean/internal/models"
	interfaces "github.com/EricBui0512/grpc-clean/pkg/v1"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(user models.User) error {
	err := repo.db.Create(&user)
	return err
}

func (repo *Repo) Get(id string) (models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error

	return user, err
}
