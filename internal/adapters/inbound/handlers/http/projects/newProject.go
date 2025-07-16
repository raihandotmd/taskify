package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/project"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
)

func NewProject(gCtx *gin.Context) error {
	var req ibModel.NewProjectRequest
	if err := gCtx.ShouldBindJSON(&req); err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return err
	}

	project, err := usecaseProject.NewProject(gCtx.Request.Context(), req.ToUcModel(), userId)
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	gCtx.JSON(http.StatusCreated, project)
	return nil
}
