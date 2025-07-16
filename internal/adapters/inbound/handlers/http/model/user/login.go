package model

import ucModel "github.com/raihandotmd/taskify/internal/usecase/models/user"

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
)

func (r *LoginRequest) ToUcModel() ucModel.LoginRequest {
	return ucModel.LoginRequest{
		Email:    r.Email,
		Password: r.Password,
	}
}
