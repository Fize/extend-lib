package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const title = "fizz"

// Claims jwt object
type Claims struct {
	Spec map[string]interface{}
	jwt.StandardClaims
}

// GenerateJwtToken create a new token for json web
func GenerateJwtToken(claims Claims, exp int64) (string, int64, error) {
	expires := time.Now().Add(time.Second * time.Duration(exp)).Unix()
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(title))
	return token, expires, err
}

// ParseJwtToken validate jwt token
func ParseJwtToken(token string) (interface{}, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(title), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
