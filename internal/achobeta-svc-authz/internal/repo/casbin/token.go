package casbin

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("btk.AchoBeta7VZD1dar")

func (c *impl) CreateToken(sub, obj, dom string) (string, error) {
	// 创建负载
	payload := jwt.MapClaims{
		"userId":  sub,
		"object":  obj,
		"domain":  dom,
		"exptime": time.Now().Add(time.Minute * 15).Unix(), // 令牌过期时间为15分钟
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// 签名并获取完整的 Token 字符串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (c *impl) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := parseToken(tokenStr)
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	// 获取负载信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token claims error")
	}
	return claims, nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	// 解析 Token 字符串
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
