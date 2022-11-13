package authorization

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/resyahrial/go-user-management/pkg/exception"
)

var (
	ErrInvalidToken = exception.NewAuthorizationException().SetMessage("invalid token")
)

type JwtAuthorization struct {
	tokenDuration time.Duration
	secretKey     string
}

func NewJwtAuthorization(tokenDuration time.Duration, secretKey string) *JwtAuthorization {
	return &JwtAuthorization{tokenDuration, secretKey}
}

func (j *JwtAuthorization) SignToken(claims map[string]interface{}) (tokenString string, err error) {
	currTime := time.Now()
	jwtClaims := jwt.StandardClaims{
		ExpiresAt: currTime.Add(j.tokenDuration).Unix(),
		IssuedAt:  currTime.Unix(),
	}
	if id, ok := claims["id"].(string); ok {
		jwtClaims.Id = id
	} else {
		log.Println("[Jwt - Sign Token] error when construct claims: claims input did not contain id")
		err = ErrInvalidToken
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	if tokenString, err = token.SignedString([]byte(j.secretKey)); err != nil {
		log.Printf("[Jwt - Sign Token] error when signing token: %v\n", err)
		err = ErrInvalidToken
		return
	}
	return
}
