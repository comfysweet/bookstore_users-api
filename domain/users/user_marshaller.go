package users

import (
	"encoding/json"
	"github.com/comfysweet/bookstore_utils-go/errors"
)

type PublicUser struct {
	Id           int64  `json:"id"`
	DateCreating string `json:"date_creating"`
	Status       string `json:"status"`
}

type PrivateUser struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	DateCreating string `json:"date_creating"`
	Status       string `json:"status"`
}

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshal(isPublic)
	}
	return result
}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:           user.Id,
			DateCreating: user.DateCreating,
			Status:       user.Status,
		}
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return errors.NewInternalServiceError("marshal json error", errors.NewError("internal service error"))
	}
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return errors.NewInternalServiceError("unmarshal json error", errors.NewError("internal service error"))
	}
	return privateUser
}
