package service

import "context"

type Service struct {
	Authorization AuthorizationService
}

type AuthorizationService interface {
	GenerateAuthorizationCode(ctx context.Context, ClientID string) (int, error)
	GenerateAccessToken(ctx context.Context, ClientID string) (string, error)
}

func NewALLService() *Service {
	return &Service{Authorization: (*authorizationService)(nil)}
}
