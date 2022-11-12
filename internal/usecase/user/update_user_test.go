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

type UpdateUserUsecaseTestSuite struct {
	suite.Suite
	userRepo *adapter_mock.MockUserRepo
	hasher   *adapter_mock.MockHasher
	ucase    entities.UserUsecase
}

func TestUpdateUserUsecase(t *testing.T) {
	suite.Run(t, new(UpdateUserUsecaseTestSuite))
}

func (s *UpdateUserUsecaseTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = adapter_mock.NewMockUserRepo(ctrl)
	s.hasher = adapter_mock.NewMockHasher(ctrl)
	s.ucase = usecase.NewUserUsecase(
		s.userRepo,
		s.hasher,
	)
}

func (s *UpdateUserUsecaseTestSuite) TestUpdateUser() {
	userId := ksuid.New().String()
	hashedPassword := "hashedPassword"

	testCases := []struct {
		name                 string
		input                *entities.User
		resultMockUpdateUser *entities.User
		errorMockUpdateUser  error
		expectedOutput       *entities.User
		expectedError        error
	}{
		{
			name: "should update user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			resultMockUpdateUser: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: hashedPassword,
			},
			expectedOutput: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: hashedPassword,
			},
		},
		{
			name: "should raise error when failed update user",
			input: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			errorMockUpdateUser: errors.New("failed update user"),
			expectedError:       errors.New("failed update user"),
		},
		{
			name: "should not update user's password if not passed by user",
			input: &entities.User{
				Name:  "user",
				Email: "user@mail.com",
			},
			resultMockUpdateUser: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: hashedPassword,
			},
			expectedOutput: &entities.User{
				ID:       userId,
				Name:     "user",
				Email:    "user@mail.com",
				Password: hashedPassword,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			hashedInput := &entities.User{
				Name:  tc.input.Name,
				Email: tc.input.Email,
			}
			if tc.input.Password != "" {
				s.hasher.EXPECT().HashPassword(tc.input.Password).Return(hashedPassword, nil)
				hashedInput.Password = hashedPassword
			}

			s.userRepo.EXPECT().Update(gomock.Any(), userId, hashedInput).Return(tc.resultMockUpdateUser, tc.errorMockUpdateUser)

			res, err := s.ucase.Update(context.Background(), userId, tc.input)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.EqualValues(tc.expectedOutput, res)
			} else {
				s.Nil(res)
			}
		})
	}
}
