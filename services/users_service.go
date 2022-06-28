package services

import (
	"github.com/Inabiel/bookstore_user_api_go/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
