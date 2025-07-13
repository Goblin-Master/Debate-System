package jwtx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type MyClaims struct {
	Userid int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// JWT 认证需要的密钥和过期时间配置
type Auth struct {
	AccessSecret string `json:"accessSecret"`
	AccessExpire int64  `json:"accessExpire"`
}

type Claims struct {
	Auth
	UserID int64 `json:"user_id"`
}

var (
	DEFAULT_ERROR  = errors.New("jwt默认错误")
	TOKEN_EMPTY    = errors.New("token为空")
	TOKEN_EXPERIED = errors.New("token已过期")
	TOKEN_INVALID  = errors.New("token无效")
)

func GenToken(c Claims) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		c.UserID,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(c.Auth.AccessExpire) * time.Second)), // 过期时间
			Issuer:    "Debate-System",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(c.Auth.AccessSecret))
}
func GetUserID(ctx context.Context) (int64, error) {
	t := ctx.Value("user_id")
	user_id, err := t.(json.Number).Int64()
	return user_id, err
}

func ParseToken(secret string, request *http.Request) (int64, error) {
	data := request.Header.Get("Authorization")
	if data == "" {
		return 0, TOKEN_EMPTY
	}
	token := strings.TrimPrefix(data, "Bearer ")
	// 解析token
	var claims MyClaims
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return 0, TOKEN_EXPERIED
		}
		if strings.Contains(err.Error(), "signature is invalid") {
			return 0, TOKEN_INVALID
		}
		if strings.Contains(err.Error(), "token contains an invalid") {
			return 0, TOKEN_INVALID
		}
		fmt.Println(err)
		return 0, DEFAULT_ERROR
	}
	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims.Userid, nil
	}
	return 0, DEFAULT_ERROR
}
