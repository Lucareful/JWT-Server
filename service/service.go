package service

import (
	"time"

	"github.com/luenci/oauth2/models"
)

type Service struct {
	Token  TokenInfoInterface
	Client ClientInfoInterface
}

type (
	// ClientInfoInterface the client information model interface
	ClientInfoInterface interface {
		GetID() string
		GetSecret() string
		GetDomain() string
		GetUserID() string
	}

	// ClientPasswordVerifierInterface the password handler interface
	ClientPasswordVerifierInterface interface {
		VerifyPassword(string) bool
	}

	// TokenInfoInterface the token information model interface
	TokenInfoInterface interface {
		GetClientID() string
		SetClientID(string)
		GetUserID() string
		SetUserID(string)
		GetRedirectURI() string
		SetRedirectURI(string)
		GetScope() string
		SetScope(string)

		GetCode() string
		SetCode(string)
		GetCodeCreateAt() time.Time
		SetCodeCreateAt(time.Time)
		GetCodeExpiresIn() time.Duration
		SetCodeExpiresIn(time.Duration)
		GetCodeChallenge() string
		SetCodeChallenge(string)

		GetAccess() string
		SetAccess(string)
		GetAccessCreateAt() time.Time
		SetAccessCreateAt(time.Time)
		GetAccessExpiresIn() time.Duration
		SetAccessExpiresIn(time.Duration)

		GetRefresh() string
		SetRefresh(string)
		GetRefreshCreateAt() time.Time
		SetRefreshCreateAt(time.Time)
		GetRefreshExpiresIn() time.Duration
		SetRefreshExpiresIn(time.Duration)
	}
)

func GetService() *Service {
	return &Service{
		Token:  (*models.Token)(nil),
		Client: (*models.Client)(nil),
	}
}
