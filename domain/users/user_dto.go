package users

import (
	"github.com/comfysweet/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DateCreating string `json:"date_creating"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		errors.NewBadRequestError("invalid email address")
	}
	return nil
}
