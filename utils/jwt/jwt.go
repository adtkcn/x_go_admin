package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

/*
*
UserId	000001	自定义字段, 用户ID, 表示这个 jwt 作用于特定用户
UserName	Tom	自定义字段, 用户名, 表示这个 jwt 作用于特定用户
GrantScope	read_user_info	自定义字段, 授权范围, 标识这个 jwt 能够干啥
Issuer	Auth_Server	标准字段, jwt 签名方, 表示是谁签发的这个 jwt
Subject	Tom	标准字段, 表示这个 jwt 作用对象, 在这里与 Username 等效, 再写一遍方便理解
Audience	jwt.ClaimStrings{"Android_APP", "IOS_APP"}	标准字段, 表示jwt 签发给谁, 比如后端某个服务(Auth_Server)签发给客户端(Android_APP, IOS_APP)使用
ExpiresAt	jwt.NewNumericDate(time.Now().Add(time.Hour))	标准字段, jwt 过期时间点
NotBefore	jwt.NewNumericDate(time.Now().Add(time.Hour))	标准字段, jwt 最早的有效时间点, 早于这个时间点无效
IssuedAt	jwt.NewNumericDate(time.Now().Add(time.Hour))	标准字段, jwt 的签发时间点
ID	随机数	标准字段, jwt的ID， 尽量唯一, 我理解为类似于在Hash之前加盐值, 更加防碰撞
*/
type MyCustomClaims struct {
	UserID int
	// exp    int64
	// iat    int64
	// nbf    int64
	jwt.RegisteredClaims
}

// 使用密钥对 token 进行签名
var secret = []byte("secretkey") // 将此处的 secret 替换为你的真实密钥

// 生成 JWT token
func GenerateToken(userId int) (string, error) {
	// 设置过期时间为10天后
	// expirationTime := time.Now().Add(10 * 24 * time.Hour)

	// 创建私有声明
	claims := MyCustomClaims{
		UserID: userId,
		// exp:    expirationTime.Unix(),
		// iat:    time.Now().Unix(),
		// nbf:    time.Now().Unix(),
	}

	// 创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
