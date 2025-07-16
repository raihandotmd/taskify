package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/project"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func EditProject(gCtx *gin.Context) {
	var req ibModel.EditProjectRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, err)
		return
	}

	req.ID = gCtx.Param("id")
	if req.ID == "" {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, "Project ID is required")
		return
	}

	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusUnauthorized, nil, "unauthorized")
		return
	}

	project, err := usecaseProject.EditProject(gCtx.Request.Context(), req.ToUcModel(), userId)
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusOK, project, nil)
}
