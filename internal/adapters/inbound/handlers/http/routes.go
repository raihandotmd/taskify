package taskifyHttp

import "github.com/gin-gonic/gin"

func SetupRoutes(ginClient *gin.Engine) {

	ginClient.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

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
