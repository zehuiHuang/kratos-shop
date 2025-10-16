package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID          int64
	NickName    string
	AuthorityId int
	jwt.RegisteredClaims
}

// CreateToken generate token
func CreateToken(c CustomClaims, key string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate token failed" + err.Error())
	}
	return signedString, nil
}
