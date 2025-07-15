package model

import ucModel "github.com/raihandotmd/taskify/internal/usecase/models/user"

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *RegisterRequest) ToUcModel() ucModel.RegisterRequest {
	return ucModel.RegisterRequest{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
