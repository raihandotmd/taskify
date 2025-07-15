package token

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/token"
)

func (uc *ucAuthToken) Generate(ctx context.Context, req ucModel.AuthTokenRequest) (ucModel.AuthTokenResponse, error) {
	tokenEnv, err := taskifyConfig.NewTokenConfig()
	if err != nil {
		return ucModel.AuthTokenResponse{}, err
	}

	timeNow := time.Now()
	accessToken := ucModel.AccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			ExpiresAt: jwt.NewNumericDate(timeNow.Add(tokenEnv.TokenExpiration)),
		},
		Data: ucModel.AccessTokenClaims{
			UserId: req.UserID,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)

	tokenString, err := jwtToken.SignedString([]byte(tokenEnv.TokenSecret))
	if err != nil {
		return ucModel.AuthTokenResponse{}, err
	}

	return ucModel.AuthTokenResponse{AccessToken: tokenString}, nil
}

type IGenerate interface {
	Generate(ctx context.Context, req ucModel.AuthTokenRequest) (ucModel.AuthTokenResponse, error)
}
