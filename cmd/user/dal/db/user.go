package db

import (
	"context"

	"micro/cmd/user/model"
)

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *model.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}