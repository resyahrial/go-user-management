package repo_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/internal/repositories/pg/repo/testhelper"
	repo "github.com/resyahrial/go-user-management/internal/repositories/pg/repo/user"
	"github.com/segmentio/ksuid"
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

	s.db, err = gorm.Open(dialector, &testhelper.DisableLog{})
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
				RoleName: "USER",
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
				RoleName: "USER",
			},
			mockErrorPersistUser: errors.New("failed to persist user"),
			expectedError:        errors.New("failed to persist user"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(`
					INSERT INTO "users" ("id","created_at","updated_at","is_deleted","name","email","password","role_name")
					VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
				`).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), false, tc.input.Name, tc.input.Email, tc.input.Password, tc.input.RoleName).
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
				s.NotEmpty(res.ID)
			} else {
				s.Nil(res)
			}
		})
	}
}

func (s *UserRepoTestSuite) TestUpdateUser() {
	userId := ksuid.New().String()

	testCases := []struct {
		name                 string
		input                *entities.User
		mockGetDetailUser    error
		mockErrorPersistUser error
		expectedError        error
		expectedOutput       *entities.User
	}{
		{
			name: "should update user",
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
			name: "should not update user when occur error when update user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockErrorPersistUser: errors.New("failed to update user"),
			expectedError:        errors.New("failed to update user"),
		},
		{
			name: "should not update user when user data not found",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			mockGetDetailUser: gorm.ErrRecordNotFound,
			expectedError:     repo.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			rows := sqlmock.NewRows([]string{"id"})
			if tc.mockGetDetailUser == nil {
				rows = rows.AddRow(userId)
			}
			s.mock.ExpectQuery(`SELECT * FROM "users" WHERE id = $1 AND is_deleted != true ORDER BY "users"."id" LIMIT 1`).
				WithArgs(userId).
				WillReturnRows(rows)

			if tc.mockGetDetailUser == nil {
				s.mock.ExpectBegin()
				s.mock.ExpectQuery(`
						UPDATE "users" 
						SET "updated_at"=$1,"name"=$2,"email"=$3,"password"=$4 
						WHERE id = $5 
						RETURNING *
					`).
					WithArgs(sqlmock.AnyArg(), tc.input.Name, tc.input.Email, tc.input.Password, userId).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userId)).
					WillReturnError(tc.mockErrorPersistUser)

				if tc.mockErrorPersistUser == nil {
					s.mock.ExpectCommit()
				} else {
					s.mock.ExpectRollback()
				}
			}

			res, err := s.repo.Update(context.Background(), userId, tc.input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.Equal(tc.expectedOutput.Name, res.Name)
				s.Equal(tc.expectedOutput.Email, res.Email)
				s.Equal(tc.expectedOutput.Password, res.Password)
				s.NotEmpty(res.ID)
			} else {
				s.Nil(res)
			}
		})
	}
}

func (s *UserRepoTestSuite) TestGetUserById() {
	userId := ksuid.New().String()

	testCases := []struct {
		name                 string
		input                *entities.User
		mockGetDetailUser    error
		mockErrorPersistUser error
		expectedError        error
		expectedOutput       *entities.User
	}{
		{
			name: "should get user detail",
			expectedOutput: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
		{
			name:              "should return error when user data not found",
			mockGetDetailUser: gorm.ErrRecordNotFound,
			expectedError:     repo.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			rows := sqlmock.NewRows([]string{"id", "name", "email", "password"})
			if tc.mockGetDetailUser == nil {
				rows = rows.AddRow(userId, "user", "user@mail.com", "anypassword")
			}
			s.mock.ExpectQuery(`SELECT * FROM "users" WHERE id = $1 AND is_deleted != true ORDER BY "users"."id" LIMIT 1`).
				WithArgs(userId).
				WillReturnRows(rows)

			res, err := s.repo.GetById(context.Background(), userId)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.Equal(tc.expectedOutput.Name, res.Name)
				s.Equal(tc.expectedOutput.Email, res.Email)
				s.Equal(tc.expectedOutput.Password, res.Password)
				s.NotEmpty(res.ID)
			} else {
				s.Nil(res)
			}
		})
	}
}

func (s *UserRepoTestSuite) TestGetUserList() {
	userId := ksuid.New().String()

	testCases := []struct {
		name           string
		input          *entities.PaginatedQueryParams
		expectedError  error
		expectedOutput []*entities.User
		expectedCount  int64
	}{
		{
			name: "should get user list",
			input: &entities.PaginatedQueryParams{
				Page:  0,
				Limit: 10,
			},
			expectedOutput: []*entities.User{
				{
					ID:       userId,
					Name:     "user",
					Email:    "user@mail.com",
					Password: "anypassword",
				},
			},
			expectedCount: 1,
		},
		{
			name: "should get user list, on next page",
			input: &entities.PaginatedQueryParams{
				Page:  1,
				Limit: 10,
			},
			expectedOutput: []*entities.User{
				{
					ID:       userId,
					Name:     "user",
					Email:    "user@mail.com",
					Password: "anypassword",
				},
			},
			expectedCount: 1,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.mock.ExpectQuery(`SELECT count(*) FROM "users" WHERE is_deleted != true`).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(int64(1)))

			query := fmt.Sprintf(`SELECT * FROM "users" WHERE is_deleted != true LIMIT %v`, tc.input.Limit)
			if tc.input.Page > 0 {
				query = fmt.Sprintf(`%s OFFSET %v`, query, tc.input.Limit*tc.input.Page)
			}

			s.mock.ExpectQuery(query).WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "email", "password"}).AddRow(userId, "user", "user@mail.com", "anypassword"),
			)

			res, count, err := s.repo.GetList(context.Background(), tc.input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.EqualValues(tc.expectedOutput, res)
				s.Equal(tc.expectedCount, count)
			} else {
				s.Nil(res)
			}
		})
	}
}

