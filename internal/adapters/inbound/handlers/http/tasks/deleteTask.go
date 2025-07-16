package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func DeleteTask(gCtx *gin.Context) {
	taskId := gCtx.Param("id")
	if taskId == "" {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, "Task ID is required")
		return
	}

	response, err := usecaseTask.DeleteTask(gCtx.Request.Context(), taskId)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, response, nil)
}
