package usecase

import (
	"context"
	"errors"

	"github.com/alexKudryavtsev-web/beyond-limits-app/config"
	"github.com/golang-jwt/jwt"
)

type AuthUseCase struct {
	adminCfg config.Admin
}

func NewAuthUseCase(adminCfg config.Admin) *AuthUseCase {
	return &AuthUseCase{adminCfg: adminCfg}
}

var _ Auth = (*AuthUseCase)(nil)

func (a *AuthUseCase) Login(ctx context.Context, login, password string) (string, error) {
	if login != a.adminCfg.Login || password != a.adminCfg.Password {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": true,
	})

	return token.SignedString([]byte(a.adminCfg.JWTSecret))
}
