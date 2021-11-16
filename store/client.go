package store

import (
	"context"
	"errors"
	"sync"

	"github.com/luenci/oauth2/service"
)

// NewClientStore create client store
func NewClientStore() *clientStore {
	return &clientStore{
		data: make(map[string]service.ClientInfoInterface),
	}
}

// clientStore client information store
type clientStore struct {
	sync.RWMutex
	data map[string]service.ClientInfoInterface
}

// GetByID according to the ID for the client information
func (cs *clientStore) GetByID(ctx context.Context, id string) (service.ClientInfoInterface, error) {
	cs.RLock()
	defer cs.RUnlock()

	if c, ok := cs.data[id]; ok {
		return c, nil
	}
	return nil, errors.New("not found")
}

// Set set client information
func (cs *clientStore) Set(id string, cli service.ClientInfoInterface) (err error) {
	cs.Lock()
	defer cs.Unlock()

	cs.data[id] = cli
	return
}
