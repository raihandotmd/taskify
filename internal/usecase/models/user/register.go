package user

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
)

type (
	RegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		AccessToken string `json:"access_token"`
	}
)

func (r RegisterRequest) ToObModel() obModel.User {
	return obModel.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
