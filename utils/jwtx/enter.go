package jwtx

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
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
