package taskifyHttp

import (
	"github.com/gin-gonic/gin"
	projectHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/projects"
	taskHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/tasks"
	userHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/users"
	middleware "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
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
	api.Use(middleware.JWTAuth())
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/projects", func(c *gin.Context) {
				if err := projectHandler.NewProject(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
			v1.GET("/projects", func(c *gin.Context) {
				if err := projectHandler.GetAllProjectByUserId(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
			v1.PUT("/projects/:id", func(c *gin.Context) {
				if err := projectHandler.EditProject(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
			v1.DELETE("/projects/:id", func(c *gin.Context) {
				if err := projectHandler.DeleteProject(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})

			// TODO: Add API routes here
			v1.POST("/tasks", func(c *gin.Context) {
				if err := taskHandler.NewTask(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
			v1.GET("/projects/:id/tasks", func(c *gin.Context) {
				if err := taskHandler.GetAllTasksByProjectId(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
			v1.PUT("/tasks/:id", func(c *gin.Context) {
				if err := taskHandler.EditTask(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
			v1.DELETE("/tasks/:id", func(c *gin.Context) {
				if err := taskHandler.DeleteTask(c); err != nil {
					c.JSON(400, gin.H{"error": err.Error()})
					return
				}
			})
		}
	}
}
