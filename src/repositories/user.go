package repositories

import (
	"github.com/google/uuid"
	"github.com/joaooliveira247/go_auth_system/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.UserModel) (uuid.UUID, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
