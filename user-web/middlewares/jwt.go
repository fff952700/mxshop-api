package middlewares

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// JWTAuth 中间件，验证 JWT Token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			respondWithError(c, http.StatusUnauthorized, "请登录")
			return
		}

		jwtHandler := NewJWT()
		claims, err := jwtHandler.ParseToken(token)
		if err != nil {
			handleTokenError(c, err)
			return
		}

		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

// JWT 结构体，包含密钥
type JWT struct {
	SigningKey []byte
}

// NewJWT 返回一个新的 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConf.JWTInfo.SigningKey),
	}
}

// CreateToken 创建一个新的 JWT Token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 JWT Token 并返回自定义声明
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, parseTokenError(err)
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 刷新 Token 的过期时间并生成新的 Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour)) // 更新过期时间
		return j.CreateToken(*claims)
	}

	return "", TokenInvalid
}

// handleTokenError 处理不同类型的 Token 错误
func handleTokenError(c *gin.Context, err error) {
	if errors.Is(err, TokenExpired) {
		respondWithError(c, http.StatusUnauthorized, "授权已过期")
		return
	}
	respondWithError(c, http.StatusUnauthorized, "未登录")
}

// respondWithError 通用的错误响应函数
func respondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, map[string]string{"msg": message})
	c.Abort()
}

// parseTokenError 解析 Token 错误
func parseTokenError(err error) error {
	if err != nil {
		// 如果错误是 jwt.ErrSignatureInvalid，表示签名验证失败
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return TokenInvalid
		}

		// 如果错误是 jwt.ErrTokenExpired，表示 Token 已过期
		if errors.Is(err, jwt.ErrTokenExpired) {
			return TokenExpired
		}

		// 如果错误是 jwt.ErrTokenNotValidYet，表示 Token 尚未生效
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return TokenNotValidYet
		}
	}

	return TokenInvalid
}
