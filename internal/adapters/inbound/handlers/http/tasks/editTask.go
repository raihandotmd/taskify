package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/task"
)

func EditTask(gCtx *gin.Context) error {
	var req ibModel.EditTaskRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	req.ID = gCtx.Param("id")
	if req.ID == "" {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
		return nil
	}

	task, err := usecaseTask.EditTask(gCtx.Request.Context(), req.ToUcModel())
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	gCtx.JSON(http.StatusOK, task)
	return nil
}
