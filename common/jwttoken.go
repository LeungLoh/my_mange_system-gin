package common

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("leungloh"),
	}
}

type CustomClaims struct {
	Username string
	jwt.StandardClaims
}

func NewCustomClaims(username string) CustomClaims {
	return CustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(300 * time.Second).Unix(), // 过期时间
			Issuer:    "leungloh",                               // 签发人
		},
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token无效")
}

func GenerateToken(username string) (string, error) {
	jwt := NewJWT()
	claims := NewCustomClaims(username)
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return token, err
	}
	return token, nil
}
