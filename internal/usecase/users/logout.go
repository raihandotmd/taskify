package users

import (
	"context"
	"time"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/user"
)

func (uc *ucUser) Logout(ctx context.Context, req ucModel.LogoutRequest) error {
	expiresAt := time.Unix(req.ExpiresAt, 0)
	ttl := time.Until(expiresAt)

	revokeTokenData := obModel.RevokeUserToken{
		UserID:    req.UserID,
		ExpiresAt: ttl,
	}

	if err := cacheUser.RevokeUserToken(ctx, revokeTokenData); err != nil {
		return err
	}

	return nil
}

type ILogout interface {
	Logout(ctx context.Context, req ucModel.LogoutRequest) error
}
