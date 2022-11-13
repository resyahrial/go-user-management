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

type ValidateUserAccessUsecaseTestSuite struct {
	suite.Suite
	userRepo     *adapter_mock.MockUserRepo
	hasher       *adapter_mock.MockHasher
	tokenHandler *adapter_mock.MockTokenHandler
	ucase        entities.AuthUsecase
}

func TestValidateUserAccessUsecase(t *testing.T) {
	suite.Run(t, new(ValidateUserAccessUsecaseTestSuite))
}

func (s *ValidateUserAccessUsecaseTestSuite) SetupTest() {
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

func (s *ValidateUserAccessUsecaseTestSuite) TestValidateUserAccess() {
	userId := ksuid.New().String()

	testCases := []struct {
		name            string
		authentication  *entities.Authentication
		mockUserGetById *entities.User
		mockErrGetById  error
		expectedError   error
	}{
		{
			name: "should give permission if all permission matched",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: userId,
				Resource:       "users",
				Action:         "WRITE",
			},
			mockUserGetById: &entities.User{
				ID: userId,
				Role: &entities.Role{
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
		{
			name: "should not give permission if user permission not matched",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: userId,
				Resource:       "users",
				Action:         "WRITE",
			},
			mockUserGetById: &entities.User{
				ID: userId,
				Role: &entities.Role{
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "READ",
							Type:     "GLOBAL",
						},
					},
				},
			},
			expectedError: usecase.ErrUnauthenticated,
		},
		{
			name: "should return error when can't get user needed",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: userId,
				Resource:       "users",
				Action:         "WRITE",
			},
			mockErrGetById: errors.New("user not found"),
			expectedError:  errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.userRepo.EXPECT().GetByIdWithPermission(gomock.Any(), tc.authentication.CurrentUserID).Return(tc.mockUserGetById, tc.mockErrGetById)
			err := s.ucase.ValidateUserAccess(context.Background(), tc.authentication)
			s.Equal(tc.expectedError, err)
		})
	}
}
