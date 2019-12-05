package entity

import (
	"github.com/tunamfv/go-jwt/domain/valueobject"
)

type User struct {
	ID valueobject.UserID `json:"id"`
	Name string `json:"name"`
}