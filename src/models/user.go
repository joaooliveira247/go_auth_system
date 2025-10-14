package models

import (
	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:id"`
	Email     string    `gorm:"type:varchar(255);not null;unique;column:email"`
	Password  string    `gorm:"type:varchar(60); not null;column:password"`
	Role      string    `gorm:"type:varchar(30);default:user;column:role"`
	CreatedAt int64     `gorm:"autoCreateTime;type:timestamp;column:created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime;type:timestamp;column:updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
