package taskifyHttp

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	projectHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/projects"
	taskHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/tasks"
	userHandler "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/users"
	middleware "github.com/raihandotmd/taskify/internal/adapters/inbound/middleware/auth"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func SetupRoutes(ginClient *gin.Engine) {

	ginClient.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	ginClient.NoRoute(func(c *gin.Context) {
		taskifyGin.NewJSONResponse(c, http.StatusNotFound, nil, errors.New("Route not found"))
	})

	auth := ginClient.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
		auth.Use(middleware.JWTAuth()).Use(middleware.TokenRevoke).GET("/logout", userHandler.Logout)
	}

	api := ginClient.Group("/api")
	api.Use(middleware.JWTAuth()).Use(middleware.TokenRevoke)
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/projects", projectHandler.NewProject)
			v1.GET("/projects", projectHandler.GetAllProjectByUserId)
			v1.PUT("/projects/:id", projectHandler.EditProject)
			v1.DELETE("/projects/:id", projectHandler.DeleteProject)

			v1.POST("/tasks", taskHandler.NewTask)
			v1.GET("/projects/:id/tasks", taskHandler.GetAllTasksByProjectId)
			v1.PUT("/tasks/:id", taskHandler.EditTask)
			v1.DELETE("/tasks/:id", taskHandler.DeleteTask)
		}
	}
}
