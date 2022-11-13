package response_test

import (
	"testing"

	response "github.com/resyahrial/go-user-management/internal/api/handlers/responses"
	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/stretchr/testify/suite"
)

type AuthResponseTestSuite struct {
	suite.Suite
}

func TestAuthResponse(t *testing.T) {
	suite.Run(t, new(AuthResponseTestSuite))
}

func (s *AuthResponseTestSuite) SetupTest() {
}

func (s *AuthResponseTestSuite) TestConvertTokenEntityToLoginResponse() {
	token := &entities.Token{
		Access: "token",
	}

	testCases := []struct {
		name           string
		expectedOutput *response.LoginResponse
		expectedError  error
	}{
		{
			name: "should success convert login response",
			expectedOutput: &response.LoginResponse{
				AccessToken: "token",
			},
		},
	}

	for _, tc := range testCases {
		res, err := response.NewLoginResponse(token)
		s.Run(tc.name, func() {
			s.Equal(tc.expectedError, err)
			s.EqualValues(tc.expectedOutput, res)
		})
	}
}
