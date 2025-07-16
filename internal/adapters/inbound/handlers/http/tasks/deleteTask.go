package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTask(gCtx *gin.Context) error {
	taskId := gCtx.Param("id")
	if taskId == "" {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return nil
	}

	response, err := usecaseTask.DeleteTask(gCtx.Request.Context(), taskId)
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	gCtx.JSON(http.StatusOK, response)
	return nil
}
