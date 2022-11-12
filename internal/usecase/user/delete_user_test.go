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

type DeleteUserUsecaseTestSuite struct {
	suite.Suite
	userRepo *adapter_mock.MockUserRepo
	hasher   *adapter_mock.MockHasher
	ucase    entities.UserUsecase
}

func TestDeleteUsecase(t *testing.T) {
	suite.Run(t, new(DeleteUserUsecaseTestSuite))
}

func (s *DeleteUserUsecaseTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = adapter_mock.NewMockUserRepo(ctrl)
	s.hasher = adapter_mock.NewMockHasher(ctrl)
	s.ucase = usecase.NewUserUsecase(
		s.userRepo,
		s.hasher,
	)
}

func (s *DeleteUserUsecaseTestSuite) TestDeleteUser() {
	userId := ksuid.New().String()
	hashedPassword := "hashedPassword"

	testCases := []struct {
		name                    string
		resultMockGetUserDetail *entities.User
		errorMockGetUserDetail  error
		expectedOutput          *entities.User
		expectedError           error
	}{
		{
			name: "should get user detail",
			resultMockGetUserDetail: &entities.User{
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
			name:                   "should return error when failed get user detail",
			errorMockGetUserDetail: errors.New("user not found"),
			expectedError:          errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.userRepo.EXPECT().GetById(gomock.Any(), userId).Return(tc.resultMockGetUserDetail, tc.errorMockGetUserDetail)

			res, err := s.ucase.GetDetail(context.Background(), userId)
			s.Equal(tc.expectedError, err)
			if err == nil {
				s.EqualValues(tc.expectedOutput, res)
			} else {
				s.Nil(res)
			}
		})
	}
}
