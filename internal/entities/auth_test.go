package entities_test

import (
	"testing"

	"github.com/resyahrial/go-user-management/internal/entities"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type AuthEntitiesTestSuite struct {
	suite.Suite
}

func TestLoginUsecase(t *testing.T) {
	suite.Run(t, new(AuthEntitiesTestSuite))
}

func (s *AuthEntitiesTestSuite) SetupTest() {
}

func (s *AuthEntitiesTestSuite) TestAuthenticationAccessPermission() {
	userId := ksuid.New().String()
	otherUserId := ksuid.New().String()

	testCases := []struct {
		name           string
		authentication *entities.Authentication
		user           *entities.User
		ok             bool
	}{
		{
			name: "should give permission to access their resource if user have global permission",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: userId,
				Resource:       "users",
				Action:         "WRITE",
			},
			user: &entities.User{
				ID: userId,
				Role: &entities.Role{
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "WRITE",
							Type:     "global",
						},
					},
				},
			},
			ok: true,
		},
		{
			name: "should give permission to access others resource if user have global permission",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: otherUserId,
				Resource:       "users",
				Action:         "WRITE",
			},
			user: &entities.User{
				ID: userId,
				Role: &entities.Role{
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "WRITE",
							Type:     "GLOBAL",
						},
					},
				},
			},
			ok: true,
		},
		{
			name: "should not give permission to access others resource if user didn't have permission",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: otherUserId,
				Resource:       "users",
				Action:         "WRITE",
			},
			user: &entities.User{
				ID: userId,
				Role: &entities.Role{
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "READ",
							Type:     "GLOBAL",
						},
					},
				},
			},
		},
		{
			name: "should not give permission to access others resource if user only have exclusive permission",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: otherUserId,
				Resource:       "users",
				Action:         "WRITE",
			},
			user: &entities.User{
				ID: userId,
				Role: &entities.Role{
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "WRITE",
							Type:     "EXCLUSIVE",
						},
					},
				},
			},
		},
		{
			name: "should not give permission to access if user passed is different with current user",
			authentication: &entities.Authentication{
				CurrentUserID:  userId,
				ResourceUserID: otherUserId,
				Resource:       "users",
				Action:         "WRITE",
			},
			user: &entities.User{
				ID: otherUserId,
				Role: &entities.Role{
					Permissions: []*entities.Permission{
						{
							Resource: "users",
							Action:   "WRITE",
							Type:     "GLOBAL",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Equal(tc.ok, tc.authentication.IsPermissionValid(tc.user))
		})
	}
}
