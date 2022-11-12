package repo_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/resyahrial/go-template/internal/entities"
	repo "github.com/resyahrial/go-template/internal/repositories/pg/repo/user"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *repo.UserRepoImpl
}

func TestUserRepo(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (s *UserRepoTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	s.Nil(err)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	s.db, err = gorm.Open(dialector)
	s.Nil(err)

	s.repo = repo.NewUserRepo(s.db)
}

func (s *UserRepoTestSuite) TestCreateUser() {
	testCases := []struct {
		name                 string
		input                *entities.User
		mockErrorPersistUser error
		expectedError        error
		expectedOutput       *entities.User
	}{
		{
			name: "should create user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			expectedOutput: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
		{
			name: "should not create user when occur error when persist user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockErrorPersistUser: errors.New("failed to persist user"),
			expectedError:        errors.New("failed to persist user"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(`
					INSERT INTO "users" ("id","created_at","updated_at","is_deleted","name","email","password")
					VALUES ($1,$2,$3,$4,$5,$6,$7)
				`).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), false, tc.input.Name, tc.input.Email, tc.input.Password).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tc.mockErrorPersistUser)

			if tc.mockErrorPersistUser == nil {
				s.mock.ExpectCommit()
			} else {
				s.mock.ExpectRollback()
			}

			res, err := s.repo.Create(context.Background(), tc.input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.Equal(tc.expectedOutput.Name, res.Name)
				s.Equal(tc.expectedOutput.Email, res.Email)
				s.Equal(tc.expectedOutput.Password, res.Password)
				s.NotEmpty(res.Id)
			} else {
				s.Nil(res)
			}
		})
	}
}
