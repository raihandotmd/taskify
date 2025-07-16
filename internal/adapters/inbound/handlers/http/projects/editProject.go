package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/project"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
)

func EditProject(gCtx *gin.Context) error {
	var req ibModel.EditProjectRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	req.ID = gCtx.Param("id")
	if req.ID == "" {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "project ID is required"})
		return nil
	}

	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return err
	}

	project, err := usecaseProject.EditProject(gCtx.Request.Context(), req.ToUcModel(), userId)
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	gCtx.JSON(http.StatusOK, project)
	return nil
}
