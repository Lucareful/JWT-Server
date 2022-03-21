package service

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type AuthorizationService interface {
	GenerateAuthorizationCode(ctx context.Context, ClientID string) (int, error)
	GenerateAccessToken(ctx context.Context, ClientID string) (string, error)
}

type JWTService interface {
	GenerateToken(userName, password string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
