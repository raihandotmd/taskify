package main

import (
	"fmt"
	"log"

	taskifyHttp "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http"
	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
	taskifyRedis "github.com/raihandotmd/taskify/pkg/redis"
)

func main() {
	if err := taskifyGorm.ConnectDatabase(); err != nil {
		log.Fatalf("[DATABASE] connection failed: %v", err)
	}
	fmt.Println("[DATABASE] connection established successfully")

	if err := taskifyRedis.ConnectCache(); err != nil {
		log.Fatalf("[REDIS] connection failed: %v", err)
	}
	fmt.Println("[REDIS] connection established successfully")

	appConfig, err := taskifyConfig.NewAppConfig()
	if err != nil {
		fmt.Println("[CONFIG] Error loading app config:", err)
		return
	}
	ginClient := taskifyGin.NewGinClient(appConfig)

	taskifyHttp.SetupRoutes(ginClient)

	serverAddr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	log.Printf("Starting server on %s", serverAddr)
	ginClient.Run(serverAddr)
}