func (s *UserRepoTestSuite) TestDeleteUser() {
	userId := ksuid.New().String()

	testCases := []struct {
		name            string
		isUserFound     bool
		mockDeleteError error
		expectedError   error
	}{
		{
			name:        "should delete user",
			isUserFound: true,
		},
		{
			name:            "should return error when occur error when delete user",
			isUserFound:     true,
			mockDeleteError: errors.New("failed to delete user"),
			expectedError:   errors.New("failed to delete user"),
		},
		{
			name:          "should return error when user data not found",
			expectedError: repo.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			rows := sqlmock.NewRows([]string{"id"})
			if tc.isUserFound {
				rows = rows.AddRow(userId)
			}
			s.mock.ExpectQuery(`SELECT * FROM "users" WHERE id = $1 AND is_deleted != true ORDER BY "users"."id" LIMIT 1`).
				WithArgs(userId).
				WillReturnRows(rows)

			if tc.isUserFound {
				s.mock.ExpectBegin()
				s.mock.ExpectQuery(`UPDATE "users" SET "updated_at"=$1,"is_deleted"=$2 WHERE id = $3 AND is_deleted != true RETURNING *`).
					WithArgs(sqlmock.AnyArg(), true, userId).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userId)).
					WillReturnError(tc.mockDeleteError)

				if tc.mockDeleteError == nil {
					s.mock.ExpectCommit()
				} else {
					s.mock.ExpectRollback()
				}
			}

			err := s.repo.Delete(context.Background(), userId)
			s.Equal(tc.expectedError, err)
		})
	}
}

func (s *UserRepoTestSuite) TestGetUserByEmail() {
	userId := ksuid.New().String()
	email := "user@mail.com"

	testCases := []struct {
		name                 string
		input                *entities.User
		mockGetDetailUser    error
		mockErrorPersistUser error
		expectedError        error
		expectedOutput       *entities.User
	}{
		{
			name: "should get user detail",
			expectedOutput: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    email,
				Password: "anypassword",
			},
		},
		{
			name:              "should return error when user data not found",
			mockGetDetailUser: gorm.ErrRecordNotFound,
			expectedError:     repo.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			rows := sqlmock.NewRows([]string{"id", "name", "email", "password"})
			if tc.mockGetDetailUser == nil {
				rows = rows.AddRow(userId, "user", email, "anypassword")
			}
			s.mock.ExpectQuery(`SELECT * FROM "users" WHERE email = $1 AND is_deleted != true ORDER BY "users"."id" LIMIT 1`).
				WithArgs(email).
				WillReturnRows(rows)

			res, err := s.repo.GetByEmail(context.Background(), email)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.Equal(tc.expectedOutput.Name, res.Name)
				s.Equal(tc.expectedOutput.Email, res.Email)
				s.Equal(tc.expectedOutput.Password, res.Password)
				s.NotEmpty(res.ID)
			} else {
				s.Nil(res)
			}
		})
	}
}

func (s *UserRepoTestSuite) TestGetUserByIdWithPermission() {
	userId := ksuid.New().String()

	testCases := []struct {
		name                 string
		input                *entities.User
		mockGetDetailUser    error
		mockErrorPersistUser error
		expectedError        error
		expectedOutput       *entities.User
	}{
		{
			name: "should get user detail with permissions",
			expectedOutput: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
				Role: &entities.Role{
					Name: "ADMIN",
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "WRITE",
							Type:     "GLOBAL",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role_name"})
			if tc.mockGetDetailUser == nil {
				rows = rows.AddRow(userId, "user", "user@mail.com", "anypassword", "ADMIN")
			}
			s.mock.ExpectQuery(`SELECT * FROM "users" WHERE id = $1 AND is_deleted != true ORDER BY "users"."id" LIMIT 1`).
				WithArgs(userId).
				WillReturnRows(rows)

			s.mock.ExpectQuery(`SELECT * FROM "roles" WHERE "roles"."name" = $1`).
				WithArgs("ADMIN").
				WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("ADMIN"))

			s.mock.ExpectQuery(`SELECT * FROM "roles_permissions" WHERE "roles_permissions"."role_name" = $1`).
				WithArgs("ADMIN").
				WillReturnRows(sqlmock.NewRows([]string{"role_name", "permission_id"}).AddRow("ADMIN", "admin_users_write_global"))

			s.mock.ExpectQuery(`SELECT * FROM "permissions" WHERE "permissions"."id" = $1`).
				WithArgs("admin_users_write_global").
				WillReturnRows(sqlmock.NewRows([]string{"id", "resource", "action", "type"}).AddRow("admin_users_write_global", "users", "WRITE", "GLOBAL"))

			res, err := s.repo.GetByIdWithPermission(context.Background(), userId)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.Equal(tc.expectedOutput.Name, res.Name)
				s.Equal(tc.expectedOutput.Email, res.Email)
				s.Equal(tc.expectedOutput.Password, res.Password)
				s.Equal(tc.expectedOutput.Role.Name, res.Role.Name)
				s.Len(res.Role.Permissions, len(tc.expectedOutput.Role.Permissions))
				s.NotEmpty(res.ID)
				if len(tc.expectedOutput.Role.Permissions) > 0 {
					expectedPermission := tc.expectedOutput.Role.Permissions[0]
					s.Equal(expectedPermission.Action, res.Role.Permissions[0].Action)
					s.Equal(expectedPermission.Type, res.Role.Permissions[0].Type)
					s.Equal(expectedPermission.Type, res.Role.Permissions[0].Type)
				}
			} else {
				s.Nil(res)
			}
		})
	}
}
