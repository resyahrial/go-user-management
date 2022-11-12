package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/resyahrial/go-user-management/internal/entities"
	usecase "github.com/resyahrial/go-user-management/internal/usecase/user"
	adapter_mock "github.com/resyahrial/go-user-management/internal/usecase/user/mocks"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type GetUserListUsecaseTestSuite struct {
	suite.Suite
	userRepo *adapter_mock.MockUserRepo
	hasher   *adapter_mock.MockHasher
	ucase    entities.UserUsecase
}

func TestGetUserListUsecase(t *testing.T) {
	suite.Run(t, new(GetUserListUsecaseTestSuite))
}

func (s *GetUserListUsecaseTestSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = adapter_mock.NewMockUserRepo(ctrl)
	s.hasher = adapter_mock.NewMockHasher(ctrl)
	s.ucase = usecase.NewUserUsecase(
		s.userRepo,
		s.hasher,
	)
}

func (s *GetUserListUsecaseTestSuite) TestGetUserList() {
	userId := ksuid.New().String()
	queryParams := &entities.PaginatedQueryParams{
		Page:  0,
		Limit: 10,
	}

	testCases := []struct {
		name                  string
		resultMockGetUserList []*entities.User
		countMockGetUserList  int64
		errorMockGetUserList  error
		expectedOutput        []*entities.User
		expectedCount         int64
		expectedError         error
	}{
		{
			name: "should get user list",
			resultMockGetUserList: []*entities.User{
				{
					ID:       userId,
					Name:     "user",
					Email:    "user@mail.com",
					Password: "anypassword",
				},
			},
			countMockGetUserList: 1,
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
			s.userRepo.EXPECT().GetList(gomock.Any(), queryParams).Return(tc.resultMockGetUserList, tc.countMockGetUserList, tc.errorMockGetUserList)

			res, count, err := s.ucase.GetList(context.Background(), queryParams)
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
