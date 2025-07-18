package taskifyAuth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	obCacheUser "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/cache/user"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func TokenRevoke(c *gin.Context) {
	var cacheUser = obCacheUser.New()

	userId, err := GetUserId(c)
	if err != nil {
		taskifyGin.AbortJSONResponse(c, http.StatusUnauthorized, nil, errors.New("unauthorized"))
		return
	}

	userIdData := obModel.RevokeUserToken{
		UserID: userId,
	}

	exists, err := cacheUser.ExistToken(c.Request.Context(), userIdData)
	if err != nil {
		taskifyGin.AbortJSONResponse(c, http.StatusInternalServerError, nil, err)
		return
	}
	if exists {
		taskifyGin.AbortJSONResponse(c, http.StatusUnauthorized, nil, errors.New("unauthorized token revoked"))
		return
	}

	c.Next()
}
