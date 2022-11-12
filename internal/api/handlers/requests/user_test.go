package request_test

import (
	"testing"

	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/stretchr/testify/suite"
)

type UserRequestTestSuite struct {
	suite.Suite
}

func TestUserRequest(t *testing.T) {
	suite.Run(t, new(UserRequestTestSuite))
}

func (s *UserRequestTestSuite) SetupTest() {
}

func (s *UserRequestTestSuite) TestConvertCreateUserRequestToUserEntity() {
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

func (s *UserRequestTestSuite) TestConvertUpdateUserRequestToUserEntity() {
	testCases := []struct {
		name           string
		input          *request.UpdateUserRequest
		expectedOutput *entities.User
		expectedError  error
	}{
		{
			name: "should success convert user",
			input: &request.UpdateUserRequest{
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
			input: &request.UpdateUserRequest{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "@newPass123",
			},
			expectedError: exception.NewBadRequestException().SetModule(entities.UserModule).SetCollectionMessage(map[string][]string{
				"password": {
					"Password can only contain alphanumeric characters",
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
