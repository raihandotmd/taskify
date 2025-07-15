package taskifyGin

import (
	"github.com/gin-gonic/gin"
	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
)

func NewGinClient(cfg taskifyConfig.App) *gin.Engine {
	gin.SetMode(cfg.Mode)
	ginClient := gin.Default()

	return ginClient
}
