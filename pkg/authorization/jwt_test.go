package authorization_test

import (
	"testing"
	"time"

	"github.com/resyahrial/go-user-management/pkg/authorization"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type JwtAuthorizationTestSuite struct {
	suite.Suite
	timeDuration time.Duration
	secretKey    string
	auth         *authorization.JwtAuthorization
}

func TestJwtAuthorization(t *testing.T) {
	suite.Run(t, new(JwtAuthorizationTestSuite))
}

func (s *JwtAuthorizationTestSuite) SetupTest() {
	s.timeDuration = 10 * time.Second
	s.secretKey = "secret"
	s.auth = authorization.NewJwtAuthorization(s.timeDuration, s.secretKey)
}

func (s *JwtAuthorizationTestSuite) TestSignToken() {
	testCases := []struct {
		name          string
		claims        map[string]interface{}
		expectedError error
	}{
		{
			name: "should success sign token",
			claims: map[string]interface{}{
				"id": ksuid.New().String(),
			},
		},
		{
			name:          "should return error when claims not contain id",
			claims:        map[string]interface{}{},
			expectedError: authorization.ErrInvalidToken,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			token, err := s.auth.SignToken(tc.claims)
			s.Equal(tc.expectedError, err)
			if tc.expectedError == nil {
				s.NotEmpty(token)
			} else {
				s.Empty(token)
			}
		})
	}
}
