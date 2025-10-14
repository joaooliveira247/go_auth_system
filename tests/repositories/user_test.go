package repositories_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joaooliveira247/go_auth_system/src/models"
	"github.com/joaooliveira247/go_auth_system/src/repositories"
	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserSuccess(t *testing.T) {
	gormDB, mock := mocks.SetupMockDB()

	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	mockUser := mocks.GenFakeUser()

	repository := repositories.NewUserRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("email","password","role","created_at","updated_at") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(mockUser.Email, mockUser.Password, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUser.ID))

	mock.ExpectCommit()

	id, err := repository.Create(
		&models.UserModel{Email: mockUser.Email, Password: mockUser.Password},
	)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}
