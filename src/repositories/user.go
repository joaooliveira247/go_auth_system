package repositories

import (
	"github.com/google/uuid"
	"github.com/joaooliveira247/go_auth_system/src/errors"
	"github.com/joaooliveira247/go_auth_system/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.UserModel) (uuid.UUID, error)
	GetUserByEmail(email string) (models.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) Create(
	user *models.UserModel,
) (uuid.UUID, error) {

	if result := repository.db.Create(&user); result.Error != nil {
		return uuid.UUID{}, errors.NewDatabaseError(result.Error)
	}

	return user.ID, nil
}
