package repository

import (
	"context"

	"github.com/luenci/oauth2/entity"
)

type UserRepositoryInterface interface {
	GetUserByName(ctx context.Context, name string) (entity.User, error)
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	GetUserID(ctx context.Context, name, password string) (entity.User, error)
	Create(ctx context.Context, user entity.User) (entity.User, error)
	Update(ctx context.Context) (entity.User, error)
}
