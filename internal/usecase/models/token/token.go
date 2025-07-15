package token

import "github.com/golang-jwt/jwt/v5"

type (
	AuthTokenRequest struct {
		UserID string `json:"user_id"`
	}

	AuthTokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	AccessToken struct {
		jwt.RegisteredClaims
		Data AccessTokenClaims `json:"dat"`
	}

	AccessTokenClaims struct {
		UserId string `json:"user_id"`
	}
)
