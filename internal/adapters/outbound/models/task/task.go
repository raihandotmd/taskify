package task

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models"
)

type (
	Task struct {
		ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
		ProjectID   string `gorm:"type:uuid;column:project_id" json:"project_id"`
		Title       string `gorm:"column:title" json:"title"`
		Description string `gorm:"column:description" json:"description"`
		Status      string `gorm:"column:status" json:"status"`
		Deadline    string `gorm:"type:date;column:deadline" json:"deadline"`
		obModel.BaseModel
	}
)
