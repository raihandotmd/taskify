package model

import obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models"

type (
	User struct {
		ID       string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
		Name     string `gorm:"column:name" json:"name"`
		Email    string `gorm:"column:email" json:"email"`
		Password string `gorm:"column:password" json:"-"`
		obModel.BaseModel
	}
)
