package project

import obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models"

type (
	Project struct {
		ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
		Name        string `gorm:"column:name" json:"name"`
		Description string `gorm:"column:description" json:"description"`
		UserID      string `gorm:"type:uuid;column:created_by" json:"created_by"`
		obModel.BaseModel
	}
)
