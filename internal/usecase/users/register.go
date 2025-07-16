package users

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	authModel "github.com/raihandotmd/taskify/internal/usecase/models/token"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/user"
	"github.com/raihandotmd/taskify/internal/usecase/token"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
	"golang.org/x/crypto/bcrypt"
)

var ucToken = token.New()

func (uc *ucUser) Register(ctx context.Context, req ucModel.RegisterRequest) (ucModel.RegisterResponse, error) {
	var orm = taskifyGorm.DB

	var (
		user obModel.User
		err  error
	)

	// hash password
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ucModel.RegisterResponse{}, err
	}
	req.Password = string(bcryptPassword)

	// create user
	if user, err = repoUser.CreateUser(ctx, orm, req.ToObModel()); err != nil {
		return ucModel.RegisterResponse{}, err
	}

	// generate access token
	accessToken, err := ucToken.Generate(ctx, authModel.AuthTokenRequest{
		UserID: user.ID,
	})
	if err != nil {
		return ucModel.RegisterResponse{}, err
	}

	return ucModel.RegisterResponse{
		AccessToken: accessToken.AccessToken,
	}, nil
}

type IRegister interface {
	Register(ctx context.Context, req ucModel.RegisterRequest) (ucModel.RegisterResponse, error)
}
