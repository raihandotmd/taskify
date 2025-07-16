package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/task"
)

func NewTask(gCtx *gin.Context) error {
	var req ibModel.NewTaskRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	task, err := usecaseTask.NewTask(gCtx.Request.Context(), req.ToUcModel())
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	gCtx.JSON(http.StatusCreated, task)
	return nil
}
