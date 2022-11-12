package response_test

import (
	"testing"

	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type UserResponseTestSuite struct {
	suite.Suite
}

func TestUserResponse(t *testing.T) {
	suite.Run(t, new(UserResponseTestSuite))
}

func (s *UserResponseTestSuite) SetupTest() {
}

func (s *UserResponseTestSuite) TestConvertUserEntityToCreateUserResponse() {
	user := &entities.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name           string
		expectedOutput *response.CreateUserResponse
		expectedError  error
	}{
		{
			name: "should create basic user",
			expectedOutput: &response.CreateUserResponse{
				Name:  "user",
				Email: "user@mail.com",
			},
		},
	}

	for _, tc := range testCases {
		res, err := response.NewCreateUserResponse(user)
		s.Run(tc.name, func() {
			s.Equal(tc.expectedError, err)
			s.EqualValues(tc.expectedOutput, res)
		})
	}
}

func (s *UserResponseTestSuite) TestConvertUserEntityToUpdateUserResponse() {
	user := &entities.User{
		ID:       ksuid.New().String(),
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name           string
		expectedOutput *response.UpdateUserResponse
		expectedError  error
	}{
		{
			name: "should create basic user",
			expectedOutput: &response.UpdateUserResponse{
				ID:    user.ID,
				Name:  "user",
				Email: "user@mail.com",
			},
		},
	}

	for _, tc := range testCases {
		res, err := response.NewCreateUserResponse(user)
		s.Run(tc.name, func() {
			s.Equal(tc.expectedError, err)
			s.EqualValues(tc.expectedOutput, res)
		})
	}
}
