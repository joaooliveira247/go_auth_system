package repositories

import "gorm.io/gorm"

type UserRepository interface{}

type userRepository struct {
	db *gorm.DB
}
