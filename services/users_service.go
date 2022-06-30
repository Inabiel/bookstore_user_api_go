package services

import (
	"github.com/Inabiel/bookstore_user_api_go/domain/users"
	"github.com/Inabiel/bookstore_user_api_go/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestfulError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}
