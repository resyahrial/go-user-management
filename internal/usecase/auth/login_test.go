package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-user-management/internal/entities"
	usecase "github.com/resyahrial/go-user-management/internal/usecase/auth"
	adapter_mock "github.com/resyahrial/go-user-management/internal/usecase/auth/mocks"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type LoginUsecaseTestSuite struct {
	suite.Suite
	userRepo     *adapter_mock.MockUserRepo
	hasher       *adapter_mock.MockHasher
	tokenHandler *adapter_mock.MockTokenHandler
	ucase        entities.AuthUsecase
}

func TestLoginUsecase(t *testing.T) {
	suite.Run(t, new(LoginUsecaseTestSuite))
}

func (s *LoginUsecaseTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = adapter_mock.NewMockUserRepo(ctrl)
	s.hasher = adapter_mock.NewMockHasher(ctrl)
	s.tokenHandler = adapter_mock.NewMockTokenHandler(ctrl)
	s.ucase = usecase.NewAuthUsecase(
		s.userRepo,
		s.hasher,
		s.tokenHandler,
	)
}

func (s *LoginUsecaseTestSuite) TestLogin() {
	userId := ksuid.New().String()
	hashedPassword := "hashedPassword"
	input := &entities.Login{
		Email:    "user@mail.com",
		Password: "nonHashedPassword",
	}

	testCases := []struct {
		name                string
		mockErrGetUser      error
		isPasswordMatch     bool
		mockAccessToken     string
		mockErrAccessToken  error
		expectedAccessToken string
		expectedError       error
	}{
		{
			name:                "should success login",
			isPasswordMatch:     true,
			mockAccessToken:     "token",
			expectedAccessToken: "token",
		},
		{
			name:               "should return error when failed sign token",
			isPasswordMatch:    true,
			mockErrAccessToken: errors.New("failed sign token"),
			expectedError:      errors.New("failed sign token"),
		},
		{
			name:          "should return error when password not match",
			expectedError: usecase.ErrInvalidLoginIput,
		},
		{
			name:           "should return error when user not found",
			mockErrGetUser: errors.New("user not found"),
			expectedError:  errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.userRepo.EXPECT().GetByEmail(gomock.Any(), input.Email).Return(&entities.User{
				ID:       userId,
				Email:    input.Email,
				Password: hashedPassword,
			}, tc.mockErrGetUser)

			if tc.mockErrGetUser == nil {
				s.hasher.EXPECT().CheckPasswordHash(input.Password, hashedPassword).Return(tc.isPasswordMatch)

				if tc.isPasswordMatch {
					s.tokenHandler.EXPECT().SignToken(map[string]interface{}{
						"id": userId,
					}).Return(tc.mockAccessToken, tc.mockErrAccessToken)
				}
			}

			token, err := s.ucase.Login(context.Background(), input)
			s.Equal(tc.expectedError, err)
			if tc.expectedError == nil {
				s.Equal(tc.expectedAccessToken, token.Access)
			} else {
				s.Nil(token)
			}
		})
	}
}
