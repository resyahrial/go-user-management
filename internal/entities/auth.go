package entities

import "context"

const (
	AuthModule = "AUTH"
)

type Login struct {
	Email    string
	Password string
}

type Token struct {
	Access string
}

type Authentication struct {
	CurrentUserID  string
	ResourceUserID string
	Resource       string
	Action         string
}

func (a *Authentication) IsPermissionValid(user *User) bool {
	var (
		permission *Permission
	)
	if a.CurrentUserID != user.ID {
		return false
	}

	for _, p := range user.Role.Permissions {
		if p.Resource == a.Resource && p.Action == a.Action {
			permission = p
			break
		}
	}
	if permission == nil {
		return false
	}

	if !permission.IsGlobalPermission() && a.ResourceUserID != user.ID {
		return false
	}
	return true
}

type AuthUsecase interface {
	Login(ctx context.Context, input *Login) (token *Token, err error)
	ValidateAccessToken(ctx context.Context, accessToken string) (user *User, err error)
	ValidateUserAccess(ctx context.Context, authentication *Authentication) (err error)
}
