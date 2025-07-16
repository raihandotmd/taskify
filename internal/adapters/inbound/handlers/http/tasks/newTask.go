package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/task"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func NewTask(gCtx *gin.Context) {
	var req ibModel.NewTaskRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, err)
		return
	}

	task, err := usecaseTask.NewTask(gCtx.Request.Context(), req.ToUcModel())
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusCreated, task, nil)
}
