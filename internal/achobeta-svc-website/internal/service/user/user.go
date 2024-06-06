package user

import (
	"achobeta-svc/internal/achobeta-svc-common/lib/tlog"
	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/internal/entity"
	"context"
	"fmt"
)

// Create 创建用户
func Create(ctx context.Context, user *entity.User) (uint, error) {
	if err := config.GetMysql().Create(user).Error; err != nil {
		tlog.CtxErrorf(ctx, "create user error: %v", err)
		return 0, err
	}
	return user.ID, nil
}

func Modify(ctx context.Context, user *entity.User) error {
	if user == nil {
		return fmt.Errorf("modify user error, user is nil")
	}
	if err := config.GetMysql().Updates(user).Error; err != nil {
		return err
	}
	tlog.Infof("modify user info success, id: %d", user.ID)
	return nil
}

func Query(ctx context.Context, user *entity.User) error {
	if err := config.GetMysql().Where(user).First(user).Error; err != nil {
		tlog.CtxErrorf(ctx, "query user error: %v", err)
		return err
	}
	return nil
}
