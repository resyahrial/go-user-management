package response_test

import (
	"testing"

	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/entities"
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

func (s *UserResponseTestSuite) TestConvertUserEntityToUserResponse() {
	user := &entities.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "anypassword",
	}

	testCases := []struct {
		name           string
		expectedOutput *response.UserResponse
		expectedError  error
	}{
		{
			name: "should create basic user",
			expectedOutput: &response.UserResponse{
				Name:  "user",
				Email: "user@mail.com",
			},
		},
	}

	for _, tc := range testCases {
		res, err := response.NewUserResponse(user)
		s.Run(tc.name, func() {
			s.Equal(tc.expectedError, err)
			s.EqualValues(tc.expectedOutput, res)
		})
	}
}
