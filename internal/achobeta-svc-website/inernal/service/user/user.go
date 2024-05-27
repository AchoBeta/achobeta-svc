package user

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/tlog"
	"achobeta-svc/internal/achobeta-svc-website/config"
	"achobeta-svc/internal/achobeta-svc-website/inernal/entity"
	"context"
)

func CreateUser(ctx context.Context, ue *entity.User) error {
	tlog.CtxInfof(ctx, "create user, username:[%s], email:[%s], phone:[%s]", ue.Username, *ue.Email, ue.Phone)
	result := config.GetMysql().Create(&ue)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
