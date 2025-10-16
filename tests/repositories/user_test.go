package repositories_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/joaooliveira247/go_auth_system/src/errors"
	"github.com/joaooliveira247/go_auth_system/src/models"
	"github.com/joaooliveira247/go_auth_system/src/repositories"
	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestCreateUserFail(t *testing.T) {
	gormDB, mock := mocks.SetupMockDB()

	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	testCases := []struct {
		testName          string
		err               error
		errStringExpected string
	}{
		{
			testName:          "UserAlreadyExist",
			err:               gorm.ErrDuplicatedKey,
			errStringExpected: "(Database): duplicated key not allowed",
		},
		{
			testName:          "UnmappedError",
			err:               errors.ErrNotExpected,
			errStringExpected: "(Database): NotExpectedTestError",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockUser := mocks.GenFakeUser()

			repository := repositories.NewUserRepository(gormDB)

			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("email","password","role","created_at","updated_at") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
				WithArgs(mockUser.Email, mockUser.Password, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt).
				WillReturnError(testCase.err)

			mock.ExpectRollback()

			id, err := repository.Create(
				&models.UserModel{
					Email:    mockUser.Email,
					Password: mockUser.Password,
				},
			)

			assert.Error(t, err)
			assert.Equal(t, uuid.Nil, id)
			assert.Equal(t, err.Error(), testCase.errStringExpected)

		})
	}
}

func TestGetUserByEmailSuccess(t *testing.T) {
	gormDB, mock := mocks.SetupMockDB()

	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	mockUser := mocks.GenFakeUser()

	row := mock.NewRows([]string{"id", "email", "password", "role", "created_at", "updated_at"}).
		AddRow(mockUser.ID, mockUser.Email, mockUser.Password, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt)

	repository := repositories.NewUserRepository(gormDB)

	mock.ExpectQuery(
		regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`,
		),
	).WithArgs(mockUser.Email, 1).WillReturnRows(row)

	result, err := repository.GetUserByEmail(mockUser.Email)

	assert.NoError(t, err)

	assert.Equal(t, mockUser, &result)
}
