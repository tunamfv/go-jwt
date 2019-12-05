package interactor

import (
	"context"
	"github.com/tunamfv/go-jwt/usecase/repository"
	"github.com/tunamfv/go-jwt/usecase/repsenter"
	"github.com/tunamfv/go-jwt/usecase/repsenter"
)

type UserInteractor interface {
	SignUp(context.Context, payload.CreateUserPayload)
	Update(context.Context, payload.UpdateUserPayload)
}

type userInteractor struct {
	r repository.UserRepository
	p presenter.UserPresenter
}

func NewUserPresenter(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return userInteractor{
		r: r,
		p: p,
	}
}

func (i userInteractor) SignUp(ctx context.Context, payload payload.CreateUserPayload) {

}

func (i userInteractor) Update(ctx context.Context, payload payload.UpdateUserPayload) {

}

