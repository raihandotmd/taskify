package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func GetAllTasksByProjectId(gCtx *gin.Context) {
	projectId := gCtx.Param("id")
	if projectId == "" {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, "Project ID is required")
		return
	}

	response, err := usecaseTask.GetAllTasksByProjectId(gCtx.Request.Context(), projectId)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, response, nil)
}
