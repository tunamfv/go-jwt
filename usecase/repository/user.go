package repository

import (
	"context"
	"github.com/tunamfv/go-jwt/usecase/payload"
	"github.com/tunamfv/go-jwt/domain/entity"
)

type UserRepository interface {
	Create(context.Context, payload.CreateUserPayload) (entity.User, error)
	Update(context.Context, payload.UpdateUserPayload) (entity.User, error)
}