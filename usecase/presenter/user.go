package presenter

import (
	"context"
	"github.com/tunamfv/go-jwt/domain/entity"
)

type UserPresenter interface {
	Render(context.Context, entity.User)
	Error(context.Context, error)
}