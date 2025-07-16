package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasksByProjectId(gCtx *gin.Context) error {
	projectId := gCtx.Param("id")
	if projectId == "" {
		return nil
	}

	response, err := usecaseTask.GetAllTasksByProjectId(gCtx.Request.Context(), projectId)
	if err != nil {
		return nil
	}

	gCtx.JSON(http.StatusOK, response)
	return nil
}
