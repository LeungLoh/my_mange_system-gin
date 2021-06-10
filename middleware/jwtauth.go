package middleware

import (
	"fmt"
	"my_mange_system/common"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var MySecret []byte

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

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("token")
		if token == "" {
			common.UnAuthorized(c)
			return
		}
		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(token)
		if err != nil {
			common.UnAuthorized(c)
			return
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)

	}
}
