package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func DeleteProject(gCtx *gin.Context) {
	projectId := gCtx.Param("id")
	if projectId == "" {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, "Project ID is required")
		return
	}

	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	response, err := usecaseProject.DeleteProject(gCtx.Request.Context(), projectId, userId)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, response, nil)
	return
}
