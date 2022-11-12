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
	input := &entities.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name                 string
		resultMockCreateUser *entities.User
		errorMockCreateUser  error
		expectedOutput       *entities.User
		expectedError        error
	}{
		{
			name: "should create user",
			resultMockCreateUser: &entities.User{
				Id:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: "password",
			},
			expectedOutput: &entities.User{
				Id:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: "password",
			},
		},
		{
			name:                "should raise error when failed persist user",
			errorMockCreateUser: errors.New("failed persist user"),
			expectedError:       errors.New("failed persist user"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.userRepo.EXPECT().Create(gomock.Any(), input).Return(tc.resultMockCreateUser, tc.errorMockCreateUser)

			res, err := s.ucase.CreateUser(context.Background(), input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.EqualValues(tc.expectedOutput, res)
			} else {
				s.Nil(res)
			}
		})
	}
}
