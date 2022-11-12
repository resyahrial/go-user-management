package response_test

import (
	"testing"

	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/stretchr/testify/suite"
)

type CreateUserResponseTestSuite struct {
	suite.Suite
}

func TestCreateUserResponse(t *testing.T) {
	suite.Run(t, new(CreateUserResponseTestSuite))
}

func (s *CreateUserResponseTestSuite) SetupTest() {
}

func (s *CreateUserResponseTestSuite) TestConvertToUserEntity() {
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
