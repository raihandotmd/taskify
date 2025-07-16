package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func GetAllProjectByUserId(gCtx *gin.Context) {
	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusUnauthorized, nil, err)
		return
	}

	response, err := usecaseProject.GetAllProjectByUserId(gCtx.Request.Context(), userId)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, response, nil)
}
