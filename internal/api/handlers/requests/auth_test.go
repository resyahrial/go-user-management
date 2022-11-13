package request_test

import (
	"testing"

	request "github.com/resyahrial/go-user-management/internal/api/handlers/requests"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/resyahrial/go-user-management/pkg/exception"
	"github.com/stretchr/testify/suite"
)

type AuthRequestTestSuite struct {
	suite.Suite
}

func TestAuthRequest(t *testing.T) {
	suite.Run(t, new(AuthRequestTestSuite))
}

func (s *AuthRequestTestSuite) SetupTest() {
}

func (s *AuthRequestTestSuite) TestConvertLoginRequestToLoginEntity() {
	testCases := []struct {
		name           string
		input          *request.LoginRequest
		expectedOutput *entities.Login
		expectedError  error
	}{
		{
			name: "should create login entity",
			input: &request.LoginRequest{
				Email:    "user@mail.com",
				Password: "anypassword",
			},
			expectedOutput: &entities.Login{
				Email:    "user@mail.com",
				Password: "anypassword",
			},
		},
		{
			name: "should return error when not pass validation",
			input: &request.LoginRequest{
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
		login, err := tc.input.CastToLoginEntity()
		s.Run(tc.name, func() {
			if tc.expectedError == nil {
				s.Nil(err)
			} else {
				s.Equal(tc.expectedError.Error(), err.Error())
			}
			s.EqualValues(tc.expectedOutput, login)
		})
	}
}
