package auth

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// token 标志
	JWT_AUTHORIZATION_KEY = "AUTHORIZATION"
	JWT_BEARER            = "Bearer "

	// claims 标志
	JWT_EXP = "exp" //有效期
	JWT_IAT = "iat" //签发时间
	JWT_SUB = "sub" //用户ID
	JWT_AUD = "aud" //受众
	JWT_NBF = "nbf" //生效时间
	JWT_ISS = "iss" //签发者

	JWT_ROLES            = "roles"             // 用户角色 暂
	JWT_CLIENTID         = "CLIENT_ID"         // 客户端
	JWT_USERNAME         = "USER_NAME"         // 用户名
	JWT_USERID           = "USER_ID"           // 用户ID
	JWT_USERIDENTITYNAME = "USER_IDENTITYNAME" // 用户唯一ID
)

func NewMapClaims(head map[string]any) jwt.MapClaims {
	return jwt.MapClaims{
		JWT_EXP: time.Now().Add(time.Hour * 24 * 7).Unix(), // 有效期七天
		JWT_IAT: time.Now().Unix(),
		JWT_SUB: head[UID],
		JWT_AUD: head[AUD],
		JWT_NBF: time.Now().Add(-time.Minute).Unix(),
		JWT_ISS: "auth-jwt",

		// 自定义信息
		// JWT_ROLES
		JWT_CLIENTID:         head[CLIENTID],
		JWT_USERNAME:         head[USERNAME],
		JWT_USERID:           head[UID],
		JWT_USERIDENTITYNAME: head[USERIDENTITYNAME],
	}
}

func GenerateTokens(claims jwt.MapClaims) (token string, refreshToken string) {

	// create token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JwtConfig.Jwt.Secret))
	if err != nil {
		// todo log here
		log.Printf("create token failed: %s", err.Error())
		return "", ""
	}

	// save info
	refreshClaims := NewMapClaims(map[string]any{})
	refreshClaims[JWT_EXP] = time.Now().Add(time.Hour * 24 * 7 * 2).Unix() // 有效期两周

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JwtConfig.Jwt.Secret))
	if err != nil {
		// todo log here
		log.Printf("create refresh token failed: %s", err.Error())
		return "", ""
	}

	return
}

func ExtractMapClaims(tokenStr string) (jwt.MapClaims, error) {

	token, err := ValidateToken(tokenStr)

	if err != nil {
		// todo log
		log.Printf("ExtractMapClaims err: %s", err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// successed convert to map claims
		return claims, nil

	} else {

		// todo log
		log.Printf("ExtractMapClaims err: claims failed ok:[%t] valid:[%t]", ok, token.Valid)
		return nil, err
	}
}

// 验证 Token
func ValidateToken(tokenStr string) (*jwt.Token, error) {
	tokenStr = trimBearer(tokenStr)

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回验证密钥
		return []byte(JwtConfig.Jwt.Secret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
}

func trimBearer(tokenStr string) string {
	return strings.TrimPrefix(tokenStr, JWT_BEARER)
}
