package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskifyAuth "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
)

func GetAllProjectByUserId(gCtx *gin.Context) error {
	userId, err := taskifyAuth.GetUserId(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return err
	}

	response, err := usecaseProject.GetAllProjectByUserId(gCtx.Request.Context(), userId)
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	gCtx.JSON(http.StatusOK, response)
	return nil
}
