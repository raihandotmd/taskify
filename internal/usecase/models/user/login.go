package user

import obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}
)

func (r LoginRequest) ToObModel() obModel.User {
	return obModel.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
