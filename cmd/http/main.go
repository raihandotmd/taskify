package main

import (
	"fmt"
	"log"

	taskifyHttp "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http"
	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func main() {
	appConfig, err := taskifyConfig.NewAppConfig()
	if err != nil {
		fmt.Println("Error loading app config:", err)
		return
	}
	ginClient := taskifyGin.NewGinClient(appConfig)

	taskifyHttp.SetupRoutes(ginClient)

	serverAddr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	log.Printf("Starting server on %s", serverAddr)
	ginClient.Run(serverAddr)
}
