package request_test

import (
	"testing"

	request "github.com/resyahrial/go-template/internal/api/handlers/requests"
	"github.com/resyahrial/go-template/internal/entities"
	"github.com/resyahrial/go-template/pkg/exception"
	"github.com/stretchr/testify/suite"
)

type CreateUserRequestTestSuite struct {
	suite.Suite
}

func TestCreateUserRequest(t *testing.T) {
	suite.Run(t, new(CreateUserRequestTestSuite))
}

func (s *CreateUserRequestTestSuite) SetupTest() {
}

func (s *CreateUserRequestTestSuite) TestConvertToUserEntity() {
	testCases := []struct {
		name           string
		input          *request.CreateUserRequest
		expectedOutput *entities.User
		expectedError  error
	}{
		{
			name: "should create basic user",
			input: &request.CreateUserRequest{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			expectedOutput: &entities.User{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
		{
			name: "should return error when not pass validation",
			input: &request.CreateUserRequest{
				Name:     "user",
				Email:    "failemail",
				Password: "anypassword",
			},
			expectedError: exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(map[string][]string{
				"email": {
					"Email must be a valid email address",
				},
			}),
		},
	}

	for _, tc := range testCases {
		user, err := tc.input.CastToUserEntity()
		s.Run(tc.name, func() {
			if tc.expectedError == nil {
				s.Nil(err)
			} else {
				s.Equal(tc.expectedError.Error(), err.Error())
			}
			s.EqualValues(tc.expectedOutput, user)
		})
	}
}
