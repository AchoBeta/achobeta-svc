package user

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"context"
)

// Create 创建用户
func Create(ctx context.Context, user *entity.User) (uint, error) {
	if err := config.GetMysql().Create(user).Error; err != nil {
		tlog.CtxErrorf(ctx, "create user error: %v", err)
		return 0, err
	}
	return user.ID, nil
}

func Query(ctx context.Context, user *entity.User) {

}
