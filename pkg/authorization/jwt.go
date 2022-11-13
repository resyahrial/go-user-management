package authorization

import (
	"fmt"
	"log"
	"strings"
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
	jwtClaims := jwt.MapClaims{
		"eat": currTime.Add(j.tokenDuration).Unix(),
		"iat": currTime.Unix(),
	}
	if id, ok := claims["id"].(string); ok {
		jwtClaims["id"] = id
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	if tokenString, err = token.SignedString([]byte(j.secretKey)); err != nil {
		log.Printf("[Jwt - Sign Token] error when signing token: %v\n", err)
		err = ErrInvalidToken
		return
	}
	return
}

func (j *JwtAuthorization) ParseToken(tokenString string) (id string, err error) {
	token, errParseToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
	if errParseToken != nil {
		log.Println(errParseToken)
		err = ErrInvalidToken
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	eat, _ := claims["eat"].(float64)
	log.Println(int64(eat), time.Now().Unix())
	if !ok || !token.Valid {
		log.Println("[Jwt - Parse Token] error when check claims: invalid claims")
		return
	}
	if eat, _ := claims["eat"].(float64); int64(eat) < time.Now().Unix() {
		log.Println(eat)
		log.Println("[Jwt - Parse Token] error when check claims: token expired")
		err = ErrInvalidToken
		return
	}
	if claimId, _ := claims["id"].(string); strings.TrimSpace(claimId) == "" {
		log.Println("[Jwt - Parse Token] error when check claims: claims not contain id")
		err = ErrInvalidToken
		return
	} else {
		id = claimId
	}
	return
}
