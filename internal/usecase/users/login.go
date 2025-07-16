package users

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	authModel "github.com/raihandotmd/taskify/internal/usecase/models/token"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/user"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
	"golang.org/x/crypto/bcrypt"
)

func (uc *ucUser) Login(ctx context.Context, req ucModel.LoginRequest) (ucModel.LoginResponse, error) {
	var orm = taskifyGorm.DB

	var (
		user obModel.User
		err  error
	)

	// find user by email
	if user, err = repoUser.GetUserByEmail(ctx, orm, req.Email); err != nil {
		return ucModel.LoginResponse{}, err
	}

	// compare password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return ucModel.LoginResponse{}, err
	}

	// generate access token
	accessToken, err := ucToken.Generate(ctx, authModel.AuthTokenRequest{
		UserID: user.ID,
	})
	if err != nil {
		return ucModel.LoginResponse{}, err
	}

	return ucModel.LoginResponse{
		AccessToken: accessToken.AccessToken,
	}, nil
}

type ILogin interface {
	Login(ctx context.Context, req ucModel.LoginRequest) (ucModel.LoginResponse, error)
}
