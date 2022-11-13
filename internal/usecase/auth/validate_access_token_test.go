package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-user-management/internal/entities"
	usecase "github.com/resyahrial/go-user-management/internal/usecase/auth"
	adapter_mock "github.com/resyahrial/go-user-management/internal/usecase/auth/mocks"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type ValidateAccessTokenUsecaseTestSuite struct {
	suite.Suite
	userRepo     *adapter_mock.MockUserRepo
	hasher       *adapter_mock.MockHasher
	tokenHandler *adapter_mock.MockTokenHandler
	ucase        entities.AuthUsecase
}

func TestValidateAccessUsecase(t *testing.T) {
	suite.Run(t, new(ValidateAccessTokenUsecaseTestSuite))
}

func (s *ValidateAccessTokenUsecaseTestSuite) SetupTest() {
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

func (s *ValidateAccessTokenUsecaseTestSuite) TestValidateAccessToken() {
	accessToken := "accessToken"
	userId := ksuid.New().String()

	testCases := []struct {
		name                 string
		mockUserIdParseToken string
		mockErrParseToken    error
		mockUserGetById      *entities.User
		mockErrGetById       error
		expectedResult       *entities.User
		expectedErr          error
	}{
		{
			name:                 "should success validate token",
			mockUserIdParseToken: userId,
			mockUserGetById: &entities.User{
				ID: userId,
			},
			expectedResult: &entities.User{
				ID: userId,
			},
		},
		{
			name:              "should return error when fail to parse token",
			mockErrParseToken: errors.New("failed to parse token"),
			expectedErr:       errors.New("failed to parse token"),
		},
		{
			name:                 "should return error when user not found",
			mockUserIdParseToken: userId,
			mockErrGetById:       exception.NewNotFoundException().SetMessage("user not found"),
			expectedErr:          usecase.ErrInvalidToken,
		},
		{
			name:                 "should return error when occur any error from get user",
			mockUserIdParseToken: userId,
			mockErrGetById:       errors.New("fail to get user by id"),
			expectedErr:          errors.New("fail to get user by id"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.tokenHandler.EXPECT().ParseToken(accessToken).Return(tc.mockUserIdParseToken, tc.mockErrParseToken)

			if tc.mockErrParseToken == nil {
				s.userRepo.EXPECT().GetById(gomock.Any(), tc.mockUserIdParseToken).Return(tc.mockUserGetById, tc.mockErrGetById)
			}

			user, err := s.ucase.ValidateAccessToken(context.Background(), accessToken)
			s.Equal(tc.expectedErr, err)
			if tc.expectedErr == nil {
				s.EqualValues(tc.expectedResult, user)
			} else {
				s.Nil(user)
			}
		})
	}
}
