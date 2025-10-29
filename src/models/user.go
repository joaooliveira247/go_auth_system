package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/joaooliveira247/go_auth_system/src/errors"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:id"`
	Email     string    `gorm:"type:varchar(255);not null;unique;column:email"           json:"email"`
	Password  string    `gorm:"type:varchar(60); not null;column:password"               json:"password"`
	Role      string    `gorm:"type:varchar(30);default:user;column:role"`
	CreatedAt int64     `gorm:"autoCreateTime;type:timestamp;column:created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime;type:timestamp;column:updated_at"`
}

func NewUserModelFromCache(cacheString string) (UserModel, error) {
	var user UserModel

	if err := json.Unmarshal([]byte(cacheString), &user); err != nil {
		return UserModel{}, errors.NewModelError(err)
	}

	return user, nil
}

func (UserModel) TableName() string {
	return "users"
}
