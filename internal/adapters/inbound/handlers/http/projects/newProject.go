package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/project"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func NewProject(gCtx *gin.Context) {
	var req ibModel.NewProjectRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, err)
		return
	}

	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusUnauthorized, nil, err)
		return
	}

	project, err := usecaseProject.NewProject(gCtx.Request.Context(), req.ToUcModel(), userId)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusCreated, project, nil)
}
