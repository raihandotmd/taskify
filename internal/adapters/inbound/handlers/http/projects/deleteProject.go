package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
)

func DeleteProject(gCtx *gin.Context) error {
	projectId := gCtx.Param("id")
	if projectId == "" {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "project ID is required"})
		return nil
	}

	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return err
	}

	response, err := usecaseProject.DeleteProject(gCtx.Request.Context(), projectId, userId)
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	gCtx.JSON(http.StatusOK, response)
	return nil
}
