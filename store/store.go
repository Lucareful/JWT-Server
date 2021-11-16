package store

import (
	"context"

	"github.com/luenci/oauth2/service"
)

type (
	// ClientStore the client information storage interface
	ClientStore interface {
		// GetByID according to the ID for the client information
		GetByID(ctx context.Context, id string) (service.ClientInfoInterface, error)
	}

	// tokenStore the token information storage interface
	TokenStore interface {
		// Create and store the new token information
		Create(ctx context.Context, info service.TokenInfoInterface) error

		// RemoveByCode delete the authorization code
		RemoveByCode(ctx context.Context, code string) error

		// RemoveByAccess use the access token to delete the token information
		RemoveByAccess(ctx context.Context, access string) error

		// RemoveByRefresh use the refresh token to delete the token information
		RemoveByRefresh(ctx context.Context, refresh string) error

		// GetByCode use the authorization code for token information data
		GetByCode(ctx context.Context, code string) (service.TokenInfoInterface, error)

		// GetByAccess use the access token for token information data
		GetByAccess(ctx context.Context, access string) (service.TokenInfoInterface, error)

		// GetByRefresh use the refresh token for token information data
		GetByRefresh(ctx context.Context, refresh string) (service.TokenInfoInterface, error)
	}
)
