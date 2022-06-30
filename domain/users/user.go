package users

import (
	"strings"

	"github.com/Inabiel/bookstore_user_api_go/utils/errors"
)

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func (user *User) Validate() *errors.RestfulError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" && !strings.Contains(user.Email, "@") {
		return errors.NewBadRequestError("email is not valid")
	}

	return nil
}
