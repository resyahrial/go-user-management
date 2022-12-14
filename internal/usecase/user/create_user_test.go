package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-user-management/internal/entities"
	usecase "github.com/resyahrial/go-user-management/internal/usecase/user"
	adapter_mock "github.com/resyahrial/go-user-management/internal/usecase/user/mocks"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type CreateUserUsecaseTestSuite struct {
	suite.Suite
	userRepo *adapter_mock.MockUserRepo
	hasher   *adapter_mock.MockHasher
	ucase    entities.UserUsecase
}

func TestCreateUserUsecase(t *testing.T) {
	suite.Run(t, new(CreateUserUsecaseTestSuite))
}

func (s *CreateUserUsecaseTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = adapter_mock.NewMockUserRepo(ctrl)
	s.hasher = adapter_mock.NewMockHasher(ctrl)
	s.ucase = usecase.NewUserUsecase(
		s.userRepo,
		s.hasher,
	)
}

func (s *CreateUserUsecaseTestSuite) TestCreateUser() {
	userId := ksuid.New().String()
	hashedPassword := "hashedPassword"

	testCases := []struct {
		name                 string
		input                *entities.User
		resultMockCreateUser *entities.User
		errorMockCreateUser  error
		expectedOutput       *entities.User
		expectedError        error
	}{
		{
			name: "should create user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
				RoleName: "USER",
			},
			resultMockCreateUser: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: hashedPassword,
				RoleName: "USER",
			},
			expectedOutput: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: hashedPassword,
				RoleName: "USER",
			},
		},
		{
			name: "should raise error when failed persist user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
				RoleName: "USER",
			},
			errorMockCreateUser: errors.New("failed persist user"),
			expectedError:       errors.New("failed persist user"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.hasher.EXPECT().HashPassword(tc.input.Password).Return(hashedPassword, nil)

			hashedInput := &entities.User{
				Name:     tc.input.Name,
				Email:    tc.input.Email,
				Password: hashedPassword,
				RoleName: tc.input.RoleName,
			}

			s.userRepo.EXPECT().Create(gomock.Any(), hashedInput).Return(tc.resultMockCreateUser, tc.errorMockCreateUser)

			res, err := s.ucase.Create(context.Background(), tc.input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.EqualValues(tc.expectedOutput, res)
			} else {
				s.Nil(res)
			}
		})
	}
}
