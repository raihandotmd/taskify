package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/task"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func EditTask(gCtx *gin.Context) {
	var req ibModel.EditTaskRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, err)
		return
	}

	req.ID = gCtx.Param("id")
	if req.ID == "" {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, "task ID is required")
		return
	}

	task, err := usecaseTask.EditTask(gCtx.Request.Context(), req.ToUcModel())
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, task, nil)
}
