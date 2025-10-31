package mocks

import (
	"log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/go-redis/redismock/v9"
	"github.com/google/uuid"
	"github.com/joaooliveira247/go_auth_system/src/models"
	"github.com/joaooliveira247/go_auth_system/src/schemas"
	"github.com/joaooliveira247/go_auth_system/src/security"
	"github.com/redis/go-redis/v9"
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

func SetupMockCache() (*redis.Client, redismock.ClientMock) {
	return redismock.NewClientMock()
}

func GenFakePassword() string {
	return faker.Password()
}

func GenFakeUserModel() *models.UserModel {
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

func GenFakeUserSchemaIn() *schemas.UserSchemaIn {
	password := GenFakePassword()
	return &schemas.UserSchemaIn{
		Email:           faker.Email(),
		Password:        password,
		ConfirmPassword: password,
	}
}

func GenHashedPassword() string {
	hashedPassword, _ := security.GenHash(GenFakePassword())
	return hashedPassword
}
