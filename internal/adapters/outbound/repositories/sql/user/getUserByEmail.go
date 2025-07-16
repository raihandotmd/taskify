package user

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	"gorm.io/gorm"
)

func (r *repository) GetUserByEmail(ctx context.Context, orm *gorm.DB, email string) (obModel.User, error) {
	var user obModel.User

	if err := orm.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return obModel.User{}, err
	}

	return user, nil
}

type IGetUserByEmail interface {
	GetUserByEmail(ctx context.Context, orm *gorm.DB, email string) (obModel.User, error)
}
