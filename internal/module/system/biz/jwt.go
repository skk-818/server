package biz

import (
	"server/internal/core/config"
	"server/pkg/jwtx"
)

type (
	JwtUsecase struct {
		cfg *config.Jwt
	}

	jwtUsecase interface {
		GenerateAccessToken(uint, string, []string) (string, error)
		GenerateRefreshToken(uint, string, []string) (string, error)
	}
)

func NewJwtUsecase(
	cfg *config.Jwt,
) *JwtUsecase {
	return &JwtUsecase{
		cfg: cfg,
	}
}

func (ju *JwtUsecase) Parse(string string) (*jwtx.CustomClaims, error) {
	j := jwtx.New(ju.cfg.Secret, ju.cfg.AccessExpire, ju.cfg.RefreshExpire)
	return j.ParseToken(string)
}

func (ju *JwtUsecase) GenerateAccessToken(userID uint, username string, roles []string) (string, error) {
	j := jwtx.New(ju.cfg.Secret, ju.cfg.AccessExpire, ju.cfg.RefreshExpire)
	return j.GenerateAccessToken(userID, username, roles)
}

func (ju *JwtUsecase) GenerateRefreshToken(userID uint, username string, roles []string) (string, error) {
	j := jwtx.New(ju.cfg.Secret, ju.cfg.AccessExpire, ju.cfg.RefreshExpire)
	return j.GenerateRefreshToken(userID, username, roles)
}
