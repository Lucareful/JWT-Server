package service

import (
	"context"

	"github.com/luenci/oauth2/store"

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

// Service defines functions used to return resource interface.
type Service interface {
	JWT() JWTService
	Authorization() AuthorizationService
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) JWT() JWTService {
	return newJWTServices(s)
}

func (s *service) Authorization() AuthorizationService {
	return newAuthorizationService(s)
}
