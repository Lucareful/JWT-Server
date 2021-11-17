package service

import (
	"context"
	"net/http"
	"time"

	"github.com/luenci/oauth2/models"

	pkg "github.com/luenci/oauth2/pkg"
)

// TokenGenerateRequest provide to generate the token request parameters
type TokenGenerateRequest struct {
	ClientID            string
	ClientSecret        string
	UserID              string
	RedirectURI         string
	Scope               string
	Code                string
	CodeChallenge       string
	CodeChallengeMethod pkg.CodeChallengeMethod
	Refresh             string
	CodeVerifier        string
	AccessTokenExp      time.Duration
	Request             *http.Request
}

// Manager authorization management interface
type Manager interface {
	// GetClient get the client information
	GetClient(ctx context.Context, clientID string) (cli models.ClientInfo, err error)

	// GenerateAuthToken generate the authorization token(code)
	GenerateAuthToken(ctx context.Context, rt pkg.ResponseType, tgr *TokenGenerateRequest) (authToken models.TokenInfo, err error)

	// GenerateAccessToken generate the access token
	GenerateAccessToken(ctx context.Context, rt pkg.GrantType, tgr *TokenGenerateRequest) (accessToken models.TokenInfo, err error)

	// RefreshAccessToken refreshing access token
	RefreshAccessToken(ctx context.Context, tgr *TokenGenerateRequest) (accessToken models.TokenInfo, err error)

	// RemoveAccessToken use the access token to delete the token information
	RemoveAccessToken(ctx context.Context, access string) (err error)

	// RemoveRefreshToken use the refresh token to delete the token information
	RemoveRefreshToken(ctx context.Context, refresh string) (err error)

	// LoadAccessToken according to the access token for corresponding token information
	LoadAccessToken(ctx context.Context, access string) (ti models.TokenInfo, err error)

	// LoadRefreshToken according to the refresh token for corresponding token information
	LoadRefreshToken(ctx context.Context, refresh string) (ti models.TokenInfo, err error)
}
