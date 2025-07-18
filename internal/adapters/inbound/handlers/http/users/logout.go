package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/user"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func Logout(gCtx *gin.Context) {
	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	tokenExpiresAt, err := taskifyAuth.GetExpiresAt(gCtx)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	req := ucModel.LogoutRequest{
		UserID:    userId,
		ExpiresAt: tokenExpiresAt,
	}

	if err := usecaseUser.Logout(gCtx.Request.Context(), req); err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, gin.H{"message": "Logged out successfully"}, nil)
}
