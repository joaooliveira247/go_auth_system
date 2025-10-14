package mocks

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
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
