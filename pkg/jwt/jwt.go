package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"server/pkg/xerror"
	"time"
)

type Jwt struct {
	secretKey          []byte
	accessTokenExpire  int64
	refreshTokenExpire int64
}

func New(secret string, accessExpire, refreshExpire int64) *Jwt {
	return &Jwt{
		secretKey:          []byte(secret),
		accessTokenExpire:  accessExpire,
		refreshTokenExpire: refreshExpire,
	}
}

func (j *Jwt) GenerateAccessToken(userID uint, username, role string) (string, error) {
	return j.GenerateToken(userID, username, role, time.Duration(j.accessTokenExpire)*time.Second)
}

func (j *Jwt) GenerateRefreshToken(userID uint, username, role string) (string, error) {
	return j.GenerateToken(userID, username, role, time.Duration(j.refreshTokenExpire)*time.Second)
}

func (j *Jwt) GenerateToken(userID uint, username, role string, expire time.Duration) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *Jwt) ParseToken(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	// 明确处理各类错误（v5）
	if err != nil {
		// 判断是否为 token 过期
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, xerror.ErrTokenExpired
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, xerror.ErrTokenMalformed
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, xerror.ErrTokenNotValidYet
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, xerror.ErrTokenSignatureInvalid
		}
		return nil, xerror.ErrTokenParseFailed
	}

	// token 是否有效
	if !token.Valid {
		return nil, xerror.ErrTokenInvalid
	}

	// 校验 claims 是否正确
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, xerror.ErrTokenExpired
	}

	return claims, nil
}
