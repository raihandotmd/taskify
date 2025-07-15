package taskifyGorm

import (
	"fmt"

	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	config, err := taskifyConfig.NewPGConfig()
	if err != nil {
		return fmt.Errorf("failed to load database configuration: %v", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.DBName, config.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}
	return nil
}
