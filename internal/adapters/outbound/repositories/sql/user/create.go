package user

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	"gorm.io/gorm"
)

func (r *repository) CreateUser(ctx context.Context, orm *gorm.DB, req obModel.User) (obModel.User, error) {

	if err := orm.WithContext(ctx).Create(&req).Error; err != nil {
		return obModel.User{}, err
	}
	return req, nil
}

type ICreate interface {
	CreateUser(ctx context.Context, orm *gorm.DB, user obModel.User) (obModel.User, error)
}
