package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JWTClaims struct {
	UserID uint64   `json:"user_id"`
	Roles  []uint32 `json:"roles"` // 字段名改为复数形式，类型改为 []uint32
	jwt.RegisteredClaims
}

// 生成JWT Token
func GenerateJWT(userID uint64, roles []uint32) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Roles:  roles, // 接收 []uint32 角色数组
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 过期时间
			Issuer:    "gomall",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// 解析JWT Token
func ParseJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
