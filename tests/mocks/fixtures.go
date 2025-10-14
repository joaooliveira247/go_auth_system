package mocks

import (
	"log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/joaooliveira247/go_auth_system/src/models"
	"github.com/joaooliveira247/go_auth_system/src/security"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{Conn: db}),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return gormDB, mock
}

func GenFakePassword() string {
	return faker.Password()
}

func GenFakeUser() *models.UserModel {
	hashedPassword, _ := security.GenHash(faker.Password())
	return &models.UserModel{
		ID:        uuid.New(),
		Email:     faker.Email(),
		Password:  hashedPassword,
		Role:      "user",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
