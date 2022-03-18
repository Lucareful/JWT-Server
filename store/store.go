package store

import "github.com/luenci/oauth2/service"

var client Factory

// Factory defines the storage interface.
type Factory interface {
	JWT() service.JWTService
	Authorization() service.AuthorizationService
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
