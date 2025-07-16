package taskifyHttp

import (
	"github.com/gin-gonic/gin"
	userHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/users"
)

func SetupRoutes(ginClient *gin.Engine) {

	ginClient.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	auth := ginClient.Group("/auth")
	{
		auth.POST("/register", func(c *gin.Context) {
			if err := userHandler.Register(c); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		})
		auth.POST("/login", func(c *gin.Context) {
			if err := userHandler.Login(c); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		})
	}

	api := ginClient.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// TODO: Add API routes here
			v1.GET("/tasks", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "List of tasks",
				})
			})
		}
	}
}
