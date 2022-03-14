package service

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	Authorization AuthorizationService
	JWT           JWTService
}

type AuthorizationService interface {
	GenerateAuthorizationCode(ctx context.Context, ClientID string) (int, error)
	GenerateAccessToken(ctx context.Context, ClientID string) (string, error)
}

type JWTService interface {
	GenerateToken(userId string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

// NewALLService service 注册中心.
func NewALLService() *Service {
	return &Service{
		Authorization: (*authorizationService)(nil),
		JWT: &jwtServices{
			secretKey: getSecretKey(),
			issure:    getName(),
		},
	}
}
