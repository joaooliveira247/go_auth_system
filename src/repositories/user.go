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
	ChangeUserPassword(id uuid.UUID, password string) error
	Delete(id uuid.UUID) error
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

func (repository *userRepository) GetUserByEmail(
	email string,
) (models.UserModel, error) {
	var user models.UserModel

	if result := repository.db.First(&user, "email = ?", email); result.Error != nil {
		return models.UserModel{}, errors.NewDatabaseError(result.Error)
	}

	return user, nil
}

func (repository *userRepository) ChangeUserPassword(
	id uuid.UUID,
	password string,
) error {

	result := repository.db.Model(&models.UserModel{}).
		Where("id = ?", id).
		Update("password", password)

	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.NewDatabaseError(errors.ErrNothingToUpdate)
	}

	return nil
}

func (repository *userRepository) Delete(id uuid.UUID) error {
	result := repository.db.Delete(models.UserModel{}, id)

	if err := result.Error; err != nil {
		return errors.NewDatabaseError(err)
	}

	if result.RowsAffected == 0 {
		return errors.NewDatabaseError(errors.ErrNothingToDelete)
	}

	return nil
}
