package user

type (
	LogoutRequest struct {
		UserID    string `json:"user_id"`
		ExpiresAt int64  `json:"expires_at"`
	}
)
