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
}

func TestJwtAuthorization(t *testing.T) {
	suite.Run(t, new(JwtAuthorizationTestSuite))
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
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			auth := authorization.NewJwtAuthorization(10*time.Second, "secret")
			token, err := auth.SignToken(tc.claims)
			s.Equal(tc.expectedError, err)
			if tc.expectedError == nil {
				s.NotEmpty(token)
			} else {
				s.Empty(token)
			}
		})
	}
}

func (s *JwtAuthorizationTestSuite) TestParseToken() {
	testCases := []struct {
		name          string
		timeDuration  time.Duration
		secretKey     string
		id            string
		expectedError error
	}{
		{
			name:         "should success parse token",
			timeDuration: 10 * time.Second,
			secretKey:    "secret",
			id:           ksuid.New().String(),
		},
		{
			name:          "should return error when token not contains id",
			timeDuration:  10 * time.Second,
			secretKey:     "secret",
			expectedError: authorization.ErrInvalidToken,
		},
		{
			name:          "should return error when token expired",
			timeDuration:  1 * time.Millisecond,
			secretKey:     "secret",
			id:            ksuid.New().String(),
			expectedError: authorization.ErrInvalidToken,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			auth := authorization.NewJwtAuthorization(tc.timeDuration, tc.secretKey)
			claims := make(map[string]interface{})
			if tc.id != "" {
				claims["id"] = tc.id
			}
			token, _ := auth.SignToken(claims)
			s.NotEmpty(token)

			// simulate token expired
			delay := 1 * time.Second
			if tc.timeDuration < delay {
				time.Sleep(delay)
			}

			resId, err := auth.ParseToken(token)
			s.Equal(tc.expectedError, err)
			if tc.expectedError == nil {
				s.Equal(tc.id, resId)
			} else {
				s.Empty(resId)
			}
		})
	}
}
